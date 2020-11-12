package scrabble

import "strings"

func Run(inputLetters string) int {
	var sum int

	type ScrabbleLetters struct {
		LetterValue []string
	}

	one := ScrabbleLetters{
		LetterValue: []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	}
	two := ScrabbleLetters{
		LetterValue: []string{"D", "G"},
	}
	three := ScrabbleLetters{
		LetterValue: []string{"B", "C", "M", "P"},
	}
	four := ScrabbleLetters{
		LetterValue: []string{"F", "H", "V", "W", "Y"},
	}
	five := ScrabbleLetters{
		LetterValue: []string{"K"},
	}
	eight := ScrabbleLetters{
		LetterValue: []string{"J", "X"},
	}
	ten := ScrabbleLetters{
		LetterValue: []string{"Q", "Z"},
	}

	if inputLetters == "nil" {
		return sum
	}

	inputLettersArray := strings.SplitAfterN(inputLetters, "", len(inputLetters))

	for _, letter := range inputLettersArray {
		for _, oneValueLetter := range one.LetterValue {
			if strings.ToUpper(letter) == oneValueLetter {
				sum++
			}
		}
		for _, TwoValueLetter := range two.LetterValue {
			if strings.ToUpper(letter) == TwoValueLetter {
				sum += 2
			}
		}
		for _, ThreeValueLetter := range three.LetterValue {
			if strings.ToUpper(letter) == ThreeValueLetter {
				sum += 3
			}
		}
		for _, FourValueLetter := range four.LetterValue {
			if strings.ToUpper(letter) == FourValueLetter {
				sum += 4
			}
		}
		for _, FiveValueLetter := range five.LetterValue {
			if strings.ToUpper(letter) == FiveValueLetter {
				sum += 5
			}
		}
		for _, EightValueLetter := range eight.LetterValue {
			if strings.ToUpper(letter) == EightValueLetter {
				sum += 8
			}
		}
		for _, TenValueLetter := range ten.LetterValue {
			if strings.ToUpper(letter) == TenValueLetter {
				sum += 10
			}
		}
	}

	return sum
}
