// Note: to actually run this with `go test`, Go requires the filename to
// end in `_test.go` (e.g. testing_test.go) — that's how `go test` discovers
// test files. It's kept as testing.go here to match the repo layout; rename
// it if you want to run it directly.
package main

import (
	"fmt"
	"testing"
)

// Add is the function under test.
func Add(a, b int) int {
	return a + b
}

// Divide returns an error instead of panicking on divide-by-zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

// TestAdd is a basic single-case test.
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

// TestAddTableDriven shows the idiomatic table-driven pattern.
func TestAddTableDriven(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -1, -2},
		{"zero", 0, 0, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

// TestDivide shows asserting on both a value and an error.
func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err) // Fatalf stops immediately
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("expected an error dividing by zero, got nil")
	}
}

// BenchmarkAdd measures Add's performance; run with `go test -bench=.`
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func main() {
	fmt.Println("This file's real purpose is its Test* and Benchmark* functions.")
	fmt.Println("Rename to testing_test.go and run: go test -v ./...")
}
