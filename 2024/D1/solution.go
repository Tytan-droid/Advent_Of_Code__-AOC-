package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func part1(inputDay string) int {
	var reponse = 0
	var l1 []int
	var l2 []int
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var v int
		var w int
		fmt.Sscanf(line, "%d %d", &v, &w)
		l1 = append(l1, v)
		l2 = append(l2, w)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	var x = len(l1)
	for i := range x {
		var o = l1[i] - l2[i]
		reponse = reponse + int(math.Abs(float64(o)))
	}
	return reponse
}

func part2(inputDay string) int {
	var reponse = 0
	var l1 []int
	var l2 []int
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var v int
		var w int
		fmt.Sscanf(line, "%d %d", &v, &w)
		l1 = append(l1, v)
		l2 = append(l2, w)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	var x = len(l1)
	for i := range x {
		var y = l1[i]
		for j := range len(l2) {
			var compteur = 0
			if l2[j] == y {
				compteur++
			}
			reponse = reponse + compteur*y
		}
	}
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2024 Day1 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "				//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "				//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
