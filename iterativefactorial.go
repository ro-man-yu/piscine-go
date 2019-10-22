package piscine

func IterativeFactorial(nb int) int {
<<<<<<< HEAD
	a := 1 //var result int
	if nb < 0 || nb > 20 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			a = a * i
		}
		return a		
	}	
=======
	a := 1

	for i := 1; i <= nb; i++ {
		a = a * i
	}
	return a
>>>>>>> b5f0687b93445af35e79b71fbaa3dcd21924cd88
}
