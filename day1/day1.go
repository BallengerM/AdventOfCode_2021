package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	answer1 := partOne(input)
	fmt.Printf("Day One, Part One: %d\n", answer1)

	answer2 := partTwo(input)
	fmt.Printf("Day One, Part Two: %d\n", answer2)
}

func readInput() ([]int, error) {
	file, err := os.Open("inputs/input_day1.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		scanner    = bufio.NewScanner(file)
		inputSlice []int
	)
	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}

		inputSlice = append(inputSlice, num)
	}

	return inputSlice, scanner.Err()
}

func partOne(input []int) int {
	var previousNum, increasedAmount int
	for i := range input {
		if input[i] > previousNum && i != 0 {
			increasedAmount++
		}

		previousNum = input[i]
	}

	return increasedAmount
}

func partTwo(input []int) int {
	var previousSum, increasedAmount int
	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		if sum > previousSum && i != 0 {
			increasedAmount++
		}

		previousSum = sum
	}

	return increasedAmount
}
