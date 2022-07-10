package gopt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult_Ok(t *testing.T) {
	assert.Equal(t, Result[int]{ok: 100, err: nil}, Ok(100))
}

func TestResult_Err(t *testing.T) {
	assert.Equal(t, Result[int]{ok: 0, err: assert.AnError}, Err[int](assert.AnError))
}

func TestResult_IsOk(t *testing.T) {
	a := Result[int]{ok: 100, err: nil}

	assert.True(t, a.IsOk())
	assert.False(t, a.IsErr())
}

func TestResult_IsErr(t *testing.T) {
	a := Result[int]{ok: 0, err: assert.AnError}

	assert.False(t, a.IsOk())
	assert.True(t, a.IsErr())
}

func TestResult_Unwrap(t *testing.T) {
	tt := []struct {
		name      string
		input     Result[int]
		expected  int
		wantPanic bool
	}{
		{
			name:     "with Ok",
			input:    Result[int]{ok: 100, err: nil},
			expected: 100,
		},
		{
			name:      "when no value",
			input:     Result[int]{ok: 0, err: assert.AnError},
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

type A struct {
	Name string
	B    Result[B]
}

type B struct {
	Name string
}

func TestResult(t *testing.T) {
	barResult := Result[B]{ok: B{Name: "Bar"}, err: nil}

	tt := []struct {
		name      string
		result    Result[B]
		input     A
		output    B
		wantPanic bool
	}{
		{
			name:   "with Ok",
			result: Result[B]{ok: B{Name: "Bar"}, err: nil},
			input:  A{Name: "Foo", B: barResult},
			output: B{Name: "Bar"},
		},
		{
			name:      "without Err",
			result:    Result[B]{err: assert.AnError},
			input:     A{Name: "Foo", B: Result[B]{err: assert.AnError}},
			wantPanic: true,
		},
	}

	for _, test := range tt {
		if test.wantPanic {
			assert.False(t, test.input.B.IsOk())
			assert.True(t, test.input.B.IsErr())
			assert.Panics(t, func() { test.input.B.Unwrap() })
		} else {
			assert.True(t, test.input.B.IsOk())
			assert.False(t, test.input.B.IsErr())
			assert.Equal(t, test.input.B.Unwrap(), test.output)
		}
	}

}
