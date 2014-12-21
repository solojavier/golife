package main

import (
	"fmt"
	"time"
)

type Rules func(Cell)

type Universe interface {
	Apply(Rules) Universe
	Rules() Rules
	fmt.Stringer
}

type Cell interface {
	LiveNeighbourCount() int
	Alive() bool
	Born()
	Die()
}

func Play(universe Universe) {
	fmt.Println(universe)
	time.Sleep(1)
	Play(tick(universe))
}

func tick(past Universe) Universe {
	return past.Apply(past.Rules())
}
