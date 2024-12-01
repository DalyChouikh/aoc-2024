package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func partition(list []int, low, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func quickSort(list []int, low, high int) []int {
	if low < high {
		pi := partition(list, low, high)
		quickSort(list, low, pi-1)
		quickSort(list, pi+1, high)
	}
	return list
}

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

func calculateDistance(list1, list2 []int) (totalDistance int) {
	for i := 0; i < len(list1); i++ {
		totalDistance = totalDistance + int(math.Abs(float64(list1[i]-list2[i])))
	}

	return totalDistance

}

func main() {
	l1, l2, err := readFile("1.input")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	totalDistance := calculateDistance(quickSort(l1, 0, len(l1)-1), quickSort(l2, 0, len(l2)-1))
	fmt.Println(totalDistance)
}
