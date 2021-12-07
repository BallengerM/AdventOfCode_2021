package day3

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

	answer1, err := partOne(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day Three, Part One: %d\n", answer1)

	answer2, err := partTwo(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day Two, Part Two: %d\n", answer2)
}

func readInput() ([]string, error) {
	file, err := os.Open("inputs/input_day3.txt")
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

type compare struct {
	ones   int
	zeroes int
}

func partOne(input []string) (int, error) {
	compareMap := make(map[int]compare)
	for _, binary := range input {
		for i, char := range binary {
			current := compareMap[i]
			if char == '1' {
				current.ones++
			} else {
				current.zeroes++
			}
			compareMap[i] = current
		}
	}

	var gamma, epsilon string
	for i := 0; i < len(compareMap); i++ {
		if compareMap[i].ones > compareMap[i].zeroes {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaDecimal, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilonDecimal, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return 0, err
	}

	total := gammaDecimal * epsilonDecimal

	return int(total), nil
}

func partTwo(input []string) (int, error) {
	length := len(input[0])

	var (
		oxygen       = input
		ones, zeroes []string
	)
	for i := 0; i < length; i++ {
		for _, binary := range oxygen {
			if binary[i] == '1' {
				ones = append(ones, binary)
			} else {
				zeroes = append(zeroes, binary)
			}
		}

		switch {
		case len(ones) > len(zeroes) || len(ones) == len(zeroes):
			oxygen = ones
			ones = nil
			zeroes = nil
		case len(ones) < len(zeroes):
			oxygen = zeroes
			ones = nil
			zeroes = nil
		}

		if len(oxygen) == 1 {
			break
		}
	}

	var (
		co2 = input
	)
	for i := 0; i < length; i++ {
		for _, binary := range co2 {
			if binary[i] == '1' {
				ones = append(ones, binary)
			} else {
				zeroes = append(zeroes, binary)
			}
		}

		switch {
		case len(ones) < len(zeroes):
			co2 = ones
			ones = nil
			zeroes = nil
		case len(ones) > len(zeroes) || len(ones) == len(zeroes):
			co2 = zeroes
			ones = nil
			zeroes = nil
		}

		if len(co2) == 1 {
			break
		}
	}

	oxygenDecimal, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		return 0, err
	}

	co2Decimal, err := strconv.ParseInt(co2[0], 2, 64)
	if err != nil {
		return 0, err
	}

	total := oxygenDecimal * co2Decimal

	return int(total), nil
}
