package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"strconv"
)

func EstraiID(s string) int {
	numeroInc := strings.Split(s, " ")[1]
	numeroString := numeroInc[:len(numeroInc) - 1]
	numeroInt, _ := strconv.Atoi(numeroString)

	return numeroInt
}

func VerificaPartita(s string, MAX map[string]int) (bool, int) {
	var valido bool = true

	sets := strings.Split(s, "; ")
	minimi := map[string]int{"red": 0, "green" : 0, "blue": 0}
	for _, s := range sets {
		for _, cubi := range strings.Split(s, ", ") {
			cubiS := strings.Split(cubi, " ")
			numero, _ := strconv.Atoi(cubiS[0])
			colore := cubiS[1]

			if numero > MAX[colore] {
				valido = false
			}

			if numero > minimi[colore] {
				minimi[colore] = numero
			}
		}
	}

	return valido, minimi["red"] * minimi["green"] * minimi["blue"]
}

func Sol() (int, int) {
	f, err := os.Open("input_d02.txt")
	if err != nil {
		Println("Errore ad aprire il file")
		return -1, -1
	}
	r := bufio.NewScanner(f)

	var p1, p2 int
	MAX := map[string]int{"red": 12, "green": 13, "blue": 14}
	for r.Scan() {
		riga := r.Text()

		id := EstraiID(riga)
		partita := strings.SplitN(riga, ": ", 2)[1]
		valido, power := VerificaPartita(partita, MAX)

		if valido {
			p1 += id
		}
		p2 += power

	}

	return p1, p2
}

func main() {
	Println(Sol())
}
