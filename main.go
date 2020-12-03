package main

import (
	"bufio"
	"os"
)

func main() {
	input1 := getInput("input1.txt")
	input2 := getInput("input2.txt")
	input3 := getInput("input3.txt")

	day1Part1(input1)
	day1Part2(input1)
	day2Part1(input2)
	day2Part2(input2)
	day3Part1(input3)
	day3Part2(input3)
}

func getInput(name string) []string {
	f, _ := os.Open(name)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}