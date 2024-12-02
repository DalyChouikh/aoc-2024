package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/DalyChouikh/aoc-2024/file"
)

func getReports(s []string) [][]int {
	result := make([][]int, len(s))
	for i, report := range s {
		r := strings.Split(report, " ")
		result[i] = make([]int, len(r))
		for j, level := range r {
			result[i][j], _ = strconv.Atoi(level)
		}
	}
	return result
}

func isSafeReport(report []int) bool {
	isInc := false
	if report[1]-report[0] > 0 {
		isInc = true
	}
	for idx := 1; idx < len(report); idx++ {
		if isInc && (report[idx]-report[idx-1]) > 0 && (report[idx]-report[idx-1]) < 4 {
			continue
		}
		if !isInc && (report[idx]-report[idx-1]) < 0 && (report[idx]-report[idx-1]) > -4 {
			continue
		}
		return false
	}
	return true
}

func isSafeReportTolerance(report []int) bool {
	inc := 0
	for idx := 1; idx < len(report); idx++ {
		if report[idx]-report[idx-1] > 0 {
			inc++
		} else {
			inc--
		}
	}
	if inc < len(report)-3 && inc > -(len(report)-3) {
		return false // there are more than one bad number
	}
	isInc := false
	if inc > 0 {
		isInc = true
	}
	for idx := 1; idx < len(report); idx++ {
		if isInc && (report[idx]-report[idx-1]) > 0 && (report[idx]-report[idx-1]) < 4 {
			continue
		}
		if !isInc && (report[idx]-report[idx-1]) < 0 && (report[idx]-report[idx-1]) > -4 {
			continue
		}
		// If there is a wrong level check current one and adjacent
		newReport := slices.Clone(report)
		newReport = append(newReport[:idx], newReport[idx+1:]...)
		if isSafeReport(newReport) {
			return true
		} else {
			newReport = slices.Clone(report)
			newReport = append(newReport[:idx-1], newReport[idx:]...)
			if isSafeReport(newReport) {
				return true
			} else {
				newReport = slices.Clone(report)
				if idx < len(report)-2 {
					newReport = append(newReport[:idx+1], newReport[idx+2:]...)
				} else {
					newReport = newReport[:len(newReport)-1]
				}
				if isSafeReport(newReport) {
					return true
				}
			}
		}
		return false
	}
	return true
}

func getNumSafeReports(s []string, tolerateBadLevel bool) int {
	result := 0
	reports := getReports(s)
	for _, r := range reports {
		if !tolerateBadLevel && isSafeReport(r) {
			result += 1
		} else if tolerateBadLevel && isSafeReportTolerance(r) {
			result += 1
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("day02/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumSafeReports(output, false))
	fmt.Println(getNumSafeReports(output, true))
}
