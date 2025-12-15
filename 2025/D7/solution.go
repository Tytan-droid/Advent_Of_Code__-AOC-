package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type pos struct {
	x, y int
}
type pos_memo struct {
	x, y, nb_w int
}

type grid map[pos]string
type visited map[pos]bool
type memo map[pos]int

func (p pos) move_down() pos {
	return pos{x: p.x, y: p.y + 1}
}

func (p pos) Split(len_x int) []pos {
	if p.x == 0 {
		return []pos{
			{x: p.x + 1, y: p.y},
		}
	} else if p.x == len_x-1 {
		return []pos{
			{x: p.x - 1, y: p.y},
		}
	} else {
		return []pos{
			{x: p.x - 1, y: p.y},
			{x: p.x + 1, y: p.y},
		}
	}
}

func (p pos_memo) nb_way(len_x, len_y int, memo memo, visited visited, g grid) int {
	if p.y == len_y-1 {
		visited[pos{p.x, p.y}] = true
		memo[pos{p.x, p.y}] = 1
		return p.nb_w
	} else if (visited[pos{p.x, p.y}]) {
		return memo[pos{p.x, p.y}]
	} else if (g[pos{p.x, p.y}] == ".") {
		visited[pos{p.x, p.y}] = true
		memo[pos{p.x, p.y}] = pos_memo{p.x, p.y + 1, p.nb_w}.nb_way(len_x, len_y, memo, visited, g)
		return memo[pos{p.x, p.y}]
	} else if (g[pos{p.x, p.y}] == "^") {
		visited[pos{p.x, p.y}] = true
		var total = 0
		var lp = pos{p.x, p.y + 1}.Split(len_x)
		for _, next := range lp {
			total += pos_memo{next.x, next.y, p.nb_w}.nb_way(len_x, len_y, memo, visited, g)
		}
		memo[pos{p.x, p.y}] = total
		return total
	}
	return pos_memo{p.x, p.y + 1, p.nb_w}.nb_way(len_x, len_y, memo, visited, g)
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var g = make(grid)
	var visited = make(visited)
	var start = pos{0, 0}
	for j, line := range lines {
		for i, c := range line {
			if c == 'S' {
				start = pos{i, j}
			}
			g[pos{i, j}] = string(c)
		}
	}
	len_x := len(lines[0])
	len_y := len(lines)
	var reponse = 0
	var queue = []pos{start}
	visited[start] = true
	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]
		if current.y <= len_y-1 {
			if g[pos{current.x, current.y + 1}] == "." && !visited[pos{current.x, current.y + 1}] {
				var next = current.move_down()
				queue = append(queue, next)
				visited[next] = true
			} else if (g[pos{current.x, current.y + 1}] == "^") && !visited[pos{current.x, current.y + 1}] {
				reponse++
				visited[pos{current.x, current.y + 1}] = true
				var next = current.move_down()
				for _, p := range next.Split(len_x) {
					if !visited[p] {
						queue = append(queue, p)
						visited[p] = true
					}
				}
			}
		}
	}
	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var g = make(grid)
	var visited = make(visited)
	var start = pos_memo{0, 0, 0}
	for j, line := range lines {
		for i, c := range line {
			if c == 'S' {
				start = pos_memo{i, j, 0}
			}
			g[pos{i, j}] = string(c)
		}
	}
	len_x := len(lines[0])
	len_y := len(lines)
	var mem = make(memo)
	start.nb_w = pos_memo{start.x, start.y, 1}.nb_way(len_x, len_y, mem, visited, g)

	return start.nb_w
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day7 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "	//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
