package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type pos struct {
	x, y int
}

type grid map[pos]int
type visited map[pos]bool

func mult_all(g grid, visit visited, j, len int) int {
	var retur = 1
	for i := range len {
		if (visit[pos{j, i}]) {
			retur *= g[pos{j, i}]
		}
	}
	return retur
}

func plus_all(g grid, visit visited, j, len int) int {
	var retur = 0
	for i := range len {
		if (visit[pos{j, i}]) {
			retur += g[pos{j, i}]
		}
	}
	return retur
}

func mult_all2(l1 []int, len int) int {
	var retur = 1
	var i = 0
	for i < len {
		retur *= l1[i]
		i++
	}
	return retur
}

func plus_all2(l1 []int, len int) int {
	var retur = 0
	var i = 0
	for i < len {
		retur += l1[i]
		i++
	}
	return retur
}

func num_colon(g grid, p pos, leny int, visit visited) int {
	var y = p.y
	var retur = 0
	for y < leny-1 {
		if visit[pos{p.x, y}] {
			retur = retur*10 + g[pos{p.x, y}]
		}
		y++
	}
	return retur
}

func part1(input string) int {
	re := regexp.MustCompile(`\s+`)
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var g = make(grid)
	var takethisnumber = make(visited)
	var operators []string
	for j, line := range lines {
		line = re.ReplaceAllString(line, " ")
		line = strings.TrimPrefix(line, " ")
		line = strings.TrimSuffix(line, " ")
		var c_line = strings.Split(line, " ")
		for i, c := range c_line {
			var st, _ = strconv.Atoi(string(c))
			if c == "*" {
				operators = append(operators, "*")
			} else if c == "+" {
				operators = append(operators, "+")
			} else {
				takethisnumber[pos{i, j}] = true
			}
			g[pos{i, j}] = st
		}
	}
	var reponse = 0
	var len = len(operators)

	for j, op := range operators {
		if op == "*" {
			reponse += mult_all(g, takethisnumber, j, len)
		} else if op == "+" {
			reponse += plus_all(g, takethisnumber, j, len)
		}
	}

	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var g = make(grid)
	var takethisnumber = make(visited)
	var operators []string
	var leny = len(lines)
	var lenx = len(lines[0])
	for j, line := range lines {
		for i, c := range line {
			var st, _ = strconv.Atoi(string(c))
			if string(c) == "*" {
				operators = append(operators, "*")
			} else if string(c) == "+" {
				operators = append(operators, "+")
			} else if string(c) != " " || j == leny-1 {
				takethisnumber[pos{i, j}] = true
			}
			g[pos{i, j}] = st
		}
	}
	var reponse = 0
	var po = pos{lenx - 1, 0}

	for po.x > 0 {
		var list []int
		for takethisnumber[pos{po.x, leny - 1}] {
			list = append(list, num_colon(g, po, leny, takethisnumber))
			po.x--
		}
		list = append(list, num_colon(g, po, leny, takethisnumber))
		po.x--
		po.x--
		var op = operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		if op == "*" {
			reponse += mult_all2(list, len(list))
		} else if op == "+" {
			reponse += plus_all2(list, len(list))
		}
	}

	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day6 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
