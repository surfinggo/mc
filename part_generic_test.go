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

func TestSliceFilter(t *testing.T) {
	cases := []struct {
		message  string
		actual   any
		expected any
	}{
		{
			message: "filtering a slice of strings",
			actual: SliceFilter([]string{"a", "b"}, func(s string) bool {
				return s == "a"
			}),
			expected: []string{"a"},
		},
		{
			message: "filtering a slice of ints",
			actual: SliceFilter([]int{1, 2}, func(i int) bool {
				return i == 1
			}),
			expected: []int{1},
		},
		{
			message: "filtering a slice of uints",
			actual: SliceFilter([]uint{uint(1), uint(2)}, func(i uint) bool {
				return i == uint(1)
			}),
			expected: []uint{uint(1)},
		},
		{
			message: "filtering a slice of floats",
			actual: SliceFilter([]float64{float64(1), float64(2)}, func(i float64) bool {
				return i == float64(1)
			}),
			expected: []float64{float64(1)},
		},
		{
			message: "filtering empty strings",
			actual: SliceFilter([]string{"", "a", "b"}, func(s string) bool {
				return s != ""
			}),
			expected: []string{"a", "b"},
		},
	}

	assertions := require.New(t)

	for _, c := range cases {
		assertions.Equal(c.expected, c.actual, "Filtering error, test case: %s", c.message)
	}
}

func TestSliceMap(t *testing.T) {
	cases := []struct {
		message  string
		actual   any
		expected any
	}{
		{
			message: "mapping a slice of strings",
			actual: SliceMap([]string{"a", "b"}, func(s string) string {
				return s + "_mapped"
			}),
			expected: []string{"a_mapped", "b_mapped"},
		},
		{
			message: "mapping a slice of ints",
			actual: SliceMap([]int{1, 2}, func(i int) int {
				return i + 1
			}),
			expected: []int{2, 3},
		},
		{
			message: "mapping a slice of uints",
			actual: SliceMap([]uint{uint(1), uint(2)}, func(i uint) uint {
				return i + 1
			}),
			expected: []uint{uint(2), uint(3)},
		},
		{
			message: "mapping a slice of floats",
			actual: SliceMap([]float64{float64(1), float64(2)}, func(i float64) float64 {
				return i + 0.1
			}),
			expected: []float64{float64(1.1), float64(2.1)},
		},
	}

	assertions := require.New(t)

	for _, c := range cases {
		assertions.Equal(c.expected, c.actual, "Mapping error, test case: %s", c.message)
	}
}

func TestPointerTo(t *testing.T) {
	assertions := require.New(t)
	stringPointer := PtrTo("a")
	acceptStringPointer := func(p *string) string {
		return *p
	}
	stringValue := acceptStringPointer(stringPointer)
	assertions.Equal("a", stringValue)

	PtrTo(true)
	PtrTo("a")
	PtrTo(1)
	PtrTo(int8(1))
	PtrTo(int16(1))
	PtrTo(int32(1))
	PtrTo(int64(1))
	PtrTo(uint(1))
	PtrTo(uint8(1))
	PtrTo(uint16(1))
	PtrTo(uint32(1))
	PtrTo(uint64(1))
	PtrTo(float32(1))
	PtrTo(float64(1))
}

func TestValueOr(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("t", ValueOr("t", "default text"))
	assertions.Equal("default text", ValueOr("", "default text"))
	assertions.Equal("default text", ValueOr("", "", "", "default text", "low priority"))
	assertions.Equal("default text", ValueOr("", "default text", "low priority"))
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
