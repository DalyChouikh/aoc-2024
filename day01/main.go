package main

import (
	"fmt"
	"math"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/DalyChouikh/aoc-2024/file"
)

func getNumLists(s []string) ([]int, []int) {
	a := make([]int, len(s), len(s))
	b := make([]int, len(s), len(s))
	for idx, str := range s {
		ss := strings.Split(str, "   ")
		a[idx], _ = strconv.Atoi(ss[0])
		b[idx], _ = strconv.Atoi(ss[1])
	}
	return a, b
}

func getTotalDistance(s []string) int {
	// Get two list integer values from input
	a, b := getNumLists(s)
	// Sort them
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	// Get distance
	result := 0
	for c := 0; c < len(a); c++ {
		d := math.Abs(float64(a[c]) - float64(b[c]))
		result += int(d)
	}
	return result
}

func getTotalSimilarityScore(s []string) int {
	// Get two list integer values from input
	a, b := getNumLists(s)
	// Create map to store num repeated times of the right column values
	m := map[int]int{}
	for _, v := range b {
		if i, ok := m[v]; ok {
			m[v] = i + 1
		} else {
			m[v] = 1
		}
	}
	// Calculate similarity score
	result := 0
	for _, v := range a {
		d := 0
		if i, ok := m[v]; ok {
			d = v * i
		}
		result += d
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("day01/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getTotalDistance(output))
	fmt.Println(getTotalSimilarityScore(output))
}
