// Canoto is a command to generate code for reading and writing the canoto
// format.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/generate"
)

const (
	canotoFlag   = "canoto"
	libraryFlag  = "library"
	protoFlag    = "proto"
	versionFlag  = "version"
	importFlag   = "import"
	internalFlag = "internal"

	formatCacheFlag      = "format-cache"
	formatNumberFlag     = "format-number"
	formatTagFlag        = "format-tag"
	formatOneOfTypeFlag  = "format-oneof-type"
	formatOneOfUnsetFlag = "format-oneof-unset"
	formatOneOfFieldFlag = "format-oneof-field"
)

func init() {
	cobra.EnablePrefixMatching = true
}

func main() {
	cmd := &cobra.Command{
		Use:   "canoto",
		Short: "Processes the provided files and generates the corresponding canoto and proto files",
		RunE: func(c *cobra.Command, args []string) error {
			flags := c.Flags()
			showVersion, err := flags.GetBool(versionFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", versionFlag, err)
			}
			if showVersion {
				fmt.Println("canoto/" + canoto.Version)
				return nil
			}

			library, err := flags.GetString(libraryFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", libraryFlag, err)
			}
			if library != "" {
				if err := generate.Library(library); err != nil {
					return fmt.Errorf("generating library in %q: %w", library, err)
				}
			}

			canoto, err := flags.GetBool(canotoFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", canotoFlag, err)
			}
			proto, err := flags.GetBool(protoFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", protoFlag, err)
			}

			canotoImport, err := flags.GetString(importFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", importFlag, err)
			}
			internal, err := flags.GetBool(internalFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", internalFlag, err)
			}

			cacheTemplate, err := flags.GetString(formatCacheFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", formatCacheFlag, err)
			}
			numberTemplate, err := flags.GetString(formatNumberFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", formatNumberFlag, err)
			}
			tagTemplate, err := flags.GetString(formatTagFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", formatTagFlag, err)
			}
			oneOfTypeTemplate, err := flags.GetString(formatOneOfTypeFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", formatOneOfTypeFlag, err)
			}
			oneOfUnsetTemplate, err := flags.GetString(formatOneOfUnsetFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", formatOneOfUnsetFlag, err)
			}
			oneOfFieldTemplate, err := flags.GetString(formatOneOfFieldFlag)
			if err != nil {
				return fmt.Errorf("getting %q: %w", formatOneOfFieldFlag, err)
			}

			opts := generate.Options{
				CanotoImport: canotoImport,
				Internal:     internal,
				Templates: generate.Templates{
					Cache:      cacheTemplate,
					Number:     numberTemplate,
					Tag:        tagTemplate,
					OneOfType:  oneOfTypeTemplate,
					OneOfUnset: oneOfUnsetTemplate,
					OneOfField: oneOfFieldTemplate,
				},
			}
			for _, arg := range args {
				if canoto {
					if err := generate.Canoto(arg, opts); err != nil {
						return fmt.Errorf("generating canoto for %q: %w", arg, err)
					}
				}
				if proto {
					if err := generate.Proto(arg, opts); err != nil {
						return fmt.Errorf("generating proto for %q: %w", arg, err)
					}
				}
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.Bool(versionFlag, false, "Display the version and exit")
	flags.Bool(canotoFlag, true, "Generate canoto file")
	flags.String(libraryFlag, "", "Generate the canoto library in the specified directory")
	flags.Bool(protoFlag, false, "Generate proto file")
	flags.String(importFlag, "", "Package to depend on for canoto serialization primitives")
	flags.Bool(internalFlag, false, "Generate a file that assumes the canoto package does not need to be imported")
	flags.String(formatCacheFlag, "", "Format to use when generating the canoto cache")
	flags.String(formatNumberFlag, "", "Format to use when generating canoto field number constants")
	flags.String(formatTagFlag, "", "Format to use when generating canoto field tag constants")
	flags.String(formatOneOfTypeFlag, "", "Format to use when generating canoto oneOf types")
	flags.String(formatOneOfUnsetFlag, "", "Format to use when generating canoto unset oneOf constants")
	flags.String(formatOneOfFieldFlag, "", "Format to use when generating canoto oneOf field constants")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed %v\n", err)
		os.Exit(1)
	}
}
