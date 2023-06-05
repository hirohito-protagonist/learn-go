package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://google.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}
