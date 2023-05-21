package main

import "testing"
import "bytes"

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
GO!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
