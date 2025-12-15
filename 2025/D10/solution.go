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

func press_or_not(switch_bi int, list_button []int, req_button []int, indice int) int {
	if all_on2(req_button) || switch_bi == 0 {
		return 0
	}
	if indice >= len(list_button) {
		return 5000
	}

	if too_push(req_button) && false {
		return 5000
	}
	var clone_switch int
	var new_req_button []int
	clone_switch, new_req_button = press_button(indice, list_button, switch_bi, req_button)

	return min(1+press_or_not(clone_switch, list_button, new_req_button, indice+1), press_or_not(switch_bi, list_button, req_button, (indice+1)))

}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var reponse = 0
	for _, line := range lines {
		var switch_bi int
		var portion = strings.Split(line, " ")
		var liste_button []int
		var req_button []int
		for _, p := range portion {
			if p[0] == '[' {
				switch_bi = cover_switch(p)

			} else if p[0] == '(' {
				liste_button = append(liste_button, cover_button(p, len(portion[0])-2))
			} else if p[0] == '{' {
				req_button = cover_req(p)
			}
		}
		var res = press_or_not(switch_bi, liste_button, req_button, 0)
		reponse += res

	}
	return reponse
}

func cover_switch(char_switch string) int {
	var res int
	var size = len(char_switch) - 2
	for i := 0; i < size; i++ {
		if char_switch[i+1] == '#' {
			res = res*2 + 1
		} else {
			res = res * 2
		}
	}
	return res
}

func cover_button(button string, size_switch int) int {
	var res = 0
	var split = strings.Split(button[1:len(button)-1], ",")
	for i := 0; i < len(split); i++ {
		var indice, _ = strconv.Atoi(split[i])
		var modif = 1 << (size_switch - indice - 1)
		res = res | modif
	}
	return res
}

func cover_req(char_raq string) []int {
	var res []int
	var split = strings.Split(char_raq[1:len(char_raq)-1], ",")
	for i := 0; i < len(split); i++ {
		var nb, _ = strconv.Atoi(split[i])
		res = append(res, nb)
	}
	return res

}

func press_button(indice int, liste_button []int, switch_bi int, req_button []int) (int, []int) {
	num_butt := liste_button[indice]

	res_int := switch_bi ^ num_butt

	res_lint := make([]int, len(req_button))
	copy(res_lint, req_button)

	for compt := 0; compt < len(req_button); compt++ {
		if (num_butt & (1 << compt)) != 0 {
			res_lint[len(res_lint)-compt-1]--
		}
	}

	return res_int, res_lint
}

func all_on2(req_button []int) bool {
	for _, comp := range req_button {
		if comp != 0 {
			return false
		}
	}

	return true
}

func too_push(req_button []int) bool {
	for _, comp := range req_button {
		if comp < 0 {
			return true
		}
	}
	return false
}

func Hashage(indice int, req_button []int) string {
	key := strconv.Itoa(indice) + "|"
	for _, v := range req_button {
		key += strconv.Itoa(v) + ","
	}
	return key
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		var buttons []int
		var req []int

		for _, p := range parts {
			if p[0] == '(' {
				buttons = append(buttons, cover_button(p, len(parts[0])-2))
			} else if p[0] == '{' {
				req = cover_req(p)
			}
		}

		n := len(req)

		part1Sols := make(map[int][][]int)
		for mask := 0; mask < (1 << len(buttons)); mask++ {
			par := 0
			var used []int
			for i := 0; i < len(buttons); i++ {
				if (mask>>i)&1 == 1 {
					par ^= buttons[i]
					used = append(used, i)
				}
			}
			part1Sols[par] = append(part1Sols[par], used)
		}

		memo := make(map[string]int)

		var f func([]int) int
		f = func(cur []int) int {
			key := fmt.Sprint(cur)
			if v, ok := memo[key]; ok {
				return v
			}

			done := true
			for _, v := range cur {
				if v != 0 {
					done = false
					break
				}
			}
			if done {
				return 0
			}

			par := 0
			for i, v := range cur {
				if v%2 != 0 {
					par |= 1 << (n - 1 - i)
				}
			}

			best := 5000

			for _, sol := range part1Sols[par] {
				next := make([]int, n)
				copy(next, cur)

				ok := true
				for _, b := range sol {
					mask := buttons[b]
					for i := 0; i < n; i++ {
						if (mask & (1 << i)) != 0 {
							next[n-1-i]--
							if next[n-1-i] < 0 {
								ok = false
								break
							}
						}
					}
					if !ok {
						break
					}
				}
				if !ok {
					continue
				}

				for i := range next {
					if next[i]%2 != 0 {
						ok = false
						break
					}
					next[i] /= 2
				}
				if !ok {
					continue
				}

				cost := len(sol) + 2*f(next)
				if cost < best {
					best = cost
				}
			}

			memo[key] = best
			return best
		}

		res := f(req)
		total += res
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day10 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "		//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
