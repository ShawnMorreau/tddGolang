package iterating

import (
	"fmt"
	"testing"
)

func TestIteratate(t *testing.T) {
	assertEquals := func(t testing.TB, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected: %q, got %q", expected, actual)
		}
	}
	t.Run("iterate and return a 5 times", func(t *testing.T) {
		assertEquals(t, "aaaaa", Iterate("a", 5))
	})
	t.Run("iterate with multiple letters twice", func(t *testing.T) {
		assertEquals(t, "baba", Iterate("ba", 2))
	})

}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("abc", 5)
	}
}

func ExampleIterate() {
	res := Iterate("a", 4)
	fmt.Println(res)
	// Output: aaaa
}
