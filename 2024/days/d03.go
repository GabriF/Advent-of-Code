package days

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func D03() {
	raws, ok := utils.ReadInput("input/input_d03.txt")
	if !ok {
		fmt.Println("ERR")
		return
	}

	do := true
	p1 := 0
	p2 := 0
	for _, raw := range raws {
		instructions := regexp.MustCompile("mul\\([0-9]{1,3}\\,[0-9]{1,3}\\)|don't\\(\\)|do\\(\\)").FindAllString(raw, -1)
		for _, instr := range instructions {
			fmt.Println(instr)
			if instr == "do()" {
				do = true
			} else if instr == "don't()" {
				do = false
			} else { // mul()
				args := strings.Split(instr[4:len(instr)-1], ",")
				n1, _ := strconv.Atoi(args[0])
				n2, _ := strconv.Atoi(args[1])
				p1 += n1 * n2
				if do {
					p2 += n1 * n2
				}
			}
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
