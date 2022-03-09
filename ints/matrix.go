package ints

import "fmt"

type Matrix [][]int

func NewMatrix(width, height int) Matrix {
	rows := make(Matrix, height)
	for i := range rows {
		rows[i] = make([]int, width)
	}
	return rows
}

func (m Matrix) SetAll(v int) {
	for y := range m {
		for x := range m[y] {
			m[y][x] = v
		}
	}
}

func (m Matrix) SumNeighbors(x, y int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			skipX := x+j < 0 || x+j >= len(m[0])
			skipY := y+i < 0 || y+i >= len(m)
			skipCenter := i == 0 && j == 0
			if skipX || skipY || skipCenter {
				continue
			}

			sum += m[y+i][x+j]
		}
	}
	return sum
}

func (m Matrix) AccumNeighbors(x, y int, f func(acc, n int) int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			skipX := x+j < 0 || x+j >= len(m[0])
			skipY := y+i < 0 || y+i >= len(m)
			skipCenter := i == 0 && j == 0
			if skipX || skipY || skipCenter {
				continue
			}

			sum = f(sum, m[y+i][x+j])
		}
	}
	return sum
}

func (m Matrix) CheckBit(x, y, mask int) bool {
	return m[y][x]&mask != 0
}

func (m Matrix) Print() {
	for y := range m {
		for x := range m[y] {
			fmt.Printf("%d ", m[y][x])
		}
		fmt.Println()
	}
}
