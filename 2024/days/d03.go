package days

import (
	"aoc/utils"
	"fmt"
)

type State int

const (
	FindFunction State = iota
	ReadArgs
)

func d3p1() {
	_, ok := utils.ReadInput("input/input_d03.txt")
	if !ok {
		fmt.Println("ERR")
		return
	}
}

func D03() {
	d3p1()
}
