package main

type TestUniverse struct {
	alive      bool
	neighbours int
}

func (t TestUniverse) Apply(r Rules) Universe {
	return TestUniverse{alive: r(t.alive, t.neighbours)}
}

func (t TestUniverse) String() string {
	if t.alive {
		return "Hello Alive Universe"
	} else {
		return "Hello Dead Universe"
	}
}

func newTestUniverse(alive bool, neighbours int) *TestUniverse {
	return &TestUniverse{
		alive:      alive,
		neighbours: neighbours,
	}
}
