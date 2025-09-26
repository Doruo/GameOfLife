package main

import "fmt"

func newMatrix(size int) [][]bool {
	m := make([][]bool, size)
	for i := range m {
		m[i] = make([]bool, size)
	}
	return m
}

func showMatrix(m [][]bool) {
	for i := range m {
		for j := range m {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println(" ")
	}
}
