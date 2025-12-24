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

		crossedZero := countRotationCrossedZero(rotationCount, dialPosition, rotationDirection)

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

		if crossedZero > 0 {
			fmt.Printf(" Also, the dial crossed zero %v times.", crossedZero)
		}

		countZeroes += crossedZero
	}

	return countZeroes
}

func countRotationCrossedZero(count int, startPos int, dir string) int {
	// switch dir {
	// case "R":
	// 	if ((startPos + count) % 100) > 99 {
	// 		fmt.Printf("\n v The dial was rotated %v.\n", (startPos+count)%100)

	// 		return 1
	// 	}
	// case "L":
	// 	if ((startPos - count) % 100) < 0 {
	// 		fmt.Printf("\n v The dial was rotated %v.\n", (startPos-count)%100)

	// 		return 1
	// 	}
	// }

	countCrossedZero := 0
	switch dir {
	case "R":
		newPos := startPos + count
		if newPos > 100 {
			countCrossedZero += newPos % 100
		}
	case "L":
		newPos := startPos - count
		if newPos < 0 {
			countCrossedZero += newPos % 100
		}
	}

	return countCrossedZero
}
