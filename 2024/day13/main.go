package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	var words [][]rune
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// need format like:
		// {A: [11, 49], B: [43, 19], Prize: [15272, 2282]}
	}
}
