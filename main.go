package main

import (
	"fmt"
	"time"
)

func main() {
	universe := NewGridUniverse(80, 20)
	n := 0
	for {
		fmt.Println("Generation ", n, ":")
		fmt.Print(universe)
		universe = Tick(universe, Conway)
		time.Sleep(1 * time.Second)
		n++
	}
}
