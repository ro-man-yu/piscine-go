package piscine

func IterativeFactorial(nb int) int {
	a := 1 //var result int
	if nb < 0 || nb > 20 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			a = a * i
		}
		return a		
	}	
}
