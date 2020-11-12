package scrabble

import "strings"

func Run(inputLetters string) int {
	var sum int

	OneValueLetters := [10]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	TwoValueLetters := [2]string{"D", "G"}
	ThreeValueLetters := [4]string{"B", "C", "M", "P"}
	FourValueLetters := [5]string{"F", "H", "V", "W", "Y"}
	FiveValueLetters := [1]string{"K"}
	EightValueLetters := [2]string{"J", "X"}
	TenValueLetters := [2]string{"Q", "Z"}

	arrayLength := len(inputLetters)
	inputLettersArray := strings.SplitAfterN(inputLetters, "", arrayLength)

	for _, letter := range inputLettersArray {
		for _, oneValueLetter := range OneValueLetters {
			if strings.ToUpper(letter) == oneValueLetter {
				sum++
			}
		}
		for _, TwoValueLetter := range TwoValueLetters {
			if strings.ToUpper(letter) == TwoValueLetter {
				sum += 2
			}
		}
		for _, ThreeValueLetter := range ThreeValueLetters {
			if strings.ToUpper(letter) == ThreeValueLetter {
				sum += 3
			}
		}
		for _, FourValueLetter := range FourValueLetters {
			if strings.ToUpper(letter) == FourValueLetter {
				sum += 4
			}
		}
		for _, FiveValueLetter := range FiveValueLetters {
			if strings.ToUpper(letter) == FiveValueLetter {
				sum += 5
			}
		}
		for _, EightValueLetter := range EightValueLetters {
			if strings.ToUpper(letter) == EightValueLetter {
				sum += 8
			}
		}
		for _, TenValueLetter := range TenValueLetters {
			if strings.ToUpper(letter) == TenValueLetter {
				sum += 10
			}
		}
	}

	return sum
}
