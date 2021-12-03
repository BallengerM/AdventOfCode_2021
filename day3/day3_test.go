package day3

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	answer, err := partOne(input)
	if err != nil {
		t.Error(err)
	}
	if answer != 198 {
		t.Fatalf("Wanted 198, received %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	answer, err := partTwo(input)
	if err != nil {
		t.Error(err)
	}
	if answer != 230 {
		t.Fatalf("Wanted 230, received %d", answer)
	}
}
