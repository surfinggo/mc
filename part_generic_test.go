package mc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSliceContains(t *testing.T) {
	assertions := require.New(t)
	assertions.True(SliceContains([]string{"a", "b"}, "a"))
	assertions.False(SliceContains([]string{"a", "b"}, "x"))
	assertions.True(SliceContains([]int{1, 2}, 1))
	assertions.False(SliceContains([]int{1, 2}, 0))
	assertions.True(SliceContains([]uint{1, 2}, uint(1)))
	assertions.False(SliceContains([]uint{1, 2}, uint(0)))
	assertions.True(SliceContains([]float64{1, 2}, float64(1)))
	assertions.False(SliceContains([]float64{1, 2}, float64(0)))
}

func TestPointerTo(t *testing.T) {
	assertions := require.New(t)
	stringPointer := PointerTo("a")
	acceptStringPointer := func(p *string) string {
		return *p
	}
	stringValue := acceptStringPointer(stringPointer)
	assertions.Equal("a", stringValue)

	PointerTo(true)
	PointerTo("a")
	PointerTo(1)
	PointerTo(int8(1))
	PointerTo(int16(1))
	PointerTo(int32(1))
	PointerTo(int64(1))
	PointerTo(uint(1))
	PointerTo(uint8(1))
	PointerTo(uint16(1))
	PointerTo(uint32(1))
	PointerTo(uint64(1))
	PointerTo(float32(1))
	PointerTo(float64(1))

	PtrTo("a")
}
