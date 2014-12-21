package main

import "fmt"

// Rules defines function to be applied to a cell.
type Rules func(Cell) bool

// Universe interface contains a representation of cells
type Universe interface {
	//Apply uses Rules function to generate a new universe
	//by applying it to every cell it contains
	Apply(Rules) Universe
	fmt.Stringer
}

// Cell is a living being inside a Universe
type Cell interface {
	LiveNeighbourCount() int
	Alive() bool
}

// Tick applies Rules function to the Universe
// to create a new generation.
func Tick(past Universe, rules Rules) Universe {
	return past.Apply(rules)
}

// Conway defines default rules for GoL
func Conway(c Cell) bool {
	if c.Alive() {
		if c.LiveNeighbourCount() == 2 || c.LiveNeighbourCount() == 3 {
			return true
		}
	} else {
		if c.LiveNeighbourCount() == 3 {
			return true
		}
	}
	return false
}
