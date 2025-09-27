package game

import (
	"math/rand"
)

func NewSeed(n int) *Grid {
	m := *NewGrid(n)
	for range randomInt(n * 4) {
		m[randomInt(n)][randomInt(n)].SetAlive(true)
	}
	return &m
}

func randomInt(n int) int {
	return rand.Intn(n)
}
