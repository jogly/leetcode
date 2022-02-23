package ints

// DedupeInPlace removes duplicate elements from a slice, and returns a slice of
// an equal or smaller size.
//
// SPACE: O(n)
// TIME: O(n)
func DedupeInPlace(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	hash := make(map[int]struct{}, len(a))
	uniq := 0

	for i := range a {
		if _, ok := hash[a[i]]; !ok {
			hash[a[i]] = struct{}{}
			a[uniq] = a[i]
			uniq++
		}
	}
	return a[:uniq]
}
