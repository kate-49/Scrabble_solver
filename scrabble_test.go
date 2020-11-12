package scrabble

import "testing"

func TestLettersHaveCorrectValues(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		// acceptance criteria tests
		{input: []string{"c", "a", "b", "b", "a", "g", "e"}, want: 14},
		{input: []string{""}, want: 0},
		{input: []string{"a"}, want: 1},
		{input: []string{"f"}, want: 4},
		{input: []string{"s", "t", "r", "e", "e", "t"}, want: 6},
		{input: []string{"q", "u", "i", "r", "k", "y"}, want: 22},
		{input: []string{"O", "X", "Y", "P", "H", "E", "N", "B", "U", "T", "A", "Z", "O", "N", "E"}, want: 41},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
