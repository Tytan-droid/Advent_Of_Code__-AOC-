package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
)

type point struct {
	x, y, z, idcircuit int
}

type distance_memo struct {
	p1, p2 *point
	d      int
}

type circuit1 struct {
	id     int
	taille int
}

//go:embed input.txt
var inputDay string

func distance(p1, p2 point) int {
	return int((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z))
}

func nearest_point(dm []distance_memo) [2]point {
	var min_d = math.MaxInt
	var res [2]point
	for _, d := range dm {
		if d.d < min_d && d.p1.idcircuit != d.p2.idcircuit {
			min_d = d.d
			res[0] = *d.p1
			res[1] = *d.p2
		}
	}
	return res
}

func part1(input string, nb_circuit int) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var xcoord, ycoord, zcoord int
	var l_points []point
	var idcircuit = 0
	var distances []distance_memo
	for id, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &xcoord, &ycoord, &zcoord)
		l_points = append(l_points, point{xcoord, ycoord, zcoord, id})
	}
	for i, p := range l_points {
		for j, q := range l_points {
			if i < j {
				var d = distance(p, q)
				distances = append(distances, distance_memo{&l_points[i], &l_points[j], d})
			}
		}
	}
	slices.SortFunc(distances, func(a, b distance_memo) int {
		return a.d - b.d
	})
	var reponse = 1
	for i := 0; i < nb_circuit; i++ {
		//var np = nearest_point(distances)
		var np = [2]point{*distances[i].p1, *distances[i].p2}
		idcircuit = np[0].idcircuit
		var old_idcircuit = np[1].idcircuit
		for j, p := range l_points {
			if p.idcircuit == old_idcircuit {
				l_points[j].idcircuit = idcircuit
			}
		}
	}
	var circuit = []circuit1{}
	for _, p := range l_points {
		var id = p.idcircuit
		var found = false
		for i, c := range circuit {
			if c.id == id {
				circuit[i].taille++
				found = true
			}
		}
		if !found {
			circuit = append(circuit, circuit1{id: id, taille: 1})
		}
	}
	slices.SortFunc(circuit, func(a, b circuit1) int {
		return b.taille - a.taille
	})
	for j := 0; j < 3; j++ {
		reponse *= circuit[j].taille
	}
	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var xcoord, ycoord, zcoord int
	var l_points []point
	var idcircuit = 0
	var circuit []circuit1
	var distances []distance_memo
	for id, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &xcoord, &ycoord, &zcoord)
		l_points = append(l_points, point{xcoord, ycoord, zcoord, id})
	}
	for i, p := range l_points {
		for j, q := range l_points {
			if i < j {
				var d = distance(p, q)
				distances = append(distances, distance_memo{&l_points[i], &l_points[j], d})
			}
		}
	}
	slices.SortFunc(distances, func(a, b distance_memo) int {
		return a.d - b.d
	})
	var reponse = 1
	var i = 0
	var boucle = true
	for boucle {
		var np = [2]point{*distances[i].p1, *distances[i].p2}
		idcircuit = np[0].idcircuit
		var old_idcircuit = np[1].idcircuit
		for j, p := range l_points {
			if p.idcircuit == old_idcircuit {
				l_points[j].idcircuit = idcircuit
			}
		}
		var nb_circuit = 0
		circuit = []circuit1{}
		for _, p := range l_points {
			var idc = p.idcircuit
			var found = false
			for i, c := range circuit {
				if c.id == idc {
					circuit[i].taille++
					found = true
				}
			}
			if !found {
				circuit = append(circuit, circuit1{id: idc, taille: 1})
				nb_circuit++
				if nb_circuit > 3 {
					break
				}
			}

		}
		if nb_circuit == 2 {
			boucle = false
		}
		i++
	}
	var near_p = nearest_point(distances)
	reponse = near_p[0].x * near_p[1].x
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day8 : ")
	fmt.Println("		-part1 : ", part1(inputDay, 1000), "		//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "		//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
