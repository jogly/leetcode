package ints

func Pop(a []int) (int, []int) {
	n := len(a) - 1
	if n == -1 {
		return 0, a
	}
	x := a[n]
	a = a[0:n]
	return x, a
}
