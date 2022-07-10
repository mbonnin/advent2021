package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := fetchInput("input1.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(findIncreasingDepths(data))
}

func fetchInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func findIncreasingDepths(depths []string) int {
	numIncreasingDepths := 0
	for i := 1; i < len(depths); i++ {
		depth, err := strconv.Atoi(depths[i])
		if err != nil {
			panic(err)
		}
		prevDepth, err := strconv.Atoi(depths[i-1])
		if err != nil {
			panic(err)
		}
		if depth > prevDepth {
			numIncreasingDepths++
		}
	}

	return numIncreasingDepths
}
