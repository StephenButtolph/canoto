package canoto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBytesEmpty(t *testing.T) {
	require := require.New(t)

	require.True(isBytesEmpty(make([]byte, 0)))
	require.True(isBytesEmpty(make([]byte, 10)))

	require.False(isBytesEmpty([]byte{0: 1}))
	require.False(isBytesEmpty([]byte{10: 1}))
}
