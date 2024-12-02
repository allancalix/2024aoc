package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/allancalix/2024aoc/aoc"
)

type reportTrajectory int

const (
	reportTrajectoryUnknown reportTrajectory = iota
	reportTrajectoryIncreasing
	reportTrajectoryDecreasing
)

type reportLevels []int

func levelsInSafeRange(prev, curr int) bool {
	delta := max(curr, prev) - min(curr, prev)

	if delta >= 1 && delta <= 3 {
		return true
	}

	return false
}

func reportHasViolations(report reportLevels) int {
	last := report[0]

	var trajectory reportTrajectory
	for i, level := range report[1:] {
		if level == last {
			return i + 1
		}

		var currentTrajectory reportTrajectory
		if level > last {
			currentTrajectory = reportTrajectoryIncreasing
		} else {
			currentTrajectory = reportTrajectoryDecreasing
		}

		if trajectory != reportTrajectoryUnknown && trajectory != currentTrajectory {
			return i + 1
		}

		trajectory = currentTrajectory

		if !levelsInSafeRange(last, level) {
			return i + 1
		}

		last = level
	}

	return -1
}

func reactorIsSafe(report reportLevels) bool {
	return reportHasViolations(report) == -1
}

func reactorIsSafeWithTolerance(report reportLevels) bool {
	violation := reportHasViolations(report)
	if violation == -1 {
		return true
	}

	// Try removing each entry in the report and check if the reactor is safe. This operation
	// is O(n * k) where n is the number of reports and k is the number of entries in the report.
	for i := 0; i < len(report); i++ {
		r := make(reportLevels, len(report)-1)

		if i == 0 {
			copy(r, report[1:])
		}

		if i == len(report)-1 {
			copy(r, report[:i])
		}

		if i > 0 && i < len(report)-1 {
			copy(r, report[:i])
			copy(r[i:], report[i+1:])
		}

		if reactorIsSafe(r) {
			return true
		}
	}

	return false
}

func main() {
	aoc.Setup()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	var reports []reportLevels
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		levels := strings.Split(line, " ")
		var report reportLevels
		for _, l := range levels {
			level, err := strconv.Atoi(l)
			if err != nil {
				log.Fatalf("failed to convert level to int: %v", err)
			}

			report = append(report, level)
		}
		reports = append(reports, report)
	}

	var totalSafe int
	if aoc.DoDayTwo() {
		for _, report := range reports {
			if reactorIsSafeWithTolerance(report) {
				totalSafe++
			}
		}
	} else {
		for _, report := range reports {
			if reactorIsSafe(report) {
				totalSafe++
			}
		}
	}

	fmt.Println("Safe", totalSafe, "Out of", len(reports))
}
