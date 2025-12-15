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

type point struct {
	x, y int
}

type form struct {
	area int
}

func ligne_to_int(str string) int {
	var size = len(str)
	var res = 0
	for i := 0; i < size; i++ {
		if str[i] == '#' {
			res = res*2 + 1
		} else {
			res = res * 2
		}
	}
	return res
}

func create_form(str string) form {
	var res form
	res.area = 0
	var lign = strings.Split(str, "\n")
	for _, c := range lign {
		if string(c) == "#" {
			res.area += 1
		}
	}
	return res
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n\n")
	var l_form []form
	var reponse = 0

	for _, line := range lines {
		if line[1] == ':' {
			l_form = append(l_form, create_form(line))
		} else {
			var lines_pb = strings.Split(line, "\n")
			for _, line_p := range lines_pb {
				var x, y int
				var ligne_p_split = strings.Split(line_p, ":")
				fmt.Sscanf(ligne_p_split[0], "%dx%d", &x, &y)

				var remain []int
				var remaining_str = strings.Split(strings.TrimPrefix(ligne_p_split[1], " "), " ")
				var place_p = 0
				var place_p2 = 0
				for i, r_str := range remaining_str {
					var r_, _ = strconv.Atoi(r_str)
					remain = append(remain, r_)
					place_p += r_ * 9
					place_p2 += r_ * l_form[i].area
				}
				if x*y >= place_p {
					reponse++
				}
				if x*y < place_p2 {
					continue
				}
			}
		}
	}
	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var reponse = 0
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day12 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "		//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "		//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
