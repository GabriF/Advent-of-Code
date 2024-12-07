package main

import (
	"fmt"
	"aoc"
	"strconv"
	"github.com/golang-collections/collections/stack"
	"slices"
)

type Punto struct {
	x int
	y int
}

func parse() [][]int {
	raws, _ := aoc.LeggiInput("d09_input.txt")
	heights := make([][]int, 0, 0)
	for _, raw:= range raws {
		r := make([]int, 0, 0)
		for _, char := range raw {
			val, _ := strconv.Atoi(string(char))
			r = append(r, val)
		}
		heights = append(heights, r)
	}
	return heights
}

func checkLow(m [][]int, i int, j int) bool {
	val := m[i][j]
	if val == 9 {
		return false
	}
	up := i - 1
	down := i + 1
	left := j - 1
	right := j + 1
	if up >= 0 && val >= m[up][j] {
		return false
	}
	if down < len(m) && val >= m[down][j] {
		return false
	}
	if left >= 0 && val >= m[i][left] {
		return false
	}
	if right < len(m[i]) && val >= m[i][right] {
		return false
	}
	return true
}

/*
 * Restituisce il numero di elementi del bacino
 * ammesso che m[i][j] sia un punto di rischio
 */

func checkBasin(m [][]int, i int, j int) int {
	return checkBasinHelp(m, Punto{i, j}, map[Punto]bool{})
}

func checkBasinHelp(m [][]int, p Punto, visited map[Punto]bool) int {
	s := stack.New()
	s.Push(p)
	size := 0;
	for s.Len() > 0 {
		act := s.Pop().(Punto)
		i, j := act.x, act.y
		if ! visited[act] && m[act.x][act.y] != 9 {
			size += 1
			if i - 1 >= 0 {
				s.Push(Punto{i - 1, j})
			}
			if j - 1 >= 0 {
				s.Push(Punto{i, j - 1})
			}
			if i + 1 < len(m) {
				s.Push(Punto{i + 1, j})
			}
			if j + 1 < len(m[i]) {
				s.Push(Punto{i, j + 1})
			}
			visited[act] = true
		}
	}
	return size;
}

func p1() {
	riskLevel := 0
	heights := parse()
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if checkLow(heights, i, j) {
				riskLevel += heights[i][j] + 1
			}
		}
	}
	fmt.Println(riskLevel)
}

func p2() {
	basins := []int{}
	heights := parse()
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if checkLow(heights, i, j) {
				b := checkBasin(heights, i, j)
				basins = append(basins, b)
			}
		}
	}
	slices.Sort(basins)
	fmt.Println(basins[len(basins) - 1] * basins[len(basins) - 2] * basins[len(basins) - 3])
}

func main() {
	p1()
	p2()
}
