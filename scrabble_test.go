package scrabble

import "testing"

func TestLettersHaveCorrectValues(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input: []string{}, want: 0},
		//one value letters
		{input: []string{"A", "E", "I", "O", "U"}, want: 5},
		{input: []string{"L", "N", "R", "S", "T"}, want: 5},
		//two value letters
		{input: []string{"D", "G"}, want: 4},
		//three value letters
		{input: []string{"B", "C", "M", "P"}, want: 12},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
