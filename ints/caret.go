package ints

import (
	"bytes"
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
	var b bytes.Buffer

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
	fatalw(w.Write(b.Bytes()))
}

func SprintCarets(arr []int, carets ...int) string {
	w := &strings.Builder{}
	FprintCarets(w, arr, carets...)
	return w.String()
}

func PrintCarets(arr []int, carets ...int) {
	FprintCarets(os.Stdout, arr, carets...)
}

func MaxWidth(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var max int
	for _, a := range arr {
		max = Max(max, Chars(a))
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
