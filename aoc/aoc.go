//go:build !test

package aoc

import (
	"flag"
)

var doTwo *bool

func Setup() {
	doTwo = flag.Bool("2", false, "Run solution for the 2nd problem.")

	flag.Parse()
}

func DoDayTwo() bool {
	if doTwo == nil {
		return false
	}

	return *doTwo
}
