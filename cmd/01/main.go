package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var doTwo = flag.Bool("2", false, "Run solution for the 2nd problem.")

type locationList []int

func distanceBetweenPoitns(a, b int) int {
	return max(a, b) - min(a, b)
}

func distance(a, b locationList) int {
	slices.Sort(a)
	slices.Sort(b)

	var distance int
	for i := range len(a) {
		distance += distanceBetweenPoitns(a[i], b[i])
	}

	return distance
}

func similarity(a, b locationList) int {
	var similarity int

	bCounts := make(map[int]int)
	for _, bi := range b {
		c, ok := bCounts[bi]
		if !ok {
			bCounts[bi] = 1

			continue
		}

		bCounts[bi] = c + 1
	}

	for _, ai := range a {
		similarity += bCounts[ai] * ai
	}

	return similarity
}

func main() {
	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	assert(err == nil, fmt.Sprintf("error: %s", err))

	var a, b locationList
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(string(line), "  ")
		assert(len(parts) == 2, fmt.Sprintf("invalid input: \"%s\"", line))

		if v, err := strconv.Atoi(strings.TrimSpace(parts[0])); err == nil {
			a = append(a, v)
		}

		if v, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			b = append(b, v)
		}
	}

	if *doTwo {
		fmt.Println(similarity(a, b))

		return
	}

	fmt.Println(distance(a, b))

	return
}

func assert(b bool, msg string) {
	if !b {
		panic(msg)
	}
}
