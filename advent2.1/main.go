package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalDistance, totalDepth := 0, 0

	fetchInput("input2.txt", totalDistance, totalDepth)
}

func fetchInput(path string, x int, y int) {
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

		x, y = updateLocation(direction, amount, x, y)
	}

	fmt.Printf("x: %d y: %d\n", x, y)
	fmt.Printf("total: %d\n", x*y)
}

func updateLocation(direction string, amount int, x int, y int, aim int) (int, int) {
	if direction == "forward" {
		x += amount
		y = aim * amount
	}
	if direction == "up" {
		y -= amount
	}
	if direction == "down" {
		y += amount
	}

	return x, y
}
