package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	safeRowCount := 0
	for _, row := range data {
		//isRowSafe := checkForRowSafety(row)
		isRowMinusElementSafe := checkForRowSafetyMinusElements(row)

		if isRowMinusElementSafe {
			safeRowCount = safeRowCount + 1
		}
	}

	fmt.Println(safeRowCount)
}

func checkForRowSafety(rowInt []int) bool {
	// first, convert to numbers
	// var rowInt []int
	// for _, num := range row {
	// 	j, err := strconv.Atoi(num)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rowInt = append(rowInt, j)
	// }

	// are all nums either increasing or decreasing?
	for i := 0; i < len(rowInt)-1; i++ {
		shouldBeIncreasing := rowInt[0] < rowInt[1]

		if shouldBeIncreasing {
			if rowInt[i] > rowInt[i+1] {
				return false
			}
		} else {
			if rowInt[i] < rowInt[i+1] {
				return false
			}
		}
	}

	// are any two adjacent levels differing by n, where 1>=n>=3?
	for i := 0; i < len(rowInt)-1; i++ {
		item := rowInt[i]
		nextItem := rowInt[i+1]
		diff := math.Abs(float64(item) - float64(nextItem))
		if diff < 1 {
			return false
		} else if diff > 3 {
			return false
		}
	}

	return true
}

func checkForRowSafetyMinusElements(row []string) bool {
	// first, convert to numbers
	var rowInt []int
	for _, num := range row {
		j, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		rowInt = append(rowInt, j)
	}

	isOneVariationTrue := checkForRowSafety(rowInt)
	if !isOneVariationTrue {
		for i := 0; i < len(rowInt); i++ {
			copy := makeCopy(rowInt)

			row := removeByIndex(copy, i)

			// if any one of them is true, return true
			if checkForRowSafety(row) {
				isOneVariationTrue = true
			}
		}
	}

	return isOneVariationTrue
}

func removeByIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func makeCopy(array []int) []int {
	result := make([]int, len(array))

	_ = copy(result, array)
	return result
}
