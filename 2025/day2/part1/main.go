package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	productIds := strings.Split(strings.TrimSpace(string(data)), ",")
	totalInvalidIds := 0
	for _, idRange := range productIds {
		ids := strings.Split(idRange, "-")

		startIdInt, _ := strconv.Atoi(ids[0])
		endIdInt, _ := strconv.Atoi(ids[1])

		for id := startIdInt; id <= endIdInt; id++ {
			s := strconv.Itoa(id)

			isInvalidId := checkForInvalidId(s)

			if isInvalidId {
				totalInvalidIds += id
			}
		}
	}

	fmt.Printf("Adding up all the invalid IDs produces %v\n", totalInvalidIds)
}

func checkForInvalidId(idStr string) bool {
	if len(idStr)%2 != 0 {
		return false
	}

	indexToSplit := len(idStr) / 2

	firstHalf := idStr[:indexToSplit]
	secondHalf := idStr[indexToSplit:]

	if firstHalf == secondHalf {
		return true
	}

	return false
}
