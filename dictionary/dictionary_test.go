package dictionary

import "testing"

func TestSearch(t *testing.T) {
	key := "test"
	value := "this is a test"
	dictionary := Dictionary{key: value}

	t.Run("in dictionary", func(t *testing.T) {
		got, err := dictionary.Search(key)
		want := value
		assertString(t, got, want)
		assertNoError(t, err)
	})
	t.Run("not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("shimshamblimblam")
		assertError(t, err, ErrInvalidKey)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"abcd": "exists"}

	t.Run("Word not in dictionary", func(t *testing.T) {
		err := dictionary.Add("word", "phrase")
		got, _ := dictionary.Search("word")
		assertNoError(t, err)
		assertString(t, got, "phrase")
	})

	t.Run("Word in dictionary", func(t *testing.T) {
		err := dictionary.Add("abcd", "exists")
		assertError(t, err, ErrKeyAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"abcd": "exists"}
	t.Run("Updates word", func(t *testing.T) {
		err := dictionary.Update("abcd", "newPhrase")
		got, _ := dictionary.Search("abcd")
		assertString(t, got, "newPhrase")
		assertNoError(t, err)
	})
	t.Run("Update called with invalid key", func(t *testing.T) {
		err := dictionary.Update("shimablima", "nope")
		assertError(t, err, ErrInvalidKey)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"abcd": "lkjdwa"}
	t.Run("Key deleted", func(t *testing.T) {
		dictionary.Delete("abcd")
		_, got := dictionary.Search("abcd")

		assertError(t, got, ErrInvalidKey)
	})
}
func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Wanted no error but got %q", err)
	}
}
