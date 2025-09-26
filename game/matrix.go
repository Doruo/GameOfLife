package game

import (
	"fmt"

	"github.com/doruo/gameoflife/cell"
)

type Matrix [][]cell.Cell

func NewMatrix(size int) Matrix {
	m := make(Matrix, size)
	for i := range m {
		m[i] = make([]cell.Cell, size)
	}
	return m
}

func Show(m Matrix) {
	for i := range m {
		for j := range m {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println(" ")
	}
}
