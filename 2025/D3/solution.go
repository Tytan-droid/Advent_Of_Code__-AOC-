package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func max(line string, len_line int) int {
	var i = 0
	var max = line[0] - 48
	for i < len_line {
		if line[i]-48 > max {
			max = line[i] - 48
		}
		i++
	}
	return int(max)
}

func argmax(line string, len_line int) int {
	var i = 0
	var max = line[0] - 48
	var index = 0
	for i < len_line {
		if line[i]-48 > max {
			max = line[i] - 48
			index = i
		}
		i++
	}
	return index
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var reponse = 0
	for _, line := range lines {
		var len_line = len(line)
		var max1 = max(line, len_line-1)
		var j = argmax(line, len_line-1)
		var max2 = max(line[j+1:], len_line-j-1)
		reponse += 10*max1 + max2

	}
	return reponse
}

func part2(input string, nb int) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var reponse = 0
	for _, line := range lines {
		var k = 0
		var len_line = len(line) - nb + 1
		var reponse_line = 0
		for k < nb {
			var max = max(line, len_line)
			var j = argmax(line, len_line)
			reponse_line = reponse_line*10 + max
			line = line[j+1:]
			len_line = len_line - j
			k++
		}
		reponse += reponse_line
	}
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day3 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay, 12), "		//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
