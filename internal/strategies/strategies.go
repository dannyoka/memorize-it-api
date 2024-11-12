package strategies

import (
	"strings"
)

func EveryNthWord(memorizedString string, n int) string {
	stringArray := strings.Split(memorizedString, " ")
	var newString string
	for i, word := range stringArray {
		if i%n == 0 {
			newString += word + " "
		} else {
			newString += replaceLettersWithUnderscore(word) + " "
		}
	}
	return newString
}

func FirstLetterOfEveryWord(memorizedString string) string {
	stringArray := strings.Split(memorizedString, " ")
	var newString string
	for _, word := range stringArray {
		newString += replaceAllButFirstLetterWithUnderscore(word) + " "
	}
	return newString
}

func replaceLettersWithUnderscore(startString string) string {
	var newString string
	// for i := 0; i < len(startString); i++ {
	// 	newString += "_"
	// }
	for _, letter := range startString {
		if checkIfCharacterIsLetter(letter) {
			newString += "_"
		} else {
			newString += string(letter)
		}
	}
	return newString
}

func checkIfCharacterIsLetter(character rune) bool {
	if (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') {
		return true
	}
	return false
}

func replaceAllButFirstLetterWithUnderscore(startString string) string {
	var newString string
	for i, letter := range startString {
		if i == 0 {
			newString += string(letter)
		} else {
			newString += "_"
		}
	}
	return newString
}
