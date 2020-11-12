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
		// {input: []string{"C", "A", "B", "B", "A", "G", "E"}, want: 14},
		// {input: []string{"C", "A", "B", "B", "A", "G", "E"}, want: 14},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
