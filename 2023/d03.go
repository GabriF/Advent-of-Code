package main

import (
	. "fmt"
	"unicode"
	"bufio"
	"os"
)

func Simbolo(r rune) bool {
	return r != '.' && ! unicode.IsDigit(r)
}

func ContieneSimbolo(s string) bool {
	for _, r := range s {
		if Simbolo(r) {
			return true
		}
	}

	return false
}

func Sol() int {
	f, err := os.Open("input_d03.txt")
	if err != nil {
		Println("Errore ad aprire il file")
		return -1
	}
	r := bufio.NewScanner(f)

	var righe []string
	for r.Scan() {
		righe = append(righe, r.Text())
	}

	var cifre []string
	for i := 0; i < len(righe); i++ {
		var cifra string
		var valido bool = false
		for j := 0; j < len(righe[i]); j++ {
			carattere := righe[i][j]
			if unicode.IsDigit(rune(carattere)) {
				cifra += string(carattere)

				if ! valido {
					if j + 1 < len(righe[i]) {
						indiceRigaI :=
						_ = indiceRigaI
					}
				}


			} else {
				if valido {
					cifre = append(cifre, cifra)
				}
				cifra = ""
				valido = false
			}
		}
	}

	Println(cifre)
	return 1
}

func main() {
	Sol()
	Println(1 % 2)
}
