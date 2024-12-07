package days

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func safeCounter(tolerance int) int {
	raws, ok := utils.ReadInput("input/input_d02.txt")
	if !ok {
		fmt.Println("ERR")
	}

	safeCounter := 0
	for _, raw := range raws {
		tokens := strings.Split(raw, " ")

		decreasing := true
		badLevels := 0
		prec, _ := strconv.Atoi(tokens[0])
		for i := 1; i < len(tokens)-1 && badLevels <= tolerance; i++ {
			cur, _ := strconv.Atoi(tokens[i])

			if i == 1 && prec < cur {
				decreasing = false
			}

			if (math.Abs(float64(cur-prec)) > 3) ||
				(cur == prec) ||
				(cur > prec && decreasing) ||
				(cur < prec && !decreasing) {
				badLevels++
				cur = prec
			}
			prec = cur
		}

		if badLevels <= tolerance {
			safeCounter++
		}
	}

	return safeCounter
}

func d2p1() {
	fmt.Println(safeCounter(0))
}

func d2p2() {
	fmt.Println(safeCounter(1))
}

func D02() {
	d2p1()
	d2p2()
}
