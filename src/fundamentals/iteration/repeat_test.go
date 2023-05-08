package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
