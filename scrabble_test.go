package scrabble

import "testing"

func TestScrabble(t *testing.T) {

	t.Run("Test that letter A returns 1", func(t *testing.T) {
		got := Run("A")
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
}
