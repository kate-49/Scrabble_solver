package scrabble

import "testing"

func TestLettersHaveCorrectValues(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		// acceptance criteria tests
		{input: "cabbage", want: 14},
		{input: "", want: 0},
		{input: "nil", want: 0},
		{input: "a", want: 1},
		{input: "f", want: 4},
		{input: "street", want: 6},
		{input: "quirky", want: 22},
		{input: "OXYPHENBUTAZONE", want: 41},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
