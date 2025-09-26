package main

import "math/rand"

func newSeed(size int) *[][]bool {
	m := newMatrix(size)
	for range size {
		m[randomInt(size)][randomInt(size)] = true
	}
	return &m
}

func randomInt(limit int) int {
	return rand.Intn(limit)
}
