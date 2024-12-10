package days

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func checkRaw(nums []string, tolerance int) bool {
	fmt.Println("EVAL:", nums)
	decreasing := true
	prec, _ := strconv.Atoi(nums[0])
	for i, num := range nums[1:] {
		curr, _ := strconv.Atoi(num)
		if i == 0 && curr > prec {
			decreasing = false
		}

		if diff := math.Abs(float64(prec - curr)); diff > 3 ||
			diff < 1 ||
			(curr > prec && decreasing) ||
			(curr < prec && (!decreasing)) {
			if tolerance > 0 {
				fmt.Println("REM:", i, num)
				return checkRaw(append(nums[:i], nums[i+1:]...), tolerance-1)
			} else {
				return false
			}
		}

		prec = curr
	}
	return true
}

func safeCounter(tolerance int) int {
	raws, ok := utils.ReadInput("input/input_d02.txt")
	if !ok {
		fmt.Println("ERR")
		return -1
	}

	safeCounter := 0
	for _, raw := range raws {
		tokens := strings.Split(raw, " ")
		if checkRaw(tokens, tolerance) {
			safeCounter++
		} else {
			fmt.Println(tokens)
		}
	}

	return safeCounter
}

func D02() {
	fmt.Println(safeCounter(0))
	fmt.Println(safeCounter(1))
}
