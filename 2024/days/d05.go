package days

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func D05() {
	raws, ok := utils.ReadInput("input/input_d05.txt")
	if !ok {
		fmt.Println("ERR")
		return
	}

	higherThan := map[int][]int{} // `higherThan[i]` contains every number `x` such that `i < x`
	var end int                   // end of first section
	for i := 0; raws[i] != ""; i++ {
		tokens := strings.Split(raws[i], "|")
		n1, _ := strconv.Atoi(tokens[0])
		n2, _ := strconv.Atoi(tokens[1])
		higherThan[n1] = append(higherThan[n1], n2)
		end = i
	}
	end += 2

	totSorted := 0
	totUnsorted := 0
	for i := end; i < len(raws); i++ {
		raw := raws[i]
		tokens := strings.Split(raw, ",")
		x := []int{}

		for j := 0; j < len(tokens); j++ {
			n, _ := strconv.Atoi(tokens[j])
			x = append(x, n)
		}

		less := func(a, b int) bool {
			elem1 := x[a]
			elem2 := x[b]
			upElems := higherThan[elem1]
			for j := 0; j < len(upElems); j++ {
				if upElems[j] == elem2 {
					return true
				}
			}
			return false
		}

		if sorted := sort.SliceIsSorted(x, less); sorted {
			totSorted += x[len(x)/2]
		} else {
			sort.Slice(x, less)
			totUnsorted += x[len(x)/2]
		}

	}
	fmt.Println(totSorted)
	fmt.Println(totUnsorted)
}
