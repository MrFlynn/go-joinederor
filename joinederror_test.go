package joinederror

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var equateErrors = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})

type testCase struct {
	name       string
	inputError error
	expected   []error
}

func Test_UnwrapMany(t *testing.T) {
	testCases := []testCase{
		{
			name:       "single_error",
			inputError: errors.Join(errors.New("lorem")),
			expected:   []error{errors.New("lorem")},
		},
		{
			name:       "single_error_with_nil",
			inputError: errors.Join(errors.New("lorem"), nil),
			expected:   []error{errors.New("lorem")},
		},
		{
			name:       "two_errors",
			inputError: errors.Join(errors.New("lorem"), errors.New("ipsum")),
			expected:   []error{errors.New("lorem"), errors.New("ipsum")},
		},
		{
			name:       "nil_error",
			inputError: nil,
			expected:   nil,
		},
		{
			name:       "single_nil",
			inputError: errors.Join(nil),
			expected:   nil,
		},
		{
			name:       "multiple_nil",
			inputError: errors.Join(nil, nil),
			expected:   nil,
		},
		{
			name:       "non_joined_error",
			inputError: errors.New("lorem ipsum"),
			expected:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			diff := cmp.Diff(tc.expected, UnwrapMany(tc.inputError), equateErrors)

			if diff != "" {
				t.Errorf("Error slice mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_UnwrapAll(t *testing.T) {
	testCases := []testCase{
		{
			name:       "nil",
			inputError: nil,
			expected:   []error{},
		},
		{
			name:       "basic_error",
			inputError: errors.New("hello world"),
			expected:   []error{errors.New("hello world")},
		},
		{
			name: "nested_simple",
			inputError: errors.Join(
				errors.New("lorem"),
				errors.Join(errors.New("ipsum"), errors.New("dolor")),
			),
			expected: []error{
				errors.New("lorem"), errors.New("ipsum"), errors.New("dolor"),
			},
		},
		{
			name: "double_nested",
			inputError: errors.Join(
				errors.New("lorem"),
				errors.Join(
					errors.New("ipsum"),
					errors.New("dolor"),
					errors.Join(errors.New("sit"), errors.New("amet")),
				),
			),
			expected: []error{
				errors.New("lorem"),
				errors.New("ipsum"),
				errors.New("dolor"),
				errors.New("sit"),
				errors.New("amet"),
			},
		},
		{
			name: "nested_complex",
			inputError: errors.Join(
				errors.New("lorem"),
				errors.Join(errors.New("ipsum"), errors.New("dolor")),
				errors.New("sit"),
				errors.New("amet"),
				errors.Join(
					errors.New("consectetur"),
					errors.New("adipiscing"),
					errors.Join(errors.New("elit"), errors.New("sed")),
				),
			),
			expected: []error{
				errors.New("lorem"),
				errors.New("ipsum"),
				errors.New("dolor"),
				errors.New("sit"),
				errors.New("amet"),
				errors.New("consectetur"),
				errors.New("adipiscing"),
				errors.New("elit"),
				errors.New("sed"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			diff := cmp.Diff(tc.expected, UnwrapAll(tc.inputError), equateErrors)

			if diff != "" {
				t.Errorf("Error slice mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
