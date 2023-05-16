package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		got := Search(dictionary, "test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})
}
