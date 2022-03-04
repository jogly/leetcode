package ints

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func FprintCarets(w io.Writer, arr []int, carets ...int) {
	if len(arr) == 0 {
		return
	}

	numWidth := MaxWidth(arr)
	sort.Ints(carets)
	var c int
	var b strings.Builder

	for i, a := range arr {
		if i > 0 {
			fatalw(w.Write([]byte{' '}))
			b.WriteByte(' ')
		}
		fmt.Fprintf(w, "%*d", numWidth, a)

		if c < len(carets) && carets[c] == i {
			b.WriteString(fmt.Sprintf("%*s", numWidth, "^"))
			c++
		} else {
			b.WriteString(strings.Repeat(" ", numWidth))
		}
	}
	fatalw(w.Write([]byte{'\n'}))
	fatalw(w.Write([]byte(b.String())))
}

func PrintCarets(arr []int, carets ...int) {
	FprintCarets(os.Stdout, arr, carets...)
}

func Width(a int) int {
	if a == 0 {
		return 1
	}
	if a < 0 {
		return Width(-a) + 1
	}

	var width int
	for a > 0 {
		a /= 10
		width++
	}
	return width
}

func MaxWidth(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var max int
	for _, a := range arr {
		max = Max(max, Width(a))
	}
	return max
}

func fatalw(_ int, e error) {
	fatal(e)
}

func fatal(e error) {
	if e != nil {
		panic(e)
	}
}
