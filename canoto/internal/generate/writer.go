package generate

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/StephenButtolph/canoto"
)

const repeated = true

func write(w io.Writer, packageName string, messages []message) error {
	const fileTemplate = `// Code generated by Canoto. DO NOT EDIT.

package ${package}

import (
	"io"
	"sync/atomic"
	"unicode/utf8"

	"github.com/StephenButtolph/canoto"
)

// Ensure that unused imports do not error
var (
	_ = io.ErrUnexpectedEOF
	_ = utf8.ValidString
)
`
	err := writeTemplate(w, fileTemplate, map[string]string{
		"package": packageName,
	})
	if err != nil {
		return err
	}

	for _, m := range messages {
		if err := writeStruct(w, m); err != nil {
			return err
		}
	}
	return nil
}

func writeStruct(w io.Writer, m message) error {
	const structTemplate = `
const (
${tagConstants}
${tagSizeConstants})

type canotoData_${structName} struct {
	size atomic.Int64
${cache}}

func (c *${structName}) UnmarshalCanoto(bytes []byte) error {
	r := canoto.Reader{
		B: bytes,
	}
	return c.UnmarshalCanotoFrom(&r)
}

func (c *${structName}) UnmarshalCanotoFrom(r *canoto.Reader) error {
	var minField uint32
	for canoto.HasNext(r) {
		field, wireType, err := canoto.ReadTag(r)
		if err != nil {
			return err
		}
		if field < minField {
			return canoto.ErrInvalidFieldOrder
		}

		switch field {
${unmarshal}		default:
			return canoto.ErrUnknownField
		}

		minField = field + 1
	}
	return nil
}

func (c *${structName}) ValidCanoto() bool {
${valid}	return true
}

func (c *${structName}) CalculateCanotoSize() int {
	var size int
${size}	c.canotoData.size.Store(int64(size))
	return size
}

func (c *${structName}) CachedCanotoSize() int {
	return int(c.canotoData.size.Load())
}

func (c *${structName}) MarshalCanoto() []byte {
	w := canoto.Writer{
		B: make([]byte, 0, c.CalculateCanotoSize()),
	}
	c.MarshalCanotoInto(&w)
	return w.B
}

func (c *${structName}) MarshalCanotoInto(w *canoto.Writer) {
${marshal}}
`

	tagConstants, err := makeTagConstants(m)
	if err != nil {
		return err
	}
	unmarshal, err := makeUnmarshal(m)
	if err != nil {
		return err
	}
	size, err := makeSize(m)
	if err != nil {
		return err
	}
	marshal, err := makeMarshal(m)
	if err != nil {
		return err
	}

	return writeTemplate(w, structTemplate, map[string]string{
		"tagConstants":     tagConstants,
		"tagSizeConstants": makeTagSizeConstants(m),
		"structName":       m.name,
		"cache":            makeCache(m),
		"unmarshal":        unmarshal,
		"valid":            makeValid(m),
		"size":             size,
		"marshal":          marshal,
	})
}

func makeTagConstants(m message) (string, error) {
	var s strings.Builder
	for _, f := range m.fields {
		_, _ = fmt.Fprintf(
			&s,
			`	canoto__%s__%s__tag = "`,
			m.canonicalizedName,
			f.canonicalizedName,
		)

		wireType, err := f.WireType()
		if err != nil {
			return "", err
		}

		tagBytes := canoto.Tag(f.fieldNumber, wireType)
		tagHex := hex.EncodeToString(tagBytes)
		for i := 0; i < len(tagHex); i += 2 {
			_, _ = fmt.Fprintf(&s, "\\x%s", tagHex[i:i+2])
		}
		_, _ = fmt.Fprintf(
			&s,
			"\" // canoto.Tag(%d, canoto.%s)\n",
			f.fieldNumber,
			wireType,
		)
	}
	return s.String(), nil
}

func makeTagSizeConstants(m message) string {
	var s strings.Builder
	for _, f := range m.fields {
		_, _ = fmt.Fprintf(
			&s,
			"\tcanoto__%s__%s__tag__size = len(canoto__%s__%s__tag)\n",
			m.canonicalizedName,
			f.canonicalizedName,
			m.canonicalizedName,
			f.canonicalizedName,
		)
	}
	return s.String()
}

func makeCache(m message) string {
	var s strings.Builder
	for _, f := range m.fields {
		if !f.repeated || !f.canotoType.IsVarint() {
			continue
		}

		_, _ = fmt.Fprintf(
			&s,
			"\t%sSize atomic.Int64\n",
			f.name,
		)
	}
	return s.String()
}

