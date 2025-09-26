package game

import "math/rand"

func NewSeed(n int) *Matrix {
	m := NewMatrix(n)
	for range n {
		m[randomInt(n)][randomInt(n)].SetAlive(true)
	}
	return &m
}

func randomInt(n int) int {
	return rand.Intn(n)
}
