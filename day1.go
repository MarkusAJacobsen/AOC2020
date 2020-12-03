package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func day1Part1(in []string) {
	var inN []int64
	for _, i := range in {
		sci, _ := strconv.ParseInt(i, 10, 64)
		inN = append(inN, sci)
	}

	// x + y = 2020
	// x = 2020 - y

	// Matrix
	// x 0, 1... 2020
	// y 2020, 2019... 0

	start := time.Now()
	var int32b []int
	for _, val := range inN {
		int32b = append(int32b, int(val))
	}
	sort.Ints(int32b)

	sorted := int32b
	sort.Sort(sort.Reverse(sort.IntSlice(int32b)))
	for _, x := range sorted {
		for _, y := range int32b {
			if x != y && 2020 - y == x {
				elapsed := time.Since(start)
				fmt.Printf("Day 1 part 1 %s\n", elapsed)
				fmt.Println(x ,y, x * y)
				return
			}
		}
	}
}

func day1Part2(in []string) {
	var inN []int64
	for _, i := range in {
		sci, _ := strconv.ParseInt(i, 10, 64)
		inN = append(inN, sci)
	}

	// x + y + z = 2020
	// x = 2020 - y - z

	start := time.Now()
	var int32b []int
	for _, val := range inN {
		int32b = append(int32b, int(val))
	}
	sort.Ints(int32b)

	sorted := int32b
	sort.Sort(sort.Reverse(sort.IntSlice(int32b)))
	for _, x := range sorted {
		for _, y := range int32b {
			for _, z := range sorted {
				if x != y && 2020 - y - z == x {
					elapsed := time.Since(start)
					fmt.Printf("Day 1 part 2 %s\n", elapsed)
					fmt.Println(x ,y, z, x * y * z)
					return
				}
			}

		}
	}
}
