package main

import "testing"

func newTestUniverse(alive bool, neighbours int) *TestUniverse {
	cell := Cell{
		alive:      alive,
		neighbours: neighbours,
	}
	return &TestUniverse{cell: cell}
}

var tests = []struct {
	alive      bool
	neighbours int
	expected   bool
}{
	{true, 0, false},
	{true, 1, false},
	{true, 2, true},
	{true, 3, true},
	{true, 4, false},
	{false, 0, false},
	{false, 1, false},
	{false, 2, false},
	{false, 3, true},
	{false, 4, false},
}

func TestTick(t *testing.T) {
	for _, c := range tests {
		u := Tick(newTestUniverse(c.alive, c.neighbours), Conway).(TestUniverse)

		if u.cell.alive != c.expected {
			t.Error(
				"for (", "neighbours:", c.neighbours, "alive:", c.alive, ")",
				"expected:", c.expected,
				"got:", u.cell.alive,
			)
		}
	}
}
