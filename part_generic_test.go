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

func TestVarOr(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("t", VarOr("t", "default text"))
	assertions.Equal("default text", VarOr("", "default text"))
	assertions.Equal(1, VarOr(1, 99))
	assertions.Equal(99, VarOr(0, 99))
	assertions.Equal(int8(1), VarOr(int8(1), int8(99)))
	assertions.Equal(int8(99), VarOr(int8(0), int8(99)))
	assertions.Equal(int16(1), VarOr(int16(1), int16(99)))
	assertions.Equal(int16(99), VarOr(int16(0), int16(99)))
	assertions.Equal(int32(1), VarOr(int32(1), int32(99)))
	assertions.Equal(int32(99), VarOr(int32(0), int32(99)))
	assertions.Equal(int64(1), VarOr(int64(1), int64(99)))
	assertions.Equal(int64(99), VarOr(int64(0), int64(99)))
	assertions.Equal(uint(1), VarOr(uint(1), uint(99)))
	assertions.Equal(uint(99), VarOr(uint(0), uint(99)))
	assertions.Equal(uint8(1), VarOr(uint8(1), uint8(99)))
	assertions.Equal(uint8(99), VarOr(uint8(0), uint8(99)))
	assertions.Equal(uint16(1), VarOr(uint16(1), uint16(99)))
	assertions.Equal(uint16(99), VarOr(uint16(0), uint16(99)))
	assertions.Equal(uint32(1), VarOr(uint32(1), uint32(99)))
	assertions.Equal(uint32(99), VarOr(uint32(0), uint32(99)))
	assertions.Equal(uint64(1), VarOr(uint64(1), uint64(99)))
	assertions.Equal(uint64(99), VarOr(uint64(0), uint64(99)))
	assertions.Equal(float32(1), VarOr(float32(1), float32(99)))
	assertions.Equal(float32(99), VarOr(float32(0), float32(99)))
	assertions.Equal(float64(1), VarOr(float64(1), float64(99)))
	assertions.Equal(float64(99), VarOr(float64(0), float64(99)))
}
