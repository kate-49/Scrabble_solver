package scrabble

func Run(inputLetters []string) int {
	var sum int

	OneValueLetters := [10]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	TwoValueLetters := [2]string{"D", "G"}
	ThreeValueLetters := [4]string{"B", "C", "M", "P"}
	FourValueLetters := [5]string{"F", "H", "V", "W", "Y"}

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
		for _, ThreeValueLetter := range ThreeValueLetters {
			if letter == ThreeValueLetter {
				sum += 3
			}
		}
		for _, FourValueLetter := range FourValueLetters {
			if letter == FourValueLetter {
				sum += 4
			}
		}
	}

	return sum
}
