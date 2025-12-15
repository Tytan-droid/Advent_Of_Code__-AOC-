package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func part1(inputDay string) int {
	var reponse = 0
	var nb_game = 0
	var v int
	var w string
	var l []int
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	for _, line := range lines {

		nb_game++
		l = append(l, nb_game)
		var liness = strings.Split(line, ": ")
		var lines2 = strings.Split(liness[1], "; ")
		for _, line2 := range lines2 {

			var lines3 = strings.Split(line2, ", ")
			for _, line3 := range lines3 {
				fmt.Sscanf(line3, "%d %s", &v, &w)
				if (w == "blue" && v > 14) || (w == "red" && v > 12) || (w == "green" && v > 13) {
					l[nb_game-1] = 0
				}
			}
		}
	}

	for _, elt := range l {
		reponse = reponse + elt
	}

	return reponse
}

func part2(inputDay string) int {
	var reponse = 0
	var v int
	var w string
	var input = strings.TrimSuffix(inputDay, "\n")
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var game_i [3]int
		var liness = strings.Split(line, ": ")
		var lines2 = strings.Split(liness[1], "; ")
		for _, line2 := range lines2 {
			var lines3 = strings.Split(line2, ", ")
			for _, line3 := range lines3 {
				fmt.Sscanf(line3, "%d %s", &v, &w)
				if w == "blue" && v > game_i[2] {
					game_i[2] = v
				} else if w == "red" && v > game_i[0] {
					game_i[0] = v
				} else if w == "green" && v > game_i[1] {
					game_i[1] = v
				}
			}
		}
		reponse += game_i[0] * game_i[1] * game_i[2]
	}
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2023 Day2 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "				//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
