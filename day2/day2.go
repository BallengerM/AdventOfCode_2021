package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	horizontal int
	depth      int
	aim        int
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
	fmt.Printf("Day Two, Part One: %d\n", answer1)

	answer2, err := partTwo(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day Two, Part Two: %d\n", answer2)
}

func readInput() ([]string, error) {
	file, err := os.Open("inputs/input_day2.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		scanner    = bufio.NewScanner(file)
		inputSlice []string
	)
	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}

	return inputSlice, scanner.Err()
}

func partOne(input []string) (int, error) {
	starting := coordinates{
		horizontal: 0,
		depth:      0,
	}

	for _, instruction := range input {
		split := strings.Split(instruction, " ")
		direction := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		switch direction {
		case "forward":
			starting.horizontal += amount
		case "up":
			starting.depth -= amount
		case "down":
			starting.depth += amount
		}
	}

	finalAnswer := starting.horizontal * starting.depth

	return finalAnswer, nil
}

func partTwo(input []string) (int, error) {
	starting := coordinates{
		horizontal: 0,
		depth:      0,
	}

	for _, instruction := range input {
		split := strings.Split(instruction, " ")
		direction := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		switch direction {
		case "forward":
			starting.horizontal += amount
			starting.depth += amount * starting.aim
		case "up":
			starting.aim -= amount
		case "down":
			starting.aim += amount
		}
	}

	finalAnswer := starting.horizontal * starting.depth

	return finalAnswer, nil
}
