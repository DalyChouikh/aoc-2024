package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/DalyChouikh/aoc-2024/file"
)

var mulPattern = regexp.MustCompile(`mul\((?<op1>[0-9]+)\,(?<op2>[0-9]+)\)`)
var doPattern = regexp.MustCompile(`do\(\)`)
var dontPattern = regexp.MustCompile(`don't\(\)`)
var combinedPattern = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|don\'t\(\)|do\(\)`)

func convertToStringSlice(intSlice [][]int) [][]string {
	stringSlice := make([][]string, len(intSlice))
	for i, row := range intSlice {
		stringSlice[i] = make([]string, len(row))
		for j, val := range row {
			stringSlice[i][j] = strconv.Itoa(val)
		}
	}
	return stringSlice
}

func extractSubmatches(input string, pattern *regexp.Regexp) [][]string {
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
	submatches := extractSubmatches(input, mulPattern)
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

func processStringWithDoAndDont(input string) (int, error) {
	result := 0
	mulEnabled := true

	matches := combinedPattern.FindAllString(input, -1)

	for _, match := range matches {
		switch {
		case mulPattern.MatchString(match):
			if mulEnabled {
				submatch := mulPattern.FindStringSubmatch(match)
				op1, _ := strconv.Atoi(submatch[1])
				op2, _ := strconv.Atoi(submatch[2])
				result += op1 * op2
			}
		case doPattern.MatchString(match):
			mulEnabled = true
		case dontPattern.MatchString(match):
			mulEnabled = false
		}
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

func processInputWithDoAndDont(inputs []string) (int, error) {
	totalResult := 0
	for _, input := range inputs {
		result, err := processStringWithDoAndDont(input)
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
	result1, err := processInput(input)
	if err != nil {
		log.Fatal(err)
	}
	result2, err := processInputWithDoAndDont(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result1)
	fmt.Println(result2)
}
