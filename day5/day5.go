package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	XPoint int
	YPoint int
}

func Run() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	answer1, err := partOne(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day Three, Part One: %d\n", answer1)
}

func readInput() (map[int][]coordinate, error) {
	file, err := os.Open("inputs/input_day5.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		scanner  = bufio.NewScanner(file)
		inputMap = make(map[int][]coordinate)
		counter  int
	)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), "->")

		var coordinatePair []coordinate
		for _, coord := range coords {
			trimmed := strings.TrimSpace(coord)
			separatedCoords := strings.Split(trimmed, ",")

			xy := coordinate{}
			xy.XPoint, err = strconv.Atoi(separatedCoords[0])
			if err != nil {
				return nil, err
			}

			xy.YPoint, err = strconv.Atoi(separatedCoords[1])
			if err != nil {
				return nil, err
			}

			coordinatePair = append(coordinatePair, xy)
		}

		inputMap[counter] = coordinatePair
		counter++
	}

	return inputMap, scanner.Err()
}

func partOne(input map[int][]coordinate) (int, error) {
	return 0, nil
}
