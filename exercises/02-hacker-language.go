package exercises

func (exercises *Exercises) HackerLanguage(input string) string {
	hackerLanguage := ""
	for index := range input {
		letter := string(input[index])
		if letter == "a" {
			hackerLanguage += "4"
		} else if letter == "e" {
			hackerLanguage += "3"
		} else if letter == "i" {
			hackerLanguage += "1"
		} else if letter == "o" {
			hackerLanguage += "0"
		} else {
			hackerLanguage += letter
		}
	}
	return hackerLanguage
}
