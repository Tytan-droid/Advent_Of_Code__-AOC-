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

func part1(inputDay string) int {
	var reponse = 0
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var sens = 0 //1 : monte et -1 : descend
		var l_elt = strings.Split(line, " ")
		var safe = true
		safe = l_safe(l_elt, safe, sens)
		if safe {
			reponse++
		}
	}
	return reponse
}

func l_safe(l_elt []string, safe bool, sens int) bool {
	var current, _ = strconv.Atoi(l_elt[0])
	for i := range len(l_elt) - 1 {
		var next, _ = strconv.Atoi(l_elt[i+1])
		if current > next {
			if sens == 1 || current-next > 3 {
				safe = false
				break
			}
			sens = -1
		} else if current < next {
			if sens == -1 || next-current > 3 {
				safe = false
				break
			}
			sens = 1
		} else {
			safe = false
			break
		}
		current = next
	}
	return safe
}

func part2(inputDay string) int {
	var reponse = 0
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var l_elt = strings.Split(line, " ")
		for i := range l_elt {
			var safe = true
			var sens = 0 //1 : monte et -1 : descend
			var l_elt_bis []string
			for j := range l_elt {
				if j != i {
					l_elt_bis = append(l_elt_bis, l_elt[j])
				}
			}
			safe = l_safe(l_elt_bis, safe, sens)
			if safe {
				reponse++
				break
			}
		}

	}
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2024 Day2 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "				//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "				//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
