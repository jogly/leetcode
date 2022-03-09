package ints

func Chars(i int) int {
	if i < 0 {
		return Chars(-i) + 1
	}
	if i < 10 {
		return 1
	}
	c := 0
	for ; i > 0; c++ {
		i /= 10
	}
	return c
}
