package main

import (
	"testing"

	"github.com/AndersonBargas/aoc2022/shape"
)

func TestMatchForThePartOne(t *testing.T) {
	want := 4
	got := matchForThePartOne(&shape.Rock{}, &shape.Rock{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 8
	got = matchForThePartOne(&shape.Rock{}, &shape.Paper{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 3
	got = matchForThePartOne(&shape.Rock{}, &shape.Scissors{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 1
	got = matchForThePartOne(&shape.Paper{}, &shape.Rock{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 5
	got = matchForThePartOne(&shape.Paper{}, &shape.Paper{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 9
	got = matchForThePartOne(&shape.Paper{}, &shape.Scissors{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 7
	got = matchForThePartOne(&shape.Scissors{}, &shape.Rock{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 2
	got = matchForThePartOne(&shape.Scissors{}, &shape.Paper{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	want = 6
	got = matchForThePartOne(&shape.Scissors{}, &shape.Scissors{})
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
