package hello

import "testing"

func TestGreeting(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to Bob in English", func(t *testing.T) {
		got := Greeting("Bob", "English")
		want := "Hello, Bob"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the World with unknown language", func(t *testing.T) {
		got := Greeting("", "ldnwainlwa")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world in a different language", func(t *testing.T) {
		got := Greeting("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Bob in a different language", func(t *testing.T) {
		got := Greeting("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}
