package day2

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	answer, err := partOne(input)
	if err != nil {
		t.Error(err)
	}
	if answer != 150 {
		t.Fatalf("Wanted 150, received %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	answer, err := partTwo(input)
	if err != nil {
		t.Error(err)
	}
	if answer != 900 {
		t.Fatalf("Wanted 900, received %d", answer)
	}
}
