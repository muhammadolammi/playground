package math

import (
	"testing"
)

func TestStandardForm(t *testing.T) {
	testCases := []struct {
		name     string
		in       float64
		expected string
	}{
		{
			name:     "test for whole number",
			in:       float64(80000),
			expected: "8 * 10^4",
		},
		{
			name:     "test for decimal point number",
			in:       float64(0.0008),
			expected: "8 * 10^-4",
		},
		{
			name:     "test for normal  number",
			in:       float64(96900),
			expected: "8 * 10^-4",
		},
		{
			name:     "test for large number",
			in:       2 << 11,
			expected: "3 * 10^153",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			result := StandardForm(tc.in)
			if result != tc.expected {
				t.Errorf("StandardForm(%f) = %s. expected: %s", tc.in, result, tc.expected)
			}

		})
	}

}

func TestNeedInt(t *testing.T) {
	testCases := []struct {
		name      string
		in        float64
		expected  int
		expectErr bool
	}{
		{
			name:      "success",
			in:        10,
			expected:  11,
			expectErr: false,
		},
		{
			name:      "failure",
			in:        1 << 100,
			expected:  0,
			expectErr: true,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			result, err := NeedInt(tc.in)
			if (err != nil) != tc.expectErr {
				t.Errorf("NeedInt(%v) error = %v, expectErr %v", tc.in, err, tc.expectErr)
				return
			}
			if result != tc.expected {
				t.Errorf("NeedInt(%v) = %v, expected: %v", tc.in, result, tc.expected)
			}

		})
	}
}
