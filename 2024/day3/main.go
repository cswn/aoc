package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzleInput := string(fileContent)

	r := regexp.MustCompile(`(mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\))`)
	mulInstructions := r.FindAllString(puzzleInput, -1)

	var filtedInstructions []string
	for _, instruction := range mulInstructions {
		//exclude mul functions that are preceeded by dont()
		r := regexp.MustCompile(`don't\(\)`)
		dont := r.FindAllString(instruction, -1)
		fmt.Println(dont)
		filtedInstructions = append(filtedInstructions, instruction)
	}

	sum := 0
	for _, instruction := range filtedInstructions {
		// get the numbers from the mul function
		r := regexp.MustCompile(`([0-9]+)`)
		numbersToMultiply := r.FindAllString(instruction, -1)

		var ints []int
		for _, i := range numbersToMultiply {
			j, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, j)
		}
		sum = sum + (ints[0] * ints[1])
	}

	fmt.Println(sum)
}
