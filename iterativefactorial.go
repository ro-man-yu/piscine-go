package piscine

import "fmt"

func IterativeFactorial(nb int) int {
	a := 1 //var result int

	for i := 1; i <= nb; i++ {
		a = a * i
	}
	return a
}
