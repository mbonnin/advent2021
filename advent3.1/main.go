package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	gamma, epsilon := scanFile("input3.txt")

	fmt.Println(binaryToInteger(gamma) * binaryToInteger(epsilon))
}

func scanFile(fileName string) ([]int, []int) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digitTally := make([]int, 12)
	var numRows int
	for scanner.Scan() {
		for i, c := range scanner.Text() {
			if string(c) == "1" {
				digitTally[i]++
			}
		}
		numRows++
	}

	gamma := make([]int, 12)
	epsilon := make([]int, 12)

	// fmt.Println(digitTally)
	// fmt.Println(numRows)

	mid := numRows / 2

	for i, n := range digitTally {
		if n > mid {
			gamma[i] = 1
		} else {
			epsilon[i] = 1
		}
	}

	return gamma, epsilon

}

func binaryToInteger(binaryString []int) int {
	length := len(binaryString) - 1
	var total float64

	for _, d := range binaryString {
		if d == 1 {
			total += math.Pow(2, float64(length))
		}
		length--
	}

	return int(total)
}
