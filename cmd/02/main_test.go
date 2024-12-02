package main

import (
	"testing"
)

func TestReactorIsSafe(t *testing.T) {
	tests := []struct {
		report reportLevels
		safe   bool
	}{
		{
			safe:   true,
			report: reportLevels{7, 6, 4, 2, 1},
		},
		{
			safe:   false,
			report: reportLevels{1, 2, 7, 8, 9},
		},
		{
			safe:   false,
			report: reportLevels{9, 7, 6, 2, 1},
		},
		{
			safe:   false,
			report: reportLevels{1, 3, 2, 4, 5},
		},
		{
			safe:   false,
			report: reportLevels{8, 6, 4, 4, 1},
		},
		{
			safe:   true,
			report: reportLevels{1, 3, 6, 7, 9},
		},
	}

	for _, test := range tests {
		if got := reactorIsSafe(test.report); got != test.safe {
			t.Errorf("reactorIsSafe(%v) = %v, want %v", test.report, got, test.safe)
		}
	}
}

func TestReactorIsSafeWithTolerance(t *testing.T) {
	tests := []struct {
		report reportLevels
		safe   bool
	}{
		{
			safe:   true,
			report: reportLevels{7, 6, 4, 2, 1},
		},
		{
			safe:   false,
			report: reportLevels{1, 2, 7, 8, 9},
		},
		{
			safe:   false,
			report: reportLevels{9, 7, 6, 2, 1},
		},
		{
			safe:   true,
			report: reportLevels{1, 3, 2, 4, 5},
		},
		{
			safe:   true,
			report: reportLevels{8, 6, 4, 4, 1},
		},
		{
			safe:   true,
			report: reportLevels{8, 8, 5, 3, 1},
		},
		{
			safe:   true,
			report: reportLevels{1, 3, 6, 7, 9},
		},
		{
			safe:   true,
			report: reportLevels{1, 3, 6, 7, 9, 100},
		},
		{
			safe:   true,
			report: reportLevels{100, 2, 3, 6, 7, 9},
		},
		{
			safe:   false,
			report: reportLevels{1, 1, 1},
		},
		{
			safe:   true,
			report: reportLevels{1, 1, 2},
		},
		{
			safe:   true,
			report: reportLevels{1, 2, 2},
		},
		{
			safe:   true,
			report: reportLevels{1, 2, 2, 3},
		},
		{
			safe:   false,
			report: reportLevels{1, 2, 2, 6},
		},
		{
			safe:   true,
			report: reportLevels{1, 2, 1, 4, 5},
		},
		{
			safe:   false,
			report: reportLevels{1, 3, 5, 2, 4, 3},
		},
		{
			safe:   false,
			report: reportLevels{1, 3, 3, 3},
		},
		{
			safe:   false,
			report: reportLevels{69, 69, 66, 63, 59, 58},
		},
		{
			safe:   true,
			report: reportLevels{84, 82, 83, 84, 85, 88, 90},
		},
	}

	for _, test := range tests {
		if got := reactorIsSafeWithTolerance(test.report); got != test.safe {
			t.Errorf("reactorIsSafeWithTolerance(%v) = %v, want %v", test.report, got, test.safe)
		}
	}
}

func TestLevelsInSafeRange(t *testing.T) {
	tests := []struct {
		prev, curr int
		safe       bool
	}{
		{1, 2, true},
		{2, 1, true},
		{1, 3, true},
		{3, 1, true},
		{0, 4, false},
		{4, 0, false},
		{4, 4, false},
	}

	for _, test := range tests {
		if got := levelsInSafeRange(test.prev, test.curr); got != test.safe {
			t.Errorf("levelsInSafeRange(%v, %v) = %v, want %v", test.prev, test.curr, got, test.safe)
		}
	}
}
