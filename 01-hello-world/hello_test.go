package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sayng hello to people", func(t *testing.T) {
		got := Hello("João", "")
		want := "Hello, João"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("João", "Spanish")
		want := "Hola, João"
		assertCorrectMessage(t, want, got)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("João", "Portuguese")
		want := "Olá, João"
		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
