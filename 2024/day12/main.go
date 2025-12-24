package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var plants [][]rune
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal("Couldn't open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		plants = append(plants, runes)
	}

	priceSum := 0

	// iterate through rows and cols
	for i := 0; i < len(plants[0]); i++ {
		for j := 0; j < len(plants); j++ {
			// for each plant, find the area and the perimeter and add their product to priceSum
			perimeter := getPerimeterOfPlant(plants[i][j], plants)
			area := getAreaOfRegion(plants[i][j], plants)
			priceSum += perimeter * area
		}
	}

	fmt.Println(priceSum)
}

func getPerimeterOfPlant(plant rune, plantsGrid [][]rune) int {
	totalPerimeter := 4

	// for each direction left, right, up, down, if that rune == plant, totalPerimeter = totalPerimeter - 1
	return totalPerimeter
}

func getAreaOfRegion(plant rune, plantsGrid [][]rune) int {
	totalArea := 0

	return totalArea
}
