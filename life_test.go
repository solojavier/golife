package main

import "testing"

func newTestUniverse(state bool, neighbours int) *TestUniverse {
	cell := &TestCell{
		alive:      state,
		neighbours: neighbours,
	}
	return &TestUniverse{cell: cell}
}

var tests = []struct {
	state      bool
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
		u := Tick(newTestUniverse(c.state, c.neighbours), Conway).(TestUniverse)

		if u.cell.Alive() != c.expected {
			t.Error(
				"for (", "neighbours:", c.neighbours, "state:", c.state, ")",
				"expected:", c.expected,
				"got:", u.cell.Alive(),
			)
		}
	}
}
