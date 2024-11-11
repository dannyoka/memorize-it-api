package strategies

import (
	"strings"
)

func EveryNthWord(memorizedString string, n int) string {
	stringArray := strings.Split(memorizedString, " ")
	var newString string
	for i, word := range stringArray {
		if i%n == 0 {
			newString += replaceLettersWithUnderscore(word) + " "
		} else {
			newString += word + " "
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
	for i := 0; i < len(startString); i++ {
		newString += "_"
	}
	return newString
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
