package main

import (
	"fmt"
	"time"

	"github.com/solojavier/golife"
)

func main() {
	universe := golife.NewGridUniverse(80, 20)
	evolution(universe, 0)
}

func evolution(universe golife.Universe, generation int) {
	fmt.Println("Generation", generation, ":")
	fmt.Print(universe)
	time.Sleep(1 * time.Second)
	evolution(golife.Tick(universe, golife.Conway), generation+1)
}
