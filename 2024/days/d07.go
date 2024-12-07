package days

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func concNumbers(n1, n2 int) int {
	return n1*int(math.Pow10(int(math.Log10(float64(n2))+1))) + n2
}

func canCalcHelper(numbers []int, curr int, i int, test int, conc bool) bool {
	if i == len(numbers)-1 {
		return curr == test
	} else if curr > test {
		return false
	}

	return canCalcHelper(numbers, curr+numbers[i+1], i+1, test, conc) ||
		canCalcHelper(numbers, curr*numbers[i+1], i+1, test, conc) ||
		(conc && canCalcHelper(numbers, concNumbers(curr, numbers[i+1]), i+1, test, conc))
}

func canCalc(numbers []int, test int, conc bool) bool {
	return canCalcHelper(numbers, numbers[0], 0, test, conc)
}

func D07() {
	raws, ok := utils.ReadInput("input/input_d07.txt")
	if !ok {
		fmt.Println("ERR")
		return
	}

	p1, p2 := 0, 0
	for _, raw := range raws {
		tokens := strings.Split(raw, " ")
		testNumber, _ := strconv.Atoi(strings.Trim(tokens[0], ":"))
		numbers := make([]int, len(tokens)-1)
		for i := 1; i < len(tokens); i++ {
			numbers[i-1], _ = strconv.Atoi(tokens[i])
		}

		if canCalc(numbers, testNumber, false) {
			p1 += testNumber
		}

		if canCalc(numbers, testNumber, true) {
			p2 += testNumber
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
