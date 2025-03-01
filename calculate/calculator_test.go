package calculate_test

import (
	"testing"

	"github.com/ragg967/GoCalculator/calculate" // Updated import path
)

func TestAdd(t *testing.T) {
	result := calculate.Add(3, 5) // Updated package name
	expected := 8.0
	if result != expected {
		t.Errorf("Add(3, 5) = %f; want %f", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := calculate.Subtract(10, 4) // Updated package name
	expected := 6.0
	if result != expected {
		t.Errorf("Subtract(10, 4) = %f; want %f", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := calculate.Multiply(3, 5) // Updated package name
	expected := 15.0
	if result != expected {
		t.Errorf("Multiply(3, 5) = %f; want %f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		result, err := calculate.Divide(10, 2) // Updated package name
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		expected := 5.0
		if result != expected {
			t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calculate.Divide(10, 0) // Updated package name
		if err == nil {
			t.Error("Divide(10, 0) expected an error, got nil")
		}
	})
}
