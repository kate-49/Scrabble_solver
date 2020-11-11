package scrabble

func Run(inputLetters []string) int {
	var sum int

	OneValueLetters := [10]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	TwoValueLetters := [2]string{"D", "G"}

	for _, letter := range inputLetters {
		for _, oneValueLetter := range OneValueLetters {
			if letter == oneValueLetter {
				sum++
			}
		}
		for _, TwoValueLetter := range TwoValueLetters {
			if letter == TwoValueLetter {
				sum += 2
			}
		}
	}

	return sum
}
