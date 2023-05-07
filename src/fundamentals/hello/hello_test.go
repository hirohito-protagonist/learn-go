package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hiro")
	want := "Hello, Hiro"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
