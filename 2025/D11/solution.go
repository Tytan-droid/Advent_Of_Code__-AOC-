package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Node struct {
	value    string
	children []string
}

func NewNode(value string) Node {
	return Node{
		value:    value,
		children: []string{},
	}
}

func (n *Node) AddChild(child string) {
	n.children = append(n.children, child)
}

func (n *Node) print_tree() {
	fmt.Println(n.value)
	for _, child := range n.children {
		fmt.Println(" -", child)
	}
}

func nb_to_out(n Node, nodes map[string]Node, memo map[string]int, visited map[string]bool) int {
	if visited[n.value] {
		return memo[n.value]
	} else if n.value == "out" {
		var result = 1
		memo[n.value] = result
		visited[n.value] = true
		return result
	} else {
		var result = 0
		for _, child := range n.children {
			var child_node = nodes[child]
			result += nb_to_out(child_node, nodes, memo, visited)
		}
		memo[n.value] = result
		visited[n.value] = true
		return result
	}
}

func nb_to_str(n Node, nodes map[string]Node, memo map[string]int, visited map[string]bool, str string) int {
	if visited[n.value] {
		return memo[n.value]
	} else if n.value == str {
		var result = 1
		memo[n.value] = result
		visited[n.value] = true
		return result
	} else {
		var result = 0
		for _, child := range n.children {
			var child_node = nodes[child]
			result += nb_to_str(child_node, nodes, memo, visited, str)
		}
		memo[n.value] = result
		visited[n.value] = true
		return result
	}
}

func part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var nodes = make(map[string]Node)
	var memo = make(map[string]int)
	var visited = make(map[string]bool)
	for _, line := range lines {
		var child = strings.Split(line, ":")
		var node = NewNode(child[0])
		child[1] = strings.TrimPrefix(child[1], " ")
		var subchilds = strings.Split(child[1], " ")

		for _, subchild := range subchilds {
			node.AddChild(subchild)
		}
		nodes[child[0]] = node
	}
	nodes["out"] = NewNode("out")
	var reponse = nb_to_out(nodes["you"], nodes, memo, visited)
	return reponse
}

func part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var nodes = make(map[string]Node)
	var memo_out = make(map[string]int)
	var memo_dac = make(map[string]int)
	var memo_fft = make(map[string]int)

	var visited_out = make(map[string]bool)
	var visited_dac = make(map[string]bool)
	var visited_fft = make(map[string]bool)

	for _, line := range lines {
		var child = strings.Split(line, ":")
		var node = NewNode(child[0])
		child[1] = strings.TrimPrefix(child[1], " ")
		var subchilds = strings.Split(child[1], " ")

		for _, subchild := range subchilds {
			node.AddChild(subchild)
		}
		nodes[child[0]] = node
	}
	nodes["out"] = NewNode("out")

	var dac_to_fft = nb_to_str(nodes["dac"], nodes, memo_fft, visited_fft, "fft")
	var fft_to_dac = nb_to_str(nodes["fft"], nodes, memo_dac, visited_dac, "dac")
	var dac_to_out = nb_to_out(nodes["dac"], nodes, memo_out, visited_out)
	var fft_to_out = nb_to_out(nodes["fft"], nodes, memo_out, visited_out)
	var svr_to_dac = nb_to_str(nodes["svr"], nodes, memo_dac, visited_dac, "dac")
	var svr_to_fft = nb_to_str(nodes["svr"], nodes, memo_fft, visited_fft, "fft")
	var reponse = svr_to_dac*dac_to_fft*fft_to_out + svr_to_fft*fft_to_dac*dac_to_out
	return reponse
}

func main() {
	start := time.Now()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("	AOC 2025 Day11 : ")
	fmt.Println("		-part1 : ", part1(inputDay), "			//made in ", time.Since(start))
	start = time.Now()
	fmt.Println("		-part2 : ", part2(inputDay), "	//made in ", time.Since(start))
	fmt.Println("--------------------------------------------------------------")
}
