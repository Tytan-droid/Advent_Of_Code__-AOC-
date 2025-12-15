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

type grid map[pos]uint8

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

type state struct {
	p   pos
	dir int
}

func (s state) N() state {
	return state{pos{s.p.x, s.p.y - 1}, UP}
}
func (s state) S() state {
	return state{pos{s.p.x, s.p.y + 1}, DOWN}
}
func (s state) E() state {
	return state{pos{s.p.x + 1, s.p.y}, RIGHT}
}
func (s state) W() state {
	return state{pos{s.p.x - 1, s.p.y}, LEFT}
}

func nextState(s state, g grid) []state {
	var c, ok = g[s.p]
	if !ok {
		fmt.Println("notok", s)
	}
	if s.dir == RIGHT {
		if c == '.' || c == '-' {
			return []state{s.E()}
		}
		if c == '/' {
			return []state{s.N()}
		}
		if c == '\\' {
			return []state{s.S()}
		}
		if c == '|' {
			return []state{s.N(), s.S()}
		}
	}
	if s.dir == LEFT {
		if c == '.' || c == '-' {
			return []state{s.W()}
		}
		if c == '/' {
			return []state{s.S()}
		}
		if c == '\\' {
			return []state{s.N()}
		}
		if c == '|' {
			return []state{s.N(), s.S()}
		}
	}
	if s.dir == DOWN {
		if c == '.' || c == '|' {
			return []state{s.S()}
		}
		if c == '/' {
			return []state{s.W()}
		}
		if c == '\\' {
			return []state{s.E()}
		}
		if c == '-' {
			return []state{s.E(), s.W()}
		}
	}

	if c == '.' || c == '|' {
		return []state{s.N()}
	}
	if c == '/' {
		return []state{s.E()}
	}
	if c == '\\' {
		return []state{s.W()}
	}
	if c == '-' {
		return []state{s.E(), s.W()}
	} else {
		fmt.Println("erreur dans next", c)
		return nil
	}

}

func part1(input string, x, y, dir int) int {
	var lines = strings.Split(input, "\n")
	var grid = make(grid)
	for j, line := range lines {
		for i, c := range line {
			grid[pos{i, j}] = uint8(c)
		}
	}
	var start = state{pos{x, y}, dir}

	var todo []state
	todo = append(todo, start)
	var visited = map[state]bool{}
	var energized = map[pos]bool{}
	for len(todo) > 0 {
		var s = todo[0]
		todo = todo[1:]
		visited[s] = true
		energized[s.p] = true
		var l = nextState(s, grid)
		for _, e := range l {
			if e.p.x < 0 || e.p.y < 0 {
				continue
			}
			if e.p.x >= len(lines[0]) || e.p.y >= len(lines) {
				continue
			}
			if visited[e] {
				continue
			}
			todo = append(todo, e)
		}
	}
	return len(energized)
}

func maxi(v1, v2 int) int {
	if v1 < v2 {
		return v2
	} else {
		return v1
	}
}

func part2(input string) int {
	var max = 0
	var val = 0
	var leninput = strings.Split(input, "\n")
	var y_max = len(leninput)
	var x_max = len(leninput[0])
	var y = 0
	var x = 0
	for y < y_max {
		val = part1(input, 0, y, RIGHT)
		max = maxi(max, val)
		val = part1(input, x_max-1, y, LEFT)
		max = maxi(max, val)
		y++
	}
	for x < x_max {
		val = part1(input, x, 0, DOWN)
		max = maxi(max, val)
		val = part1(input, x, y_max-1, UP)
		max = maxi(max, val)
		x++
	}

	return max
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2023 Day2 : ")
	fmt.Println("		-part1 : ", part1(inputDay, 0, 0, RIGHT), "				//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "				//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
