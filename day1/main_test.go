package main

import "testing"

func TestSumUpTheElfCalories(t *testing.T) {
	elfCalories := []int{0, 20, 30}
	want := 50

	got := sumUpTheElfCalories(elfCalories)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
