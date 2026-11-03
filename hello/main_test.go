package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("chris", "English")
		want := "Hello, chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hola' + name when Spanish is passed as a parameter", func(t *testing.T) {
		got := Hello("e", "Spanish")
		want := "Hola, e"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say, 'Bonjour' + name when Spanish is passed as a parameter", func(t *testing.T) {
		got := Hello("e", "French")
		want := "Bonjour, e"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)

	}
}
