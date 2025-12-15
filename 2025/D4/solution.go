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

type grid map[pos]int

const (
	UP    = 0
	LEFT  = 1
	RIGHT = 2
	DOWN  = 3
	UL    = 4
	UR    = 5
	DL    = 6
	DR    = 7
)

func rouleau_dans_dir(p pos, dir int, g grid, x_max, y_max int) bool {
	if dir == UP {
		if p.y > 0 && g[pos{p.x, p.y - 1}] == 1 {
			return true
		}
	} else if dir == LEFT {
		if p.x > 0 && g[pos{p.x - 1, p.y}] == 1 {
			return true
		}
	} else if dir == RIGHT {
		if p.x < x_max && g[pos{p.x + 1, p.y}] == 1 {
			return true
		}
	} else if dir == DOWN {
		if p.y < y_max && g[pos{p.x, p.y + 1}] == 1 {
			return true
		}
	} else if dir == UL {
		if p.x > 0 && p.y > 0 && g[pos{p.x - 1, p.y - 1}] == 1 {
			return true
		}
	} else if dir == UR {
		if p.x < x_max && p.y > 0 && g[pos{p.x + 1, p.y - 1}] == 1 {
			return true
		}
	} else if dir == DL {
		if p.x > 0 && p.y < y_max && g[pos{p.x - 1, p.y + 1}] == 1 {
			return true
		}
	} else if dir == DR {
		if p.x < x_max && p.y < y_max && g[pos{p.x + 1, p.y + 1}] == 1 {
			return true
		}
	}
	return false
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var taille_x = len(lines)
	var taille_y = len(lines[0])
	var grid = make(grid)
	for j, line := range lines {
		for i, c := range line {
			if c == '@' {
				grid[pos{i, j}] = 1
			} else {
				grid[pos{i, j}] = 0
			}
		}
	}
	var reponse = 0
	for j, line := range lines {
		for i, c := range line {
			var p = pos{i, j}
			var nb_voisin = 0
			if c == '@' {
				var ind = 0
				for ind < 8 {
					if rouleau_dans_dir(p, ind, grid, taille_x, taille_y) {
						nb_voisin++
					}
					ind++
				}
				if nb_voisin < 4 {
					reponse++
				}
			}
		}
	}

	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var taille_x = len(lines)
	var taille_y = len(lines[0])
	var grid = make(grid)
	for j, line := range lines {
		for i, c := range line {
			if c == '@' {
				grid[pos{i, j}] = 1
			} else {
				grid[pos{i, j}] = 0
			}
		}
	}
	var reponse_t = 1
	var reponse = -1
	for reponse_t > 0 {
		reponse += reponse_t
		reponse_t = 0
		for j, line := range lines {
			for i, _ := range line {
				var p = pos{i, j}
				var nb_voisin = 0
				if grid[p] == 1 {
					var ind = 0
					for ind < 8 {
						if rouleau_dans_dir(p, ind, grid, taille_x, taille_y) {
							nb_voisin++
						}
						ind++
					}
					if nb_voisin < 4 {
						reponse_t++
						grid[pos{i, j}] = 0
					}
				}
			}
		}
	}
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day4 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
