package ints

func NCR(n, r int) int {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}

func NPR(n, r int) int {
	return Factorial(n) / Factorial(n-r)
}
