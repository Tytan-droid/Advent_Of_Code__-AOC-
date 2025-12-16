package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func part1(inputDay string) int {
	var reponse = 0
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "mul(")
	var x1, x2 int
	r, _ := regexp.Compile("^[0-9]+,[0-9]+$")
	for _, line := range lines {
		var lines2 = strings.Split(line, ")")
		for _, line2 := range lines2 {
			var ok_ = r.MatchString(line2)
			if ok_ {
				fmt.Sscanf(line2, "%d,%d", &x1, &x2)
				reponse += x1 * x2
			}
		}
	}
	return reponse
}

func part2(inputDay string) int {
	var reponse = 0
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines_dnt = strings.Split(input, "do()")
	for _, line_dnt := range lines_dnt {
		var lines_d = strings.Split(line_dnt, "don't()")
		for i, line_d := range lines_d {
			if i == 0 {
				reponse += part1(line_d)
			}
		}
	}

	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2024 Day3 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
