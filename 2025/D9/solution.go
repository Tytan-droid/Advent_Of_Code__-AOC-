package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x int
	y int
}

func intersection_deux_rect(p1, q1, p2, q2 Point) bool {
	if p1 == p2 && q1 == q2 {
		return false
	}
	var xbd Point = Point{max(p1.x, q1.x), max(p1.y, q1.y)}
	var xhg Point = Point{min(p1.x, q1.x), min(p1.y, q1.y)}

	var ybd Point = Point{max(p2.x, q2.x), max(p2.y, q2.y)}
	var yhg Point = Point{min(p2.x, q2.x), min(p2.y, q2.y)}

	if xhg.x >= ybd.x || xhg.y >= ybd.y {
		return false
	}
	if xbd.x <= yhg.x || xbd.y <= yhg.y {
		return false
	}
	if xhg.x >= ybd.x && xhg.y >= ybd.y && xbd.x <= ybd.x && xbd.y <= ybd.y {
		return false
	}
	if yhg.x >= xbd.x && yhg.y >= xbd.y && ybd.x <= xbd.x && ybd.y <= xbd.y {
		return false
	}

	return true
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var l_points = []Point{}

	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		point := Point{x: x, y: y}
		l_points = append(l_points, point)
	}
	var reponse = 0
	for _, p := range l_points {
		for _, q := range l_points {
			reponse = max(reponse, int(math.Abs(float64((p.x-q.x)))+1)*int(math.Abs(float64((p.y-q.y)))+1))
		}
	}

	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var l_points = []Point{}
	var points_ok map[Point]bool = make(map[Point]bool)
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		point := Point{x: x, y: y}
		l_points = append(l_points, point)
	}

	for _, p := range l_points {
		for _, q := range l_points {
			if p.x == q.x {
				for y := min(p.y, q.y); y <= max(p.y, q.y); y++ {
					points_ok[Point{x: p.x, y: y}] = true
				}
			}
			if p.y == q.y {
				for x := min(p.x, q.x); x <= max(p.x, q.x); x++ {
					points_ok[Point{x: x, y: p.y}] = true
				}
			}
		}
	}
	var reponse = 0
	var l_segments []struct{ p, q Point }
	for _, p := range l_points {
		for _, q := range l_points {
			if p.x == q.x || p.y == q.y {
				l_segments = append(l_segments, struct{ p, q Point }{p: p, q: q})
			}
		}
	}

	for _, p := range l_points {
		for _, q := range l_points {
			var ajout = true
			for _, segments := range l_segments {
				if intersection_deux_rect(p, q, segments.p, segments.q) {
					ajout = false
				}
			}
			if ajout {
				reponse = max(reponse, int(math.Abs(float64((p.x-q.x)))+1)*int(math.Abs(float64((p.y-q.y)))+1))
			}
		}
	}

	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day9 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "			//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
