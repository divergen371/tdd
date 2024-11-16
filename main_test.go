package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	asserCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Atsushi", "")
		want := "Hello, Atsushi"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"
			asserCorrectMessage(t, got, want)
		})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		asserCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Atsushi", "French")
		want := "Bonjour, Atsushi"
		asserCorrectMessage(t, got, want)
	})
}
