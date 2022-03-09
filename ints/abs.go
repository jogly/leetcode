package ints

func Abs(a int) int {
	return Tern(a < 0, -a, a)
}
