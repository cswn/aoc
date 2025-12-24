package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	diskMap := string(fileContent)
	compacted := compactFile(diskMap)

	compactedArray := stringToIntArray(compacted)
	checkSum := calculateCheckSum(compactedArray)

	fmt.Println(checkSum)
}

func compactFile(diskMap string) []string {
	var representedFileBlocks []string
	diskMapSlice := strings.Split(diskMap, "")

	for i, fileNumber := range diskMapSlice {
		fileNumberInt, _ := strconv.Atoi(fileNumber)

		if i%2 == 0 {
			id := i / 2
			for j := 0; j < fileNumberInt; j++ {
				representedFileBlocks = append(representedFileBlocks, strconv.Itoa(id))
			}
		} else { // for odd index, add periods to represent free space
			for j := 0; j < fileNumberInt; j++ {
				representedFileBlocks = append(representedFileBlocks, ".")
			}
		}
	}

	representedFileBlocksStr := shiftBlocksFromEndToLeft(representedFileBlocks)
	return representedFileBlocksStr
}

func shiftBlocksFromEndToLeft(blocks []string) []string {
	result := make([]string, 0)
	files := 0
	for _, block := range blocks {
		if block != "." {
			files++
		}
	}

	for i, block := range blocks {
		if i == files {
			break
		}

		if block == "." {
			// if block is a period, replace it in the result with the last el in blocks
			for i := 1; i < files; i++ {
				lastItem := blocks[len(blocks)-i]
				if lastItem != "." {
					lastItem := string(lastItem)
					blocks = remove(blocks, len(blocks)-i)
					blocks = append(blocks, ".")

					result = append(result, lastItem)
					break
				}
			}
		} else {
			result = append(result, block)
		}
	}
	return result
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func calculateCheckSum(file []int) int {
	checkSum := 0
	for i, val := range file {
		sum := i * val
		checkSum += sum
	}
	return checkSum
}

func stringToIntArray(str []string) []int {
	var nums []int
	for _, val := range str {
		if val == "." {
			continue
		}
		num, _ := strconv.Atoi(val)
		nums = append(nums, num)
	}
	return nums
}
