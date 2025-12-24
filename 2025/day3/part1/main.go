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

	var batteryBanks []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		batteryBanks = append(batteryBanks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	totalJoltage := 0
	for _, bank := range batteryBanks {
		maxJoltage := findMaxJoltage(bank)
		fmt.Printf("max joltage is %v\n", maxJoltage)

		totalJoltage += maxJoltage
	}

	fmt.Printf("\ntotal output joltage is %v", totalJoltage)
}

func findMaxJoltage(bank string) int {
	largestDigitIndex, largestDigit := findLargest(bank, 0)
	secondLargestDigit := -1

	for i := largestDigitIndex; i < len(bank); i++ {
		if i == largestDigitIndex {
			break
		}

		secondBatteryInt, _ := strconv.Atoi(string(bank[i]))

		if secondBatteryInt > secondLargestDigit {
			secondLargestDigit = secondBatteryInt
		}
	}

	// for i, battery := range bank {
	// 	if i == largestDigitIndex {
	// 		continue
	// 	}

	// 	batteryInt, _ := strconv.Atoi(string(battery))
	// 	if batteryInt > secondLargestDigit {

	// 		secondLargestDigit = batteryInt
	// 		secondLargestIndex = i
	// 	}
	// }

	firstStr := strconv.Itoa(largestDigit)
	secondStr := strconv.Itoa(secondLargestDigit)
	maxJoltage := firstStr + secondStr

	maxJoltageInt, _ := strconv.Atoi(maxJoltage)
	return maxJoltageInt
}

func findLargest(bank string, startIndex int) (int, int) {
	largest, _ := strconv.Atoi(string(bank[0]))
	index := 0

	for i := startIndex; i < len(bank); i++ {
		fmt.Printf("iteration #%v looking at battery value %v. largest val is currently %v\n", i, string(bank[i]), largest)

		if i == index {
			continue
		}

		batteryInt, _ := strconv.Atoi(string(bank[i]))
		if batteryInt > largest {
			largest = batteryInt
			index = i
		}
	}

	return index, largest
}
