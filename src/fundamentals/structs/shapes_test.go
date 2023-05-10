package shapes

import "testing"

func TestPerimiter(t *testing.T) {
	got := Perimiter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}
