package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileRules, err := os.Open("rules.txt")
	if err != nil {
		log.Fatal("Couldn't open input file")
	}
	defer fileRules.Close()

	var rules [][]int
	rulesScanner := bufio.NewScanner(fileRules)
	for rulesScanner.Scan() {
		line := rulesScanner.Text()
		slice := strings.Split(line, "|")

		rule := make([]int, len(slice))
		for i, str := range slice {
			rule[i], _ = strconv.Atoi(str)
		}
		rules = append(rules, rule)
	}

	file, err := os.Open("pages.txt")
	if err != nil {
		log.Fatal("Couldn't open input file")
	}
	defer file.Close()

	var middleNumbers []int
	var middleNumbersIncorrect []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arrayLine := convertToIntArray(line)
		isCorrect := checkForCorrectOrder(arrayLine, rules)
		if isCorrect {
			middleIndex := int(math.Round(float64(len(arrayLine) / 2)))
			middleNum := arrayLine[middleIndex]
			middleNumbers = append(middleNumbers, middleNum)
		} else {
			arr := reorderPageNumbers(arrayLine, rules)
			middleIndex := int(math.Round(float64(len(arr) / 2)))
			middleNum := arr[middleIndex]
			middleNumbersIncorrect = append(middleNumbersIncorrect, middleNum)
		}
	}

	sum := 0
	for _, val := range middleNumbers {
		sum += val
	}
	sumIncorrect := 0
	for _, val := range middleNumbersIncorrect {
		sumIncorrect += val
	}

	fmt.Println(sum)
	fmt.Println(sumIncorrect)
}

func checkForCorrectOrder(pageNums []int, rules [][]int) bool {
	isCorrect := true

	// iterate through rules to check the page numbers against each rule pair
	for _, rulePair := range rules {
		firstNum := rulePair[0]
		secondNum := rulePair[1]
		firstNumIndex := -1
		secondNumIndex := -1

		for i := 0; i < len(pageNums); i++ {
			if pageNums[i] == firstNum {
				firstNumIndex = i
			}

			if pageNums[i] == secondNum {
				secondNumIndex = i
			}
		}

		if firstNumIndex != -1 && secondNumIndex != -1 && firstNumIndex > secondNumIndex {
			isCorrect = false
		}
	}

	return isCorrect
}

func convertToIntArray(s string) []int {
	stringSlice := strings.Split(s, ",")
	result := make([]int, len(stringSlice))

	for i, str := range stringSlice {
		result[i], _ = strconv.Atoi(str)
	}

	return result
}

func reorderPageNumbers(pageNums []int, rules [][]int) []int {
	result := make([]int, len(pageNums))
	copy(result, pageNums)

	for !checkForCorrectOrder(result, rules) {
		for _, rulePair := range rules {
			firstNum := rulePair[0]
			secondNum := rulePair[1]
			firstNumIndex := -1
			secondNumIndex := -1

			for i := 0; i < len(result); i++ {
				if result[i] == firstNum {
					firstNumIndex = i
				}

				if result[i] == secondNum {
					secondNumIndex = i
				}
			}

			if firstNumIndex != -1 && secondNumIndex != -1 && firstNumIndex > secondNumIndex {
				firstNumInLine := result[firstNumIndex]
				secondNumInLine := result[secondNumIndex]

				result[firstNumIndex] = secondNumInLine
				result[secondNumIndex] = firstNumInLine
			}
		}
	}

	return result
}
