package main

import (
	"fmt"
	"time"
)

func day3Part1(in []string) {
	start := time.Now()

	xGrad := 3
	count := 0

	x := xGrad
	for i := 1; i < len(in); i++ {
		li := x % len(in[i])

		if in[i][li] == '#' {
			count++
		}

		x += xGrad
	}
	elapsed := time.Since(start)
	fmt.Printf("Day 3 part 1 %s\n", elapsed)
	fmt.Println(count)
}

func day3Part2(in []string) {
	start := time.Now()

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var counts []int

	for _, s := range slopes {
		xGrad := s[0]
		yGrad := s[1]
		count := 0

		x := xGrad
		for i := yGrad; i < len(in); i += yGrad {
			li := x % len(in[i])

			if in[i][li] == '#' {
				count++
			}

			x += xGrad
		}

		counts = append(counts, count)
	}

	acc := 0
	for _, c := range counts {
		if acc == 0 {
			acc += c
		} else {
			acc *= c
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Day 3 part 2 %s\n", elapsed)
	fmt.Println(counts, acc)
}
