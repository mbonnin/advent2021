package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	depth    int
	distance int
	aim      int
}

func main() {
	var loc Location

	fetchInput("input2.txt", loc)
}

func fetchInput(path string, loc Location) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parsedLine := strings.Fields(scanner.Text())
		direction := parsedLine[0]
		amount, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			loc.distance = loc.distance + amount
			loc.depth = loc.depth + (loc.aim * amount)
		case "up":
			loc.aim = loc.aim - amount
		case "down":
			loc.aim = loc.aim + amount

		}
	}

	fmt.Printf("total: %d\n", loc.distance*loc.depth)
}
