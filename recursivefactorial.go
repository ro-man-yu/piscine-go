package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb > 1 && nb < 20 {
		return RecursiveFactorial(nb-1) * nb
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	arg := 19
	fmt.Println(RecursiveFactorial(arg))
}