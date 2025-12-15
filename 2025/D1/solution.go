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
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var nb = 50
	var compteur = 0
	for _, line := range lines {
		var v = line[0:1]
		var w = (line[1:])
		c, _ := strconv.Atoi(w)
		if v == "R" {
			nb = (nb + c) % 100
		} else {
			nb = (nb - c + 100) % 100
		}
		if nb == 0 {
			compteur++
		}
	}
	return compteur
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var nb = 50
	var compteur = 0
	for _, line := range lines {
		var v = line[0:1]
		var w = (line[1:])
		c, _ := strconv.Atoi(w)
		if v == "R" {
			nb = (nb + c)
			for nb > 99 {
				nb = nb - 100
				compteur++
			}
		} else {
			nb = nb - c
			for nb < 0 {
				if nb+c != 0 {
					compteur++
				}
				nb = nb + 100
			}
			if nb == 0 {
				compteur++
			}
		}
	}
	return compteur
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day1 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
