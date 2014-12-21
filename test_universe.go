package main

type TestUniverse struct {
	cell Cell
}

func (t TestUniverse) Apply(r Rules) Universe {
	cell := Cell{
		alive:      r(t.cell),
		neighbours: 0,
	}
	return TestUniverse{cell: cell}
}

func (t TestUniverse) String() string {
	if t.cell.alive {
		return "Hello Alive Universe"
	} else {
		return "Hello Dead Universe"
	}
}

func newTestUniverse(alive bool, neighbours int) *TestUniverse {
	cell := Cell{
		alive:      alive,
		neighbours: neighbours,
	}
	return &TestUniverse{cell: cell}
}
