package walker

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with string and int field",
			struct {
				Name string
				Age  int
			}{"Chris", 55},
			[]string{"Chris"},
		},
		{
			"nested struct",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Using a pointer",
			&Person{
				"Chris",
				Profile{22, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"What if it's a slice",
			[]Profile{
				{44, "overhere"},
				{23, "overthere"},
			},
			[]string{"overhere", "overthere"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{22, "hello"},
			},
			[]string{"London", "hello"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"baz": "box",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "bar")
		assertContains(t, got, "box")
	})

}
func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
