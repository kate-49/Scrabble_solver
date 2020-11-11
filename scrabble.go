package scrabble

func Run(inputLetters []string) int {
	var sum int

	OneValueLetters := [10]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}

	for _, letter := range inputLetters {
		for _, oneValueLetter := range OneValueLetters {
			if letter == oneValueLetter {
				sum++
			}
		}
	}

	return sum
}
