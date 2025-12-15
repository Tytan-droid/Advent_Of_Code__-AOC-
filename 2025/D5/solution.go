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

func part1(input string) int {
	var list_tiret = true
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var reponse = 0
	var fresh_d []int
	var fresh_f []int

	for _, line := range lines {
		if len(line) == 0 {
			list_tiret = false
		}
		if list_tiret {
			var rangel = strings.Split(line, "-")
			var start, _ = strconv.Atoi(rangel[0])
			var end, _ = strconv.Atoi(rangel[1])
			fresh_d = append(fresh_d, start)
			fresh_f = append(fresh_f, end)
		} else {
			var i, _ = strconv.Atoi(line)
			var ajout = false
			for ind, _ := range fresh_d {
				if fresh_d[ind] <= i && fresh_f[ind] >= i && ajout == false {
					reponse++
					ajout = true
				}
			}
		}
	}
	return reponse
}

func part2(input string) int {
	var list_tiret = true
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var reponse = 0
	var fresh_d []int
	var fresh_f []int
	var max = 0
	for _, line := range lines {
		if len(line) == 0 {
			list_tiret = false
		}
		if list_tiret {
			var rangel = strings.Split(line, "-")
			var start, _ = strconv.Atoi(rangel[0])
			var end, _ = strconv.Atoi(rangel[1])

			for ind := range fresh_d {
				if start <= fresh_f[ind] && start >= fresh_d[ind] {
					start = fresh_f[ind] + 1
				} else if end >= fresh_d[ind] && end <= fresh_f[ind] {
					end = fresh_d[ind] - 1
				}
				if fresh_d[ind] <= end && fresh_d[ind] >= start {
					fresh_d[ind] = end + 1
				} else if fresh_f[ind] >= start && fresh_f[ind] <= end {
					fresh_f[ind] = start - 1
				}
			}
			if start <= end {
				fresh_d = append(fresh_d, start)
				fresh_f = append(fresh_f, end)
				if end > max {
					max = end
				}
			}
		}
	}
	for ind := range fresh_d {
		if fresh_d[ind] <= fresh_f[ind] {
			reponse += fresh_f[ind] - fresh_d[ind] + 1
		}
	}

	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day5 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "	//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
