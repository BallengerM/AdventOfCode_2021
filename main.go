package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"git.bisonline.com/makayla.ballenger/advent/day1"
	"git.bisonline.com/makayla.ballenger/advent/day2"
	"git.bisonline.com/makayla.ballenger/advent/day3"
	"git.bisonline.com/makayla.ballenger/advent/day5"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which day would you like the solution to?: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.TrimSpace(text)

	switch {
	case text == "1" || text == "One":
		day1.Run()
	case text == "2" || text == "Two":
		day2.Run()
	case text == "3" || text == "Three":
		day3.Run()
	case text == "4" || text == "Four":
		fmt.Println("Day 4 has not been completed.")
	case text == "5" || text == "Five":
		day5.Run()
	default:
		fmt.Println("The requested day is invalid.")
	}
}
