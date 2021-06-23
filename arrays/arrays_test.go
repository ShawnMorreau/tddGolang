package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertEqual := func(t testing.TB, expected, actual int) {
		if expected != actual {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	}
	t.Run("sum a bunch of numbers", func(t *testing.T) {
		arr := []int{1, 2, 3, 5, 2}
		assertEqual(t, 13, Sum(arr))
	})
	t.Run("empty", func(t *testing.T) {
		arr := []int{}
		assertEqual(t, 0, Sum(arr))
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 3})
	expected := []int{3, 7}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
