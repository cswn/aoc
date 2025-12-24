package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var connections [][]string
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal("Couldn't open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// filter out connections that don't start with t
		//if line[0] == 116 || line[3] == 116 {
		firstComputer := string(line[0]) + string(line[1])
		secondComputer := string(line[3]) + string(line[4])
		computers := []string{firstComputer, secondComputer}
		connections = append(connections, computers)
		//}
	}

	result := checkForSetsOfThree(connections)

	fmt.Println(result)
	fmt.Println(len(result))
}

func checkForSetsOfThree(connections [][]string) [][]string {
	var result [][]string

	for i := 0; i < len(connections); i++ {
		curr := connections[i]
		set := []string{curr[0], curr[1]}
		for _, pair := range connections {
			var matchedItem string
			var thirdInSet string

			// skip over curr
			if pair[0] == set[0] && pair[1] == set[1] {
				continue
			}

			if pair[0] == set[0] {
				matchedItem = pair[0]
				thirdInSet = pair[1]
			} else if pair[0] == set[1] {
				matchedItem = pair[0]
				thirdInSet = pair[1]
			} else if pair[1] == set[0] {
				matchedItem = pair[1]
				thirdInSet = pair[0]
			} else if pair[1] == set[1] {
				matchedItem = pair[1]
				thirdInSet = pair[0]
			} else {
				continue
			}

			// searchFor should be the one in curr that isnt matchedItem
			var searchFor string
			if curr[0] == matchedItem {
				searchFor = curr[1]
			} else if curr[1] == matchedItem {
				searchFor = curr[0]
			}

			// fmt.Println(curr)
			// fmt.Println(matchedItem)
			// fmt.Println(searchFor)
			// fmt.Println(thirdInSet)

			makesASet := checkIfSetOfThreeIsMade(connections, searchFor, thirdInSet)

			if makesASet {
				set = append(set, thirdInSet)

				// if set is already in result, don't append
				setInResult := checkIfSetIsAlreadyInResult(set, result)
				if !setInResult {
					result = append(result, set)
				}
				break
			}
		}
	}

	return result
}

func checkIfSetOfThreeIsMade(rowsToCheck [][]string, item1 string, item2 string) bool {
	found := false
	for _, row := range rowsToCheck {
		if (row[0] == item1 || row[1] == item1) && (row[0] == item2 || row[1] == item2) {
			found = true
		}
	}
	return found
}

func checkIfSetIsAlreadyInResult(set []string, result [][]string) bool {
	found := false
	for _, row := range result {
		if (row[0] == set[0] || row[1] == set[0] || row[2] == set[0]) && (row[0] == set[1] || row[1] == set[1] || row[2] == set[1]) && (row[0] == set[2] || row[1] == set[2] || row[2] == set[2]) {
			found = true
		}
	}
	return found
}
