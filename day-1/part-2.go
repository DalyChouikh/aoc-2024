package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("Invalid line: %s", line)
		}
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("Error converting numbers: %s", line)
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Error reading file: %w", err)
	}
	return list1, list2, nil
}

func countOccurences(list []int, element int) (count int) {
	for _, v := range list {
		if v == element {
			count = count + 1
		}
	}
	return count
}

func calculateSimilarity(list1, list2 []int) (similarityScore int) {
	for i := 0; i < len(list1); i++ {
		similarityScore = similarityScore + (list1[i] * countOccurences(list2, list1[i]))
	}

	return similarityScore

}

func main() {
	l1, l2, err := readFile("day-1.input")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	similarityScore := calculateSimilarity(l1, l2)
	fmt.Println(similarityScore)
}
