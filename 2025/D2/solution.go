package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func is_invalid(num, div int) bool {
	var s = string(strconv.Itoa(num))
	if (len(s) % div) != 0 {
		return false
	} else {
		for i := 0; i < int(len(s)/div); i++ {
			for j := 1; j < div; j++ {
				if s[i] != (s[j*len(s)/div+i]) {
					return false
				}
			}

		}
		return true
	}
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, ",")
	var reponse = 0
	for _, line := range lines {
		var rangel = strings.Split(line, "-")
		var start, _ = strconv.Atoi(rangel[0])
		var end, _ = strconv.Atoi(rangel[1])
		var i int
		for i = start; i <= end; i++ {
			if is_invalid(i, 2) {
				reponse += i
			}
		}
	}
	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, ",")
	var reponse = 0
	for _, line := range lines {
		var rangel = strings.Split(line, "-")
		var start, _ = strconv.Atoi(rangel[0])
		var end, _ = strconv.Atoi(rangel[1])
		var i int
		var div int

		for i = start; i <= end; i++ {
			var s = string(strconv.Itoa(i))
			for div = 2; div <= int(len(s)); div++ {
				if is_invalid(i, div) {
					reponse += i
					break
				}
			}
		}
	}
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day2 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
