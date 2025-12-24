package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../test_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rotations []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	password := findPassword(rotations)
	fmt.Printf("\nThe password is %v", password)
}

func findPassword(rotations []string) int {
	dialPosition := 50
	countZeroes := 0

	for _, rotation := range rotations {
		rotationDirection := string(rotation[0])
		rotationCountStr := rotation[1:]
		rotationCount, err := strconv.Atoi(rotationCountStr)
		if err != nil {
			panic(err)
		}

		switch rotationDirection {
		case "R":
			newPos := dialPosition + rotationCount
			if newPos > 99 {
				newPos = (newPos - 100)
			}
			dialPosition = newPos
		case "L":
			newPos := dialPosition - rotationCount
			if newPos < 0 {
				newPos = (100 + newPos)
			}
			dialPosition = newPos
		}
		fmt.Printf("\nThe dial is rotated %v to point at %v.", rotation, dialPosition)

		if dialPosition%100 == 0 {
			countZeroes++
		}
	}

	return countZeroes
}
