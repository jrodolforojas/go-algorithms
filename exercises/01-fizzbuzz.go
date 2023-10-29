package exercises

import (
	"strconv"
)

func (exercises *Exercises) Fizzbuzz(list []int) []string {
	results := []string{}
	for _, item := range list {
		fizz := item%3 == 0
		buzz := item%5 == 0

		if fizz && buzz {
			results = append(results, "fizzbuzz")
		} else if fizz {
			results = append(results, "fizz")
		} else if buzz {
			results = append(results, "buzz")
		} else {
			results = append(results, strconv.Itoa(item))
		}
	}
	return results
}
