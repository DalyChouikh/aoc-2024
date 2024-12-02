package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(line []int) bool {
	direction := "increasing"
	if line[0] > line[1] {
		direction = "decreasing"
	}
	for i := 1; i < len(line); i++ {
		if direction == "increasing" && line[i] < line[i-1] {
			return false
		}
		if direction == "decreasing" && line[i] > line[i-1] {
			return false
		}
	}
	for i := 0; i < len(line)-1; i++ {
		diff := int(math.Abs(float64(line[i] - line[i+1])))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func readFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reportsCount int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var report []int
		for _, v := range line {
			value, _ := strconv.Atoi(v)
			report = append(report, value)
		}
		if isReportSafe(report) {
			reportsCount++
		}
	}
	return reportsCount
}

func main() {
	reportsCount := readFile("day-2.input")
	fmt.Println(reportsCount)
}
