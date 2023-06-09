package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Flinston", "")
		want := "Hello, Flinston"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when an empty string 	is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Eloide", "Spanish")
		want := "Hola, Eloide"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("Gus", "French")
		want := "Bonjour, Gus"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Portuguese", func(t *testing.T) {
		got := Hello("Vitor", "Portuguese")
		want := "Olá, Vitor"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
