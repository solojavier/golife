package golife

import (
	"bytes"
	"math/rand"
)

// gridUniverse contains a representation of a rectangular universe
type gridUniverse struct {
	height int
	width  int
	grid   [][]bool
}

// NewGridUniverse creates a Universe of the specified size.
// It also generates an initial randome state with ~25% alive cells.
func NewGridUniverse(width, height int) Universe {
	grid := createGrid(width, height)
	for i := 0; i < (width * height / 4); i++ {
		grid[rand.Intn(height)][rand.Intn(width)] = true
	}
	return gridUniverse{
		width:  width,
		height: height,
		grid:   grid,
	}
}

// Apply uses the rules to generate a new generation
// using current universe grid.
func (u gridUniverse) Apply(rules Rules) Universe {
	future := createGrid(u.width, u.height)

	for y := 0; y < u.height; y++ {
		for x := 0; x < u.width; x++ {
			future[y][x] = rules(u.grid[y][x], u.neighbours(x, y))
		}
	}

	u.grid = future
	return u
}

// String represents the Universe as a string.
// Alive cells are marked with an '*'.
func (u gridUniverse) String() string {
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

func (u gridUniverse) neighbours(x, y int) (count int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && u.wrap(x+i, y+j) {
				count++
			}
		}
	}
	return
}

func (u gridUniverse) wrap(x, y int) bool {
	wrapped_x := (x + u.width) % u.width
	wrapped_y := (y + u.height) % u.height

	return u.grid[wrapped_y][wrapped_x]
}

func createGrid(width, height int) [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}
