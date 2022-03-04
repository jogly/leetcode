package ints

const maxFactorial = 100

var factorials = make([]int, maxFactorial)

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	if n > maxFactorial {
		panic("factorials table too small")
	}
	if factorials[n] == 0 {
		factorials[n] = n * Factorial(n-1)
	}
	return factorials[n]
}
