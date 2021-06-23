package adder

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertEqual := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("Add two numbers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		assertEqual(t, sum, expected)
	})
	t.Run("add two negative numbers", func(t *testing.T) {
		sum := Add(-2, -3)
		expected := -5
		assertEqual(t, sum, expected)
	})
	t.Run("nothing sent in args", func(t *testing.T) {
		sum := Add()
		expected := 0
		assertEqual(t, sum, expected)
	})
	t.Run("add many numbers", func(t *testing.T) {
		sum := Add(-2, -3, 5, 2, 1, 2)
		expected := 5
		assertEqual(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 2, 3)
	fmt.Println(sum)
	// Output: 6
}
