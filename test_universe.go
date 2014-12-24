package golife

type testUniverse struct {
	alive      bool
	neighbours int
}

func (t testUniverse) Apply(r Rules) Universe {
	return testUniverse{alive: r(t.alive, t.neighbours)}
}

func (t testUniverse) String() string {
	if t.alive {
		return "Hello Alive Universe"
	} else {
		return "Hello Dead Universe"
	}
}

func newTestUniverse(alive bool, neighbours int) *testUniverse {
	return &testUniverse{
		alive:      alive,
		neighbours: neighbours,
	}
}
