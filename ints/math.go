package ints

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Pow(x, y int) int {
	if y < 0 {
		return 0
	}
	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		z := Pow(x, y/2)
		return z * z
	}
	return x * Pow(x, y-1)
}
