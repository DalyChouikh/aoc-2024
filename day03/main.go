package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/DalyChouikh/aoc-2024/file"
)

var pattern = regexp.MustCompile(`mul\((?<op1>[0-9]+)\,(?<op2>[0-9]+)\)`)

func extractSubmatches(input string) [][]string {
	return pattern.FindAllStringSubmatch(input, -1)
}

func parseOperands(match []string) (int, int, error) {
	operand1, err1 := strconv.Atoi(match[1])
	operand2, err2 := strconv.Atoi(match[2])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error parsing operands: %v, %v", err1, err2)
	}
	return operand1, operand2, nil
}

func processString(input string) (int, error) {
	submatches := extractSubmatches(input)
	result := 0
	for _, match := range submatches {
		op1, op2, err := parseOperands(match)
		if err != nil {
			return 0, err
		}
		result += op1 * op2
	}
	return result, nil
}

func processInput(inputs []string) (int, error) {
	totalResult := 0
	for _, input := range inputs {
		result, err := processString(input)
		if err != nil {
			return 0, err
		}
		totalResult += result
	}
	return totalResult, nil
}

func main() {
	input, err := file.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result, err := processInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
