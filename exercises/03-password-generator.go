package exercises

import "math/rand"

func (exe *Exercises) passwordGenerator(length int, allowMayus bool, allowNumbers bool, allowSymbols bool) string {
	allowedCharacters := "qwertyuiopasdfghjklñzxcvbnm"
	mayus := "QWERTYUIOPASDFGHJKLÑZXCVBNM"
	numbers := "1234567890"
	symbols := "!#$%&/()=.-_,+*;:"

	password := ""
	if length < 8 || length > 16 {
		return password
	}

	if allowMayus {
		allowedCharacters += mayus
	}
	if allowNumbers {
		allowedCharacters += numbers
	}
	if allowSymbols {
		allowedCharacters += symbols
	}

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(allowedCharacters))
		randomCharacter := string(allowedCharacters[randomIndex])
		password += randomCharacter
	}

	return password
}
