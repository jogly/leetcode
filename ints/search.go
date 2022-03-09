package ints

func SearchAny(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		} else if a[m] > x {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

func SearchFirst(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		} else if a[m] > x {
			r = m - 1
		} else {
			if m == 0 || a[m-1] != x {
				return m
			}
			r = m - 1
		}
	}
	return -1
}

func SearchLast(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		} else if a[m] > x {
			r = m - 1
		} else {
			if m == len(a)-1 || a[m+1] != x {
				return m
			}
			l = m + 1
		}
	}
	return -1
}

func SearchClosest(a []int, x int) int {
	if len(a) == 0 {
		return -1
	}
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] < x {
			l = m + 1
		} else if a[m] > x {
			r = m - 1
		} else {
			return m
		}
	}
	// l = r+1
	if x < a[l] {
		return r
	}
	return l
}
