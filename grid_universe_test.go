package main

import (
	"strings"
	"testing"
)

func TestSeed(t *testing.T) {
	u := NewGridUniverse(10, 10)

	if !strings.Contains(u.String(), "*") {
		t.Error("Universe should have some alive cells")
	}
}

func TestApply(t *testing.T) {
	u := NewGridUniverse(10, 10)
	nu := u.Apply(func(alive bool, neighbours int) bool {
		return false
	})

	if strings.Contains(nu.String(), "*") {
		t.Error("Universe should have killed all cells")
	}
}

func TestWrap(t *testing.T) {
	var tests = []struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}{
		{0, 0, 10, 0},
		{1, 0, 11, 0},
		{0, 2, 0, 12},
		{0, 4, 0, 14},
		{5, 5, 15, 15},
	}
	u := NewGridUniverse(10, 10).(GridUniverse)

	for _, c := range tests {
		if u.wrap(c.x1, c.y1) != u.wrap(c.x2, c.y2) {
			t.Error(
				"Universe is not wrapped correctly:",
				c.x1, ",", c.y1,
				"should equal:",
				c.x2, ",", c.y2,
			)
		}
	}
}
