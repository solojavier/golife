package golife

import "fmt"

// Rules defines function that can be used to determine
// a cell next state.
// Returns true if should be alive and false if it should die.
type Rules func(bool, int) bool

// Universe interface contains a representation of cells
type Universe interface {
	//Apply uses Rules function to generate a new universe
	//by applying it to every cell it contains
	Apply(Rules) Universe
	fmt.Stringer
}

// Conway defines default rules for GoL
func Conway(alive bool, neighbours int) bool {
	if alive {
		if neighbours == 2 || neighbours == 3 {
			return true
		}
	} else {
		if neighbours == 3 {
			return true
		}
	}
	return false
}

// Tick applies Rules function to the Universe
// to create a new generation.
func Tick(past Universe, rules Rules) Universe {
	return past.Apply(rules)
}
