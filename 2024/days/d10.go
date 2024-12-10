package days

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func buildMatrix(raws []string) [][]int {
	m := make([][]int, len(raws))
	for i, raw := range raws {
		m[i] = make([]int, len(raw))
		for j, char := range raw {
			num, _ := strconv.Atoi(string(char))
			m[i][j] = num
		}
	}
	return m
}

func countRating(x, y int, m [][]int) int {
	if m[x][y] == 9 {
		return 1
	}
	curr := m[x][y]
	s := 0
	if up := x + 1; up < len(m) && m[up][y] == curr+1 {
		s += countRating(up, y, m)
	}
	if down := x - 1; down >= 0 && m[down][y] == curr+1 {
		s += countRating(down, y, m)
	}
	if right := y + 1; right < len(m[x]) && m[x][right] == curr+1 {
		s += countRating(x, right, m)
	}
	if left := y - 1; left >= 0 && m[x][left] == curr+1 {
		s += countRating(x, left, m)
	}
	return s
}

func countScore(x, y int, reached [][]bool, m [][]int) int {
	if m[x][y] == 9 {
		if reached[x][y] {
			return 0
		} else {
			reached[x][y] = true
			return 1
		}
	}
	curr := m[x][y]
	s := 0
	if up := x + 1; up < len(m) && m[up][y] == curr+1 {
		s += countScore(up, y, reached, m)
	}
	if down := x - 1; down >= 0 && m[down][y] == curr+1 {
		s += countScore(down, y, reached, m)
	}
	if right := y + 1; right < len(m[x]) && m[x][right] == curr+1 {
		s += countScore(x, right, reached, m)
	}
	if left := y - 1; left >= 0 && m[x][left] == curr+1 {
		s += countScore(x, left, reached, m)
	}
	return s
}

func D10() {
	raws, ok := utils.ReadInput("input/input_d10.txt")
	if !ok {
		fmt.Println("ERR")
		return
	}
	m := buildMatrix(raws)
	p1, p2 := 0, 0
	for i, raw := range m {
		for j, num := range raw {
			if num == 0 {
				reached := make([][]bool, len(m))
				for z := 0; z < len(m); z++ {
					reached[z] = make([]bool, len(m[z]))
				}
				p1 += countScore(i, j, reached, m)
				p2 += countRating(i, j, m)
			}
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
