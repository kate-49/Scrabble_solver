package scrabble

import "testing"

func TestLettersWithValue1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "A", want: 1},
		{input: "E", want: 1},
		{input: "I", want: 1},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
