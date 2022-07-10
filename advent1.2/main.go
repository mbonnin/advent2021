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
	fmt.Println(findDepthByRange(data))
}

func fetchInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func findDepthByRange(depths []int) int {
	numIncreasingRanges := 0
	for i := 3; i < len(depths); i++ {
		currentRange := depths[i] + depths[i-1] + depths[i-2]
		previousRange := depths[i-1] + depths[i-2] + depths[i-3]
		if currentRange > previousRange {
			numIncreasingRanges++
		}
	}
	return numIncreasingRanges
}
