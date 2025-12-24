package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
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

	var list1 []int
	var list2 []int
	for _, row := range data {
		j, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, j)

		k, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		list2 = append(list2, k)
	}

	// sort the two lists
	sort.Ints(list1[:])
	sort.Ints(list2[:])

	length := len(list1)
	var pairs [][]int

	for i := 0; i < length; i++ {
		pairs = append(pairs, []int{})
		pairArr := []int{}
		pairArr = append(pairArr, list1[i], list2[i])
		pairs[i] = pairArr
	}
	similarityScore := calculateSimilarityScore(pairs, list2)
	fmt.Println(similarityScore)

	// part 1 code below
	// sum := 0
	// for _, pair := range pairs {
	// 	var distance int
	// 	if pair[0] < pair[1] {
	// 		distance = pair[1] - pair[0]
	// 	} else {
	// 		distance = pair[0] - pair[1]
	// 	}

	// 	sum = sum + distance
	// }

	// fmt.Println(sum)
}

func calculateSimilarityScore(pairs [][]int, haystack []int) int {
	similarityScore := 0
	for _, pair := range pairs {
		product := numAppearances(pair[0], haystack) * pair[0]
		similarityScore = similarityScore + product
	}

	return similarityScore
}

func numAppearances(needle int, haystack []int) int {
	appearances := 0
	for _, item := range haystack {
		if item == needle {
			appearances = appearances + 1
		}
	}
	return appearances
}
