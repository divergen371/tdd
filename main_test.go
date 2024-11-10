package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Atsushi")
	want := "Hello, Atsushi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
