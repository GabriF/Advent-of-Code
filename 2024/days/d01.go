package days

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func d1p1() {
	raws, ok := utils.ReadInput("input/input_d01.txt")
	if !ok {
		fmt.Println("ERR")
	}

	s1 := make([]int, len(raws))
	s2 := make([]int, len(raws))
	for i, r := range raws {
		tokens := strings.Split(r, "   ")
		s1[i], _ = strconv.Atoi(tokens[0])
		s2[i], _ = strconv.Atoi(tokens[1])
	}

	sort.Ints(s1)
	sort.Ints(s2)

	s := 0
	for i := 0; i < len(s1); i++ {
		s += int(math.Abs(float64(s1[i] - s2[i])))
	}

	fmt.Println(s)
}

func d1p2() {
	raws, ok := utils.ReadInput("input/input_d01.txt")
	if !ok {
		fmt.Println("ERR")
	}

	s1 := make(map[int]int, len(raws))
	s2 := make(map[int]int, len(raws))
	for _, raw := range raws {
		tokens := strings.Split(raw, "   ")
		val1, _ := strconv.Atoi(tokens[0])
		val2, _ := strconv.Atoi(tokens[1])
		s1[val1]++
		s2[val2]++
	}

	total := 0
	for key, val := range s2 {
		if count, ok := s1[key]; ok {
			total += key * val * count
		}
	}

	fmt.Println(total)
}

func D01() {
	d1p1()
	d1p2()
}
