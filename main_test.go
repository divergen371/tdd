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

	t.Run("say 'Hello, World' when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("")
			want := "Hello, World"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
}
