package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"git.bisonline.com/makayla.ballenger/advent/day1"
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
	default:
		fmt.Println("The requested day is invalid.")
	}
}
