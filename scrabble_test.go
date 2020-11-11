package scrabble

import "testing"

func TestLettersWithValue1IncludedNonValueLetters(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input: []string{"A", "E", "I", "O", "U"}, want: 5},
		{input: []string{"A", "B", "C", "E", "I", "K", "O", "X"}, want: 4},
		{input: []string{"L", "N", "R", "S", "T"}, want: 5},
		{input: []string{"N", "E", "U", "X"}, want: 3},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
