package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fields[0] = strings.Trim(fields[0], ":")

		var fieldsInt []int
		for _, num := range fields {
			intNum, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			fieldsInt = append(fieldsInt, intNum)
		}
		result := 0
		fmt.Println(checkForEvaluation(fieldsInt))
		sums = append(sums, result)
	}

	result := sumOfVals(sums)

	fmt.Println(result)
}

func checkForEvaluation(numbers []int) interface{} {
	// operators := []string{"+", "*"}
	testValue := numbers[0]
	numbers = removeFromSlice(numbers, 0)
	operatorSpots := len(numbers) - 1

	resultSum := 0
	resultProduct := 0
	for _, val := range numbers {
		resultSum += val
		resultProduct *= val
	}
	if resultSum == testValue || resultProduct == testValue {
		return testValue
	}

	expr := "a"
	expression, _ := govaluate.NewEvaluableExpression(expr)
	for i := 0; i < operatorSpots; i++ {
		// if one of the operations == testValue, return testValue
		expression, _ = govaluate.NewEvaluableExpression(expr)
		expr += expression.String()
	}
	parameters := make(map[string]interface{}, 8)
	parameters["a"] = numbers[0]
	result, _ := expression.Evaluate(parameters)

	return result
}

func sumOfVals(values []int) int {
	sum := 0
	for _, val := range values {
		sum = sum + val
	}

	return sum
}

func removeFromSlice(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
