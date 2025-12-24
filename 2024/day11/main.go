package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "2 54 992917 5270417 2514 28561 0 990"
	inputArr := strings.Fields(input)
	//stones := make([]int, len(inputArr))
	stones := make(map[int]int)
	for _, str := range inputArr {
		//stones[i], _ = strconv.Atoi(str)
		num, _ := strconv.Atoi(str)
		stones[num]++
	}

	// for i := 0; i < 75; i++ {
	// 	var newStones []int
	// 	for _, stone := range stones {
	// 		lenNum := getNumDigits(stone)
	// 		if stone == 0 {
	// 			newStones = append(newStones, 1)
	// 		} else if lenNum%2 == 0 {
	// 			firstHalf, secondHalf := cutIntInHalf(stone, lenNum)
	// 			newStones = append(newStones, firstHalf)
	// 			newStones = append(newStones, secondHalf)
	// 		} else {
	// 			newStones = append(newStones, (stone * 2024))
	// 		}
	// 	}
	// 	stones = newStones
	// }

	for i := 0; i < 75; i++ {
		updatedStones := make(map[int]int)
		for stone, count := range stones {
			lenNum := getNumDigits(stone)
			if stone == 0 {
				updatedStones[1] += count
			} else if lenNum%2 == 0 {
				firstHalf, secondHalf := cutIntInHalf(stone, lenNum)
				updatedStones[firstHalf] += count
				updatedStones[secondHalf] += count
			} else {
				updatedStones[stone*2024] += count
			}
		}
		stones = updatedStones
	}
	sum := 0
	for _, count := range stones {
		sum += count
	}

	fmt.Println(stones)
	fmt.Println(sum)
	//fmt.Println(len(stones))
}

func getNumDigits(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func cutIntInHalf(num int, length int) (int, int) {
	str := strconv.Itoa(num)
	halfLen := length / 2
	firstHalf := str[0:halfLen]
	secondHalf := str[halfLen:]

	firstHalfStr, _ := strconv.Atoi(firstHalf)
	secondHalfStr, _ := strconv.Atoi(secondHalf)

	return firstHalfStr, secondHalfStr
}
