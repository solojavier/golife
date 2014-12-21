package main

type TestUniverse struct {
	cell Cell
}

type TestCell struct {
	neighbours int
	alive      bool
}

func (t TestUniverse) Apply(r Rules) Universe {
	cell := &TestCell{
		alive:      r(t.cell),
		neighbours: 0,
	}
	return TestUniverse{cell: cell}
}

func (t TestUniverse) String() string {
	if t.cell.Alive() {
		return "Hello Alive Universe"
	} else {
		return "Hello Dead Universe"
	}
}

func (c TestCell) Alive() bool {
	return c.alive
}

func (c TestCell) LiveNeighbourCount() int {
	return c.neighbours
}
