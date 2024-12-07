package main

import (
	"fmt"
	"strings"
	"strconv"
	"aoc"
)

type Punto struct {
	x int
	y int
}

type Istruzione struct {
	dir rune
	val int
}

func parse() ([]Punto, []Istruzione) {
	punti := []Punto{}
	istruzioni := []Istruzione{}

	righe, _ := aoc.LeggiInput("d13_input.txt")

	leggiIstruzioni := false
	for _, linea := range righe {
		if linea == "" {
			leggiIstruzioni = true
			continue
		}
		if ! leggiIstruzioni {
			lineaSplit := strings.Split(linea, ",")
			x, _ := strconv.Atoi(lineaSplit[0])
			y, _ := strconv.Atoi(lineaSplit[1])
			punti = append(punti, Punto{x, y})
		} else {
			lineaSplit := strings.Split(linea, " ")
			istruzione := strings.Split(lineaSplit[2], "=")
			dir := rune(istruzione[0][0])
			val, _ := strconv.Atoi(istruzione[1])
			istruzioni = append(istruzioni, Istruzione{dir, val})
		}

	}
	return punti, istruzioni
}

func piega(p Punto, i Istruzione) Punto {
	dir := i.dir
	val := i.val
	x := p.x
	y := p.y
	if dir == 'x' && x > val {
		x = 2 * val - x
	} else if dir == 'y' && y > val {
		y = 2 * val - y
	}
	return Punto{x, y}
}

func stampa(punti map[Punto]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 50; j++ {
			if punti[Punto{j, i}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func p1() {
	punti, istruzioni := parse()
	istr := istruzioni[0]
	puntiFold := map[Punto]bool{}
	for _, p := range punti {
		p = piega(p, istr)
		puntiFold[p] = true
	}

	fmt.Println(len(puntiFold))
}

func p2() {
	punti, istruzioni := parse()
	puntiFold := map[Punto]bool{}
	for _, p := range punti {
		for _, istr := range istruzioni {
			p = piega(p, istr)
		}
		puntiFold[p] = true
	}
	stampa(puntiFold)
}

func main() {
	p1()
	p2()
}
