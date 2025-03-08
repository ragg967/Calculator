package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		shouldFail bool
	}{
		{"2 * (3 + 3)", 12, false},
		{"5 ^ 2", 25, false},
		{"5 * 2", 10, false},
		{"10 / 2", 5, false},
		{"5 + 5", 10, false},
		{"10 - 5", 5, false},
		{"25 % 2", 5, false},
		{"invalid expression", 0, true},
	}

	for _, test := range tests {
		result, err := Calculate(test.expression)
		if test.shouldFail {
			if err == nil {
				t.Errorf("Expected error for expression %q, but got none", test.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for expression %q: %v", test.expression, err)
			}
			if result != test.expected {
				t.Errorf("For expression %q, expected %v, but got %v", test.expression, test.expected, result)
			}
		}
	}
}
