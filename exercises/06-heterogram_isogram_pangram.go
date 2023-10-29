package exercises

import (
	"strings"
)

func (exe *Exercises) getCharacterCounter(text string) map[string]int {
	characterCounter := make(map[string]int)

	formattedText := strings.ReplaceAll(text, " ", "") // remove white spaces
	formattedText = strings.ToLower(formattedText)     // all characters in lower case

	for i := 0; i < len(formattedText); i++ {
		letter := string(formattedText[i])
		characterCounter[letter] = characterCounter[letter] + 1
	}

	return characterCounter
}

func (exe *Exercises) isHeterogram(text string) bool {
	textCounter := exe.getCharacterCounter(text)
	for _, value := range textCounter {
		if value > 1 {
			return false
		}
	}
	return true
}

func (exe *Exercises) isIsogram(text string) bool {
	order := 0
	for _, value := range exe.getCharacterCounter(text) {
		if order == 0 {
			order = value
		}
		if order != value {
			return false
		}
	}
	return true
}

func (exe *Exercises) isPangram(text string) bool {
	alphabet := "qwertyuiopasdfghjklzxcvbnm"
	alphabetCounter := exe.getCharacterCounter(alphabet)
	textCounter := exe.getCharacterCounter(text)

	return len(textCounter) == len(alphabetCounter)
}
