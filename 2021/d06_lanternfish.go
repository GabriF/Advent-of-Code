package main

import (
	"aoc"
	"fmt"
	"strconv"
)

func parse(riga string) map[int]int {
	lanternfish := make(map[int]int)
	elems := []rune(riga)
	for i := 0; i < len(elems); i += 2 {
		n, _ := strconv.Atoi(string(elems[i]))
		lanternfish[n]++
	}
	return lanternfish
}

func main() {
	righe, _ := aoc.ReadInput("d06_input.txt")
	lanternfish := parse(righe[0])
	const GIORNI = 256
	for i := 0; i < GIORNI; i++ {
		nuovi := lanternfish[0]
		for j := 0; j <= 7; j++ {
			lanternfish[j] = lanternfish[j+1]
		}
		lanternfish[8] = nuovi
		lanternfish[6] += nuovi
		fmt.Println(lanternfish)
	}

	s := 0
	for _, v := range lanternfish {
		s += v
	}
	fmt.Println(s)
}
