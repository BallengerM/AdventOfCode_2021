package day1

import "testing"

func TestPartOne(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	answer := partOne(input)
	if answer != 7 {
		t.Fatalf("Wanted 7, received %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	answer := partTwo(input)
	if answer != 5 {
		t.Fatalf("Wanted 7, received %d", answer)
	}
}
