package main

import (
	"bytes"
	"math/rand"
)

type GridUniverse struct {
	height int
	width  int
	grid   [][]bool
}

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

func (u GridUniverse) Apply(r Rules) Universe {
	future := createGrid(u.width, u.height)

	for y := 0; y < u.height; y++ {
		for x := 0; x < u.width; x++ {
			cell := Cell{
				alive:      u.grid[y][x],
				neighbours: u.aliveNeighbours(x, y),
			}
			future[y][x] = r(cell)
		}
	}

	u.grid = future

	return u
}

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
			if (j != 0 || i != 0) && u.alive(x+i, y+j) {
				n++
			}
		}
	}
	return
}

func (u GridUniverse) alive(x, y int) bool {
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
