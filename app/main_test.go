package app

import (
	"fmt"
	"testing"
)

// TestSuite for Add function
func TestAddSuite(t *testing.T) {
	t.Run("Positive Numbers", func(t *testing.T) {
		result := Add(5, 3)
		expected := 8
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Negative Numbers", func(t *testing.T) {
		result := Add(-1, -2)
		expected := -3
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Zero Case", func(t *testing.T) {
		result := Add(0, 0)
		expected := 0
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}

// TestSuite for Multiply function
func TestMultiplySuite(t *testing.T) {
	t.Run("Positive Numbers", func(t *testing.T) {
		result := Multiply(5, 3)
		expected := 15
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Negative Numbers", func(t *testing.T) {
		result := Multiply(-1, -2)
		expected := 2
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Zero Case", func(t *testing.T) {
		result := Multiply(0, 3)
		expected := 0
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}
