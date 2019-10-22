package piscine

func IterativeFactorial(nb int) int {
	0 < nb <= 20
	a := 1 //var result int

	for i := 1; i <= nb; i++ {
		a = a * i
	}
	return a
}
