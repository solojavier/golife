package main

import (
	"fmt"
	"time"
)

func main() {
	universe := NewGridUniverse(80, 20)
	evolution(universe, 0)
}

func evolution(universe Universe, generation int) {
	fmt.Println("Generation ", generation, ":")
	fmt.Print(universe)
	time.Sleep(1 * time.Second)
	evolution(Tick(universe, Conway), generation+1)
}
