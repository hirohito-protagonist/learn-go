package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hiro")

	got := buffer.String()
	want := "Hello, Hiro"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
