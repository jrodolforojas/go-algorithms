package exercises

import "fmt"

func (exe *Exercises) primoFibonacciPar(number int) string {
	isEven := number%2 == 0

	isPrime := true
	for i := 2; i < number; i++ {
		if number%i != 0 {
			isPrime = false
		}
	}

	isFibonacci := true
	a := 0
	b := 1
	for a < number {
		a = b
		b = a + b
	}

	isFibonacci = a == number

	return fmt.Sprintf("%d %t %t %t", number, isPrime, isFibonacci, isEven)
}
