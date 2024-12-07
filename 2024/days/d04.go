package days

import (
	"aoc/utils"
	"fmt"
)

func d4p1() {
	raws, ok := utils.ReadInput("input/input_d04.txt")
	if !ok {
		fmt.Println("ERR")
		return
	}

	//total := 0
	for i := 0; i < len(raws); i++ {
		for j := 0; j < len(raws[i]); j++ {
			// check down

		}
	}
}

func D04() {
	d4p1()
}
