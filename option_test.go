package gopt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption_Some(t *testing.T) {
	val := 100

	assert.Equal(t, Option[int]{value: &val}, Some(100))
}

func TestOption_None(t *testing.T) {
	assert.Equal(t, Option[int]{value: nil}, None[int]())
}

func TestOption_IsSome(t *testing.T) {
	val := 100
	a := Option[int]{value: &val}

	assert.True(t, a.IsSome())
}

func TestOption_IsNone(t *testing.T) {
	a := Option[int]{value: nil}

	assert.True(t, a.IsNone())
}

func TestOption_Unwrap(t *testing.T) {
	val := 100

	tt := []struct {
		name      string
		input     Option[int]
		expected  int
		wantPanic bool
	}{
		{
			name:     "when has value",
			input:    Option[int]{value: &val},
			expected: 100,
		},
		{
			name:      "when no value",
			input:     Option[int]{value: nil},
			wantPanic: true,
		},
	}

	for _, test := range tt {

		if test.wantPanic {
			assert.Panics(t, func() {
				test.input.Unwrap()
			})
		} else {
			assert.Equal(t, test.expected, test.input.Unwrap())
		}
	}

}

func TestOption_UnwrapOr(t *testing.T) {
	val := 100
	defVal := 1

	tt := []struct {
		name     string
		input    Option[int]
		defInput int
		expected int
	}{
		{
			name:     "when has value",
			input:    Option[int]{value: &val},
			expected: 100,
		},
		{
			name:     "when no value",
			input:    Option[int]{value: nil},
			defInput: defVal,
			expected: 1,
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, test.input.UnwrapOr(test.defInput))
	}

}
