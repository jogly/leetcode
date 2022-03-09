package ints

type PQI interface {
	Priority() int
}

// A PQ is a min-heap of PQI.
type PQ []PQI

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
func (h PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PQ) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PQI))
}

func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}
