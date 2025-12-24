package main

import (
	"bufio"
	"fmt"
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
		runes := []rune(line)
		words = append(words, runes)
	}

	//count := findCountOfXmasInGrid(words, "XMAS")
	count := findCountOfXmasInGrid(words, "MAS")

	fmt.Println(count)
}

func findCountOfXmasInGrid(grid [][]rune, word string) int {
	sum := 0
	totalRows := len(grid)
	totalInRow := len(grid[0])
	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalInRow; j++ {
			runes := []rune(word)
			//count := search2DArray(grid, i, j, runes)
			count := search2DArrayForCross(grid, i, j, runes)
			sum = sum + count
		}
	}

	return sum
}

func search2DArray(grid [][]rune, row int, col int, word []rune) int {
	totalRows := len(grid)
	totalInRow := len(grid[0])

	// first, check if first letter of given coord matches with first letter of word
	if grid[row][col] != word[0] {
		return 0
	}

	len := len(word)

	// coordinates to search in
	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	foundCount := 0
	// search in all 8 possible directions
	for dir := 0; dir < 8; dir++ {

		// initialize starting points
		currentXPos := row + x[dir]
		currentYPos := col + y[dir]
		var k int

		for k = 1; k < len; k++ {
			// out of bounds
			if currentXPos >= totalRows || currentXPos < 0 || currentYPos >= totalInRow || currentYPos < 0 {
				break
			}

			if grid[currentXPos][currentYPos] != word[k] {
				break
			}

			currentXPos += x[dir]
			currentYPos += y[dir]
		}

		// if all characters matched, then k must = length of word at this poiunt
		if k == len {
			foundCount++
		}
	}

	return foundCount
}

func search2DArrayForCross(grid [][]rune, row int, col int, word []rune) int {
	totalRows := len(grid)
	totalInRow := len(grid[0])

	// first, check if first letter of given coord matches with first letter of word
	if grid[row][col] != word[1] {
		return 0
	}

	len := len(word)

	// coordinates to search in
	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	foundCount := 0
	// search in all 8 possible directions
	for dir := 0; dir < 8; dir++ {

		// initialize starting points
		currentXPos := row + x[dir]
		currentYPos := col + y[dir]
		var k int

		for k = 1; k < len; k++ {
			// out of bounds
			if currentXPos >= totalRows || currentXPos < 0 || currentYPos >= totalInRow || currentYPos < 0 {
				break
			}

			if grid[currentXPos][currentYPos] != word[k] {
				break
			}

			currentXPos += x[dir]
			currentYPos += y[dir]
		}

		// if all characters matched, then k must = length of word at this poiunt
		if k == len {
			foundCount++
		}
	}

	return foundCount
}
