package main

import (
	"bytes"
	"math/rand"
)

// GridUniverse contains a representation of a rectangular universe
type GridUniverse struct {
	height int
	width  int
	grid   [][]bool
}

// NewGridUniverse creates a Universe of the specified size.
// It also generates an initial randome state with ~25% alive cells.
func NewGridUniverse(w, h int) Universe {
	grid := createGrid(w, h)
	for i := 0; i < (w * h / 4); i++ {
		grid[rand.Intn(h)][rand.Intn(w)] = true
	}
	return GridUniverse{
		width:  w,
		height: h,
		grid:   grid,
	}
}

// Apply uses the rules to generate a new generation
// using current universe grid.
func (u GridUniverse) Apply(rules Rules) Universe {
	future := createGrid(u.width, u.height)

	for y := 0; y < u.height; y++ {
		for x := 0; x < u.width; x++ {
			future[y][x] = rules(u.grid[y][x], u.aliveNeighbours(x, y))
		}
	}

	u.grid = future
	return u
}

// String represents the Universe as a string.
// Alive cells are marked with an '*'.
func (u GridUniverse) String() string {
	var buf bytes.Buffer
	for y := 0; y < u.height; y++ {
		for x := 0; x < u.width; x++ {
			b := byte(' ')
			if u.grid[y][x] {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (u GridUniverse) aliveNeighbours(x, y int) (n int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && u.wrap(x+i, y+j) {
				n++
			}
		}
	}
	return
}

func (u GridUniverse) wrap(x, y int) bool {
	wrapped_x := (x + u.width) % u.width
	wrapped_y := (y + u.height) % u.height

	return u.grid[wrapped_y][wrapped_x]
}

func createGrid(w, h int) [][]bool {
	g := make([][]bool, h)
	for i := range g {
		g[i] = make([]bool, w)
	}
	return g
}
