package scrabble

import "testing"

func TestScrabble(t *testing.T) {

	t.Run("Test that 5 letters of 1 value letters return 5", func(t *testing.T) {
		got := Run([]string{"A", "E", "I", "O", "U"})
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
	t.Run("Test mix of 1 value letters and other letters", func(t *testing.T) {
		got := Run([]string{"A", "B", "C", "E", "I", "K", "O", "X"})
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
}

// func TestLettersWithValue1(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  int
// 	}{
// 		{input: "A", want: 1},
// 		{input: "E", want: 1},
// 		{input: "I", want: 1},
// 		{input: "0", want: 1},
// 		{input: "U", want: 1},
// 		{input: "L", want: 1},
// 		{input: "N", want: 1},
// 		{input: "R", want: 1},
// 		{input: "S", want: 1},
// 		{input: "T", want: 1},
// 	}

// 	for _, tc := range tests {
// 		got := Run(tc.input)
// 		if tc.want != got {
// 			t.Errorf("expected: %v, got: %v", tc.want, got)
// 		}
// 	}
// }
