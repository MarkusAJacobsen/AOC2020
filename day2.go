package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func day2Part1(in []string) {
	start := time.Now()

	m := 0
	// Pattern low-hi let: pw
	for _, p := range in {
		re := regexp.MustCompile("[: -]+")
		split := re.Split(p, -1)

		var bound []string
		for i := range split {
			bound = append(bound, split[i])
		}

		low, _ := strconv.Atoi(bound[0])
		hi, _ := strconv.Atoi(bound[1])
		t := bound[2]
		pw := bound[3]

		req := strings.Split(pw, t)
		if (len(req) - 1) >= low && (len(req) -1) <= hi {
			m++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Day 2 part 1 %s\n", elapsed)
	fmt.Println(m)
}

func day2Part2(in []string) {
	start := time.Now()

	m := 0
	// Pattern low-hi let: pw
	for _, p := range in {
		re := regexp.MustCompile("[: -]+")
		split := re.Split(p, -1)

		var bound []string
		for i := range split {
			bound = append(bound, split[i])
		}

		low, _ := strconv.Atoi(bound[0])
		hi, _ := strconv.Atoi(bound[1])
		t := bound[2]
		pw := bound[3]

		if !(string(pw[low - 1]) == t && string(pw[hi - 1]) == t) {
			if string(pw[low - 1]) == t || string(pw[hi - 1]) == t {
				m++
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Day 2 part 2 %s\n", elapsed)
	fmt.Println(m)
}