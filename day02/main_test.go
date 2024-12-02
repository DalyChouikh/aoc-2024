package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsSafeReport(t *testing.T) {
	tcs := []struct {
		desc   string
		report []int
		isSafe bool
	}{
		{
			desc:   "Safe because the levels are all decreasing by 1 or 2",
			report: []int{7, 6, 4, 2, 1},
			isSafe: true,
		},
		{
			desc:   "Unsafe because 2 7 is an increase of 5",
			report: []int{1, 2, 7, 8, 9},
			isSafe: false,
		},
		{
			desc:   "Unsafe because 6 2 is a decrease of 4",
			report: []int{9, 7, 6, 2, 1},
			isSafe: false,
		},
		{
			desc:   "Unsafe because 1 3 is increasing but 3 2 is decreasing",
			report: []int{1, 3, 2, 4, 5},
			isSafe: false,
		},
		{
			desc:   "Unsafe because 4 4 is neither an increase or a decrease",
			report: []int{8, 6, 4, 4, 1},
			isSafe: false,
		},
		{
			desc:   "Safe because the levels are all increasing by 1, 2, or 3",
			report: []int{1, 3, 6, 7, 9},
			isSafe: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := isSafeReport(tc.report)
			if diff := cmp.Diff(tc.isSafe, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestIsSafeReportTolerance(t *testing.T) {
	tcs := []struct {
		desc   string
		report []int
		isSafe bool
	}{
		{
			desc:   "Safe without removing any level",
			report: []int{7, 6, 4, 2, 1},
			isSafe: true,
		},
		{
			desc:   "Unsafe regardless of which level is removed",
			report: []int{1, 2, 7, 8, 9},
			isSafe: false,
		},
		{
			desc:   "Unsafe regardless of which level is removed",
			report: []int{9, 7, 6, 2, 1},
			isSafe: false,
		},
		{
			desc:   "Safe by removing the second level, 3",
			report: []int{1, 3, 2, 4, 5},
			isSafe: true,
		},
		{
			desc:   "Safe by removing the third level, 4",
			report: []int{8, 6, 4, 4, 1},
			isSafe: true,
		},
		{
			desc:   "Safe without removing any level",
			report: []int{1, 3, 6, 7, 9},
			isSafe: true,
		},
		{
			desc:   "Safe by removing the third level, 2",
			report: []int{1, 3, 6, 7, 2},
			isSafe: true,
		},
		{
			desc:   "Safe by removing the third level, 9",
			report: []int{9, 3, 6, 7, 9},
			isSafe: true,
		},
		{
			desc:   "Safe by removing the third level, 75",
			report: []int{75, 76, 75, 72, 71, 68},
			isSafe: true,
		},
		{
			desc:   "Safe by removing the third level, 32",
			report: []int{34, 32, 37, 38, 40, 43, 45, 46},
			isSafe: true,
		},
		{
			desc:   "Unsafe regardless of which level is removed",
			report: []int{30, 33, 34, 36, 39, 36, 38, 40},
			isSafe: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := isSafeReportTolerance(tc.report)
			if diff := cmp.Diff(tc.isSafe, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumSafeReports(t *testing.T) {
	tcs := []struct {
		desc        string
		reports     []string
		badLevel    bool
		expectedVal int
	}{
		{
			desc: "without tolerance, in this example, 2 reports are safe",
			reports: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			badLevel:    false,
			expectedVal: 2,
		},
		{
			desc: "with tolerance, in this example, 4 reports are safe",
			reports: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			badLevel:    true,
			expectedVal: 4,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumSafeReports(tc.reports, tc.badLevel)
			if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
