package ints

func Tern(cond bool, a, b int) int {
	if cond {
		return a
	}
	return b
}
