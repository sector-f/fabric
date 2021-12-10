package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	width := flag.Float64("w", 0, "Width of fabric in inches")
	length := flag.Float64("l", 0, "Length of fabric in yards")
	weight := flag.Float64("t", 0, "Weight of fabric in oz. per square yard")
	flag.Parse()

	if *width == 0 || *length == 0 || *weight == 0 {
		fmt.Fprintln(os.Stderr, "Some flags are missing")
		os.Exit(1)
	}

	totalWeight := calcWeight(*width, *length, *weight)

	fmt.Printf(
		"%.2f oz. (%.2f lbs.)\n",
		totalWeight,
		totalWeight/16,
	)
}

func calcWeight(width, length, weight float64) float64 {
	totalSquareYards := (width * length) / 36
	return weight * totalSquareYards
}