func makeUnmarshal(m message) (string, error) {
	const (
		intTemplate = `		case ${fieldNumber}:
			if wireType != canoto.${wireType} {
				return canoto.ErrInvalidWireType
			}
			c.${fieldName}, err = canoto.Read${readFunction}(r)
			if err != nil {
				return err
			}
			if canoto.IsZero(c.${fieldName}) {
				return canoto.ErrZeroValue
			}
`
		fixedRepeatedIntTemplate = `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			msgBytes, err := canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}

			remainingBytes := r.B
			r.B = msgBytes
			for i := range c.${fieldName} {
				v, err := canoto.Read${readFunction}(r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.${fieldName}[i] = v
			}
			hasNext := canoto.HasNext(r)
			r.B = remainingBytes
			if hasNext {
				return io.ErrUnexpectedEOF
			}
			if canoto.IsZero(c.${fieldName}) {
				return canoto.ErrZeroValue
			}
`
		repeatedFixedSizeTemplate = `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			msgBytes, err := canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			numMsgBytes := len(msgBytes)
			if numMsgBytes == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.${fieldName} = make([]${goType}, 0, numMsgBytes/canoto.Size${sizeConstant})
			for canoto.HasNext(r) {
				v, err := canoto.Read${readFunction}(r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.${fieldName} = append(c.${fieldName}, v)
			}
			r.B = remainingBytes
`
		bytesTemplate = `		case ${fieldNumber}:
			if wireType != canoto.${wireType} {
				return canoto.ErrInvalidWireType
			}
			c.${fieldName}, err = canoto.Read${readFunction}(r)
			if err != nil {
				return err
			}
			if len(c.${fieldName}) == 0 {
				return canoto.ErrZeroValue
			}
`
		repeatedBytesTemplate = `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			v, err := canoto.Read${readFunction}(r)
			if err != nil {
				return err
			}

			count, err := canoto.CountBytes(r.B, canoto__${escapedStructName}__${escapedFieldName}__tag)
			if err != nil {
				return err
			}

			c.${fieldName} = make([]${goType}, 1, 1+count)
			c.${fieldName}[0] = v
			for range count {
				r.B = r.B[canoto__${escapedStructName}__${escapedFieldName}__tag__size:]
				v, err := canoto.Read${readFunction}(r)
				if err != nil {
					return err
				}
				c.${fieldName} = append(c.${fieldName}, v)
			}
`
		fixedRepeatedBytes = `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			v, err := canoto.Read${readFunction}(r)
			if err != nil {
				return err
			}

			c.${fieldName}[0] = v
			isZero := len(v) == 0
			for i := range len(c.${fieldName})-1 {
				if !canoto.HasPrefix(r.B, canoto__${escapedStructName}__${escapedFieldName}__tag) {
					return canoto.ErrUnknownField
				}
				r.B = r.B[canoto__${escapedStructName}__${escapedFieldName}__tag__size:]
				v, err := canoto.Read${readFunction}(r)
				if err != nil {
					return err
				}
				c.${fieldName}[1+i] = v
				isZero = isZero && len(v) == 0
			}
			if isZero {
				return canoto.ErrZeroValue
			}
`
	)
	return writeMessage(m, messageTemplate{
		ints: typeTemplate{
			single: intTemplate,
			repeated: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			msgBytes, err := canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.${fieldName} = make([]${goType}, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.Read${readFunction}(r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.${fieldName} = append(c.${fieldName}, v)
			}
			r.B = remainingBytes
`,
			fixedRepeated: fixedRepeatedIntTemplate,
		},
		fints: typeTemplate{
			single:        intTemplate,
			repeated:      repeatedFixedSizeTemplate,
			fixedRepeated: fixedRepeatedIntTemplate,
		},
		bools: typeTemplate{
			single:        intTemplate,
			repeated:      repeatedFixedSizeTemplate,
			fixedRepeated: fixedRepeatedIntTemplate,
		},
		strings: typeTemplate{
			single:        bytesTemplate,
			repeated:      repeatedBytesTemplate,
			fixedRepeated: fixedRepeatedBytes,
		},
		bytesTemplate:         bytesTemplate,
		repeatedBytesTemplate: repeatedBytesTemplate,
		fixedBytesTemplate: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			length, err := canoto.ReadInt[int32](r)
			if err != nil {
				return err
			}

			const (
				expectedLength      = len(c.${fieldName})
				expectedLengthInt32 = int32(expectedLength)
			)
			if length != expectedLengthInt32 {
				return canoto.ErrInvalidLength
			}
			if expectedLength > len(r.B) {
				return io.ErrUnexpectedEOF
			}

			copy(c.${fieldName}[:], r.B)
			if canoto.IsZero(c.${fieldName}) {
				return canoto.ErrZeroValue
			}
			r.B = r.B[expectedLength:]
`,
		repeatedFixedBytesTemplate: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			length, err := canoto.ReadInt[int32](r)
			if err != nil {
				return err
			}

			const (
				expectedLength      = len(c.${fieldName}[0])
				expectedLengthInt32 = int32(expectedLength)
			)
			if length != expectedLengthInt32 {
				return canoto.ErrInvalidLength
			}
			if expectedLength > len(r.B) {
				return io.ErrUnexpectedEOF
			}

			firstEntry := r.B[:expectedLength]
			r.B = r.B[expectedLength:]
			count, err := canoto.CountBytes(r.B, canoto__${escapedStructName}__${escapedFieldName}__tag)
			if err != nil {
				return err
			}

			c.${fieldName} = canoto.MakeSlice(c.${fieldName}, 1+count)
			copy(c.${fieldName}[0][:], firstEntry)
			for i := range count {
				r.B = r.B[canoto__${escapedStructName}__${escapedFieldName}__tag__size:]
				length, err := canoto.ReadInt[int32](r)
				if err != nil {
					return err
				}
				if length != expectedLengthInt32 {
					return canoto.ErrInvalidLength
				}
				if expectedLength > len(r.B) {
					return io.ErrUnexpectedEOF
				}

				copy(c.${fieldName}[1+i][:], r.B)
				r.B = r.B[expectedLength:]
			}
`,
		fixedRepeatedBytesTemplate: fixedRepeatedBytes,
		fixedRepeatedFixedBytesTemplate: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			length, err := canoto.ReadInt[int32](r)
			if err != nil {
				return err
			}

			const (
				expectedLength      = len(c.${fieldName}[0])
				expectedLengthInt32 = int32(expectedLength)
			)
			if length != expectedLengthInt32 {
				return canoto.ErrInvalidLength
			}
			if expectedLength > len(r.B) {
				return io.ErrUnexpectedEOF
			}

			copy(c.${fieldName}[0][:], r.B)
			r.B = r.B[expectedLength:]
			for i := range len(c.${fieldName})-1 {
				if !canoto.HasPrefix(r.B, canoto__${escapedStructName}__${escapedFieldName}__tag) {
					return canoto.ErrUnknownField
				}
				r.B = r.B[canoto__${escapedStructName}__${escapedFieldName}__tag__size:]

				length, err := canoto.ReadInt[int32](r)
				if err != nil {
					return err
				}
				if length != expectedLengthInt32 {
					return canoto.ErrInvalidLength
				}
				if expectedLength > len(r.B) {
					return io.ErrUnexpectedEOF
				}

				copy(c.${fieldName}[1+i][:], r.B)
				r.B = r.B[expectedLength:]
			}
			if canoto.IsZero(c.${fieldName}) {
				return canoto.ErrZeroValue
			}
`,
		customs: typeTemplate{
			single: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			msgBytes, err := canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			err = c.${fieldName}.UnmarshalCanotoFrom(r)
			r.B = remainingBytes
			if err != nil {
				return err
			}
`,
			repeated: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			msgBytes, err := canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}

			remainingBytes := r.B
			count, err := canoto.CountBytes(remainingBytes, canoto__${escapedStructName}__${escapedFieldName}__tag)
			if err != nil {
				return err
			}

			c.${fieldName} = make([]${goType}, 1+count)
			r.B = msgBytes
			err = c.${fieldName}[0].UnmarshalCanotoFrom(r)
			r.B = remainingBytes
			if err != nil {
				return err
			}

			for i := range count {
				r.B = r.B[canoto__${escapedStructName}__${escapedFieldName}__tag__size:]
				r.Unsafe = true
				msgBytes, err := canoto.ReadBytes(r)
				r.Unsafe = originalUnsafe
				if err != nil {
					return err
				}

				remainingBytes := r.B
				r.B = msgBytes
				err = c.${fieldName}[1+i].UnmarshalCanotoFrom(r)
				r.B = remainingBytes
				if err != nil {
					return err
				}
			}
`,
			fixedRepeated: `		case ${fieldNumber}:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			msgBytes, err := canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}

			remainingBytes := r.B
			r.B = msgBytes
			err = c.${fieldName}[0].UnmarshalCanotoFrom(r)
			r.B = remainingBytes
			if err != nil {
				return err
			}

			isZero := len(msgBytes) == 0
			for i := range len(c.${fieldName})-1 {
				if !canoto.HasPrefix(r.B, canoto__${escapedStructName}__${escapedFieldName}__tag) {
					return canoto.ErrUnknownField
				}
				r.B = r.B[canoto__${escapedStructName}__${escapedFieldName}__tag__size:]
				r.Unsafe = true
				msgBytes, err := canoto.ReadBytes(r)
				r.Unsafe = originalUnsafe
				if err != nil {
					return err
				}

				remainingBytes := r.B
				r.B = msgBytes
				err = c.${fieldName}[1+i].UnmarshalCanotoFrom(r)
				r.B = remainingBytes
				if err != nil {
					return err
				}
				isZero = isZero && len(msgBytes) == 0
			}
			if isZero {
				return canoto.ErrZeroValue
			}
`,
		},
	})
}

func makeValid(m message) string {
	const (
		stringTemplate = `	if !utf8.ValidString(c.${fieldName}) {
		return false
	}
`
		repeatedStringTemplate = `	for _, v := range c.${fieldName} {
		if !utf8.ValidString(v) {
			return false
		}
	}
`
		customTemplate = `	if !c.${fieldName}.ValidCanoto() {
		return false
	}
`
		repeatedCustomTemplate = `	for i := range c.${fieldName} {
		if !c.${fieldName}[i].ValidCanoto() {
			return false
		}
	}
`
	)
	var (
		stringTemplates = map[bool]string{
			!repeated: stringTemplate,
			repeated:  repeatedStringTemplate,
		}
		customTemplates = map[bool]string{
			!repeated: customTemplate,
			repeated:  repeatedCustomTemplate,
		}
		s strings.Builder
	)
	for _, f := range m.fields {
		if f.canotoType != canotoBytes || f.goType == goBytes {
			continue
		}

		// goType is either string or a custom type
		var template string
		if f.goType == goString {
			template = stringTemplates[f.repeated]
		} else {
			template = customTemplates[f.repeated]
		}
		_ = writeTemplate(&s, template, map[string]string{
			"fieldName": f.name,
		})
	}
	return s.String()
}

func makeSize(m message) (string, error) {
	const (
		fixedSizeTemplate = `	if !canoto.IsZero(c.${fieldName}) {
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.Size${sizeConstant}
	}
`
		repeatedFixedSizeTemplate = `	if num := len(c.${fieldName}); num != 0 {
		fieldSize := num * canoto.Size${sizeConstant}
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
`
		fixedRepeatedFixedSizeTemplate = `	if !canoto.IsZero(c.${fieldName}) {
		const fieldSize = len(c.${fieldName}) * canoto.Size${sizeConstant}
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + fieldSize + canoto.SizeInt(int64(fieldSize))
	}
`
		bytesTemplate = `	if len(c.${fieldName}) != 0 {
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(c.${fieldName})
	}
`
		repeatedBytesTemplate = `	for _, v := range c.${fieldName} {
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(v)
	}
`
	)
	return writeMessage(m, messageTemplate{
		ints: typeTemplate{
			single: `	if !canoto.IsZero(c.${fieldName}) {
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.Size${sizeFunction}(c.${fieldName})
	}
`,
			repeated: `	if len(c.${fieldName}) != 0 {
		var fieldSize int
		for _, v := range c.${fieldName} {
			fieldSize += canoto.Size${sizeFunction}(v)
		}
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
		c.canotoData.${fieldName}Size.Store(int64(fieldSize))
	}
`,
			fixedRepeated: `	if !canoto.IsZero(c.${fieldName}) {
		var fieldSize int
		for _, v := range c.${fieldName} {
			fieldSize += canoto.Size${sizeFunction}(v)
		}
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
		c.canotoData.${fieldName}Size.Store(int64(fieldSize))
	}
`,
		},
		fints: typeTemplate{
			single:        fixedSizeTemplate,
			repeated:      repeatedFixedSizeTemplate,
			fixedRepeated: fixedRepeatedFixedSizeTemplate,
		},
		bools: typeTemplate{
			single:        fixedSizeTemplate,
			repeated:      repeatedFixedSizeTemplate,
			fixedRepeated: fixedRepeatedFixedSizeTemplate,
		},
		strings: typeTemplate{
			single:   bytesTemplate,
			repeated: repeatedBytesTemplate,
			fixedRepeated: `	if !canoto.IsZero(c.${fieldName}) {
		for _, v := range c.${fieldName} {
			size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(v)
		}
	}
`,
		},
		bytesTemplate:         bytesTemplate,
		repeatedBytesTemplate: repeatedBytesTemplate,
		fixedBytesTemplate: `	if !canoto.IsZero(c.${fieldName}) {
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(c.${fieldName}[:])
	}
`,
		repeatedFixedBytesTemplate: `	if num := len(c.${fieldName}); num != 0 {
		fieldSize := canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(c.${fieldName}[0][:])
		size += num * fieldSize
	}
`,
		fixedRepeatedBytesTemplate: `	{
		isZero := true
		for _, v := range c.${fieldName} {
			if len(v) != 0 {
				isZero = false
				break
			}
		}
		if !isZero {
			for _, v := range c.${fieldName} {
				size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(v)
			}
		}
	}
`,
		fixedRepeatedFixedBytesTemplate: `	if !canoto.IsZero(c.${fieldName}) {
		for i := range c.${fieldName} {
			size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeBytes(c.${fieldName}[i][:])
		}
	}
`,
		customs: typeTemplate{
			single: `	if fieldSize := c.${fieldName}.CalculateCanotoSize(); fieldSize != 0 {
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
`,
			repeated: `	for i := range c.${fieldName} {
		fieldSize := c.${fieldName}[i].CalculateCanotoSize()
		size += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
`,
			fixedRepeated: `	{
		var (
			fieldSizeSum int
			totalSize    int
		)
		for i := range c.${fieldName} {
			fieldSize := c.${fieldName}[i].CalculateCanotoSize()
			fieldSizeSum += fieldSize
			totalSize += canoto__${escapedStructName}__${escapedFieldName}__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
		}
		if fieldSizeSum != 0 {
			size += totalSize
		}
	}
`,
		},
	})
}

func makeMarshal(m message) (string, error) {
	const (
		intTemplate = `	if !canoto.IsZero(c.${fieldName}) {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.Append${sizeFunction}${sizeConstant}(w, c.${fieldName})
	}
`
		repeatedFintTemplate = `	if num := len(c.${fieldName}); num != 0 {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendInt(w, int64(num*canoto.Size${sizeConstant}))
		for _, v := range c.${fieldName} {
			canoto.Append${sizeConstant}(w, v)
		}
	}
`
		fixedRepeatedFintTemplate = `	if !canoto.IsZero(c.${fieldName}) {
		const fieldSize = len(c.${fieldName})*canoto.Size${sizeConstant}
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendInt(w, int64(fieldSize))
		for _, v := range c.${fieldName} {
			canoto.Append${sizeConstant}(w, v)
		}
	}
`
		bytesTemplate = `	if len(c.${fieldName}) != 0 {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendBytes(w, c.${fieldName})
	}
`
		repeatedBytesTemplate = `	for _, v := range c.${fieldName} {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendBytes(w, v)
	}
`
	)
	return writeMessage(m, messageTemplate{
		ints: typeTemplate{
			single: intTemplate,
			repeated: `	if len(c.${fieldName}) != 0 {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendInt(w, c.canotoData.${fieldName}Size.Load())
		for _, v := range c.${fieldName} {
			canoto.Append${sizeFunction}(w, v)
		}
	}
`,
			fixedRepeated: `	if !canoto.IsZero(c.${fieldName}) {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendInt(w, c.canotoData.${fieldName}Size.Load())
		for _, v := range c.${fieldName} {
			canoto.Append${sizeFunction}(w, v)
		}
	}
`,
		},
		fints: typeTemplate{
			single:        intTemplate,
			repeated:      repeatedFintTemplate,
			fixedRepeated: fixedRepeatedFintTemplate,
		},
		bools: typeTemplate{
			single: `	if !canoto.IsZero(c.${fieldName}) {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendBool(w, true)
	}
`,
			repeated:      repeatedFintTemplate,
			fixedRepeated: fixedRepeatedFintTemplate,
		},
		strings: typeTemplate{
			single:   bytesTemplate,
			repeated: repeatedBytesTemplate,
			fixedRepeated: `	if !canoto.IsZero(c.${fieldName}) {
		for _, v := range c.${fieldName} {
			canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
			canoto.AppendBytes(w, v)
		}
	}
`,
		},
		bytesTemplate:         bytesTemplate,
		repeatedBytesTemplate: repeatedBytesTemplate,
		fixedBytesTemplate: `	if !canoto.IsZero(c.${fieldName}) {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendBytes(w, c.${fieldName}[:])
	}
`,
		repeatedFixedBytesTemplate: `	for i := range c.${fieldName} {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendBytes(w, c.${fieldName}[i][:])
	}
`,
		fixedRepeatedBytesTemplate: `	{
		isZero := true
		for _, v := range c.${fieldName} {
			if len(v) != 0 {
				isZero = false
				break
			}
		}
		if !isZero {
			for _, v := range c.${fieldName} {
				canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
				canoto.AppendBytes(w, v)
			}
		}
	}
`,
		fixedRepeatedFixedBytesTemplate: `	if !canoto.IsZero(c.${fieldName}) {
		for i := range c.${fieldName} {
			canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
			canoto.AppendBytes(w, c.${fieldName}[i][:])
		}
	}
`,
		customs: typeTemplate{
			single: `	if fieldSize := c.${fieldName}.CachedCanotoSize(); fieldSize != 0 {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendInt(w, int64(fieldSize))
		c.${fieldName}.MarshalCanotoInto(w)
	}
`,
			repeated: `	for i := range c.${fieldName} {
		canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
		canoto.AppendInt(w, int64(c.${fieldName}[i].CachedCanotoSize()))
		c.${fieldName}[i].MarshalCanotoInto(w)
	}
`,
			fixedRepeated: `	{
		isZero := true
		for i := range c.${fieldName} {
			if c.${fieldName}[i].CachedCanotoSize() != 0 {
				isZero = false
				break
			}
		}
		if !isZero {
			for i := range c.${fieldName} {
				canoto.Append(w, canoto__${escapedStructName}__${escapedFieldName}__tag)
				canoto.AppendInt(w, int64(c.${fieldName}[i].CachedCanotoSize()))
				c.${fieldName}[i].MarshalCanotoInto(w)
			}
		}
	}
`,
		},
	})
}

type messageTemplate struct {
	ints    typeTemplate
	fints   typeTemplate
	bools   typeTemplate
	strings typeTemplate

	bytesTemplate                   string
	repeatedBytesTemplate           string
	fixedBytesTemplate              string
	repeatedFixedBytesTemplate      string
	fixedRepeatedBytesTemplate      string
	fixedRepeatedFixedBytesTemplate string

	customs typeTemplate
}

type typeTemplate struct {
	single        string
	repeated      string
	fixedRepeated string
}

func (t *typeTemplate) Template(repeated, fixedLength bool) string {
	switch {
	case !repeated:
		return t.single
	case !fixedLength:
		return t.repeated
	default:
		return t.fixedRepeated
	}
}

func writeMessage(m message, t messageTemplate) (string, error) {
	var s strings.Builder
	for _, f := range m.fields {
		var template string
		switch f.canotoType {
		case canotoInt, canotoSint:
			template = t.ints.Template(f.repeated, f.fixedLength[0])
		case canotoFint:
			template = t.fints.Template(f.repeated, f.fixedLength[0])
		case canotoBool:
			template = t.bools.Template(f.repeated, f.fixedLength[0])
		case canotoBytes:
			switch f.goType {
			case goString:
				template = t.strings.Template(f.repeated, f.fixedLength[0])
			case goBytes:
				switch {
				case f.fixedLength[0] && !f.repeated:
					template = t.fixedBytesTemplate
				case !f.fixedLength[0] && f.fixedLength[1]:
					template = t.repeatedFixedBytesTemplate
				case f.fixedLength[0] && !f.fixedLength[1]:
					template = t.fixedRepeatedBytesTemplate
				case f.fixedLength[0] && f.fixedLength[1]:
					template = t.fixedRepeatedFixedBytesTemplate
				case !f.repeated:
					template = t.bytesTemplate
				default:
					template = t.repeatedBytesTemplate
				}
			default:
				template = t.customs.Template(f.repeated, f.fixedLength[0])
			}
		default:
			return "", fmt.Errorf("%w: %q", errUnexpectedCanotoType, f.canotoType)
		}
		_ = writeTemplate(&s, template, f.templateArgs)
	}
	return s.String(), nil
}

func writeTemplate(w io.Writer, template string, args map[string]string) error {
	s := os.Expand(template, func(key string) string {
		return args[key]
	})
	_, err := w.Write([]byte(s))
	return err
}
