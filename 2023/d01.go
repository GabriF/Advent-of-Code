package main

import (
	. "fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
)

func sol(p2 bool) int {
	f, err := os.Open("input_d01.txt")
	if err != nil {
		Println(err)
		return -1
	}

	r := bufio.NewScanner(f)

	var s int
	for r.Scan() {
		riga := r.Text()

		numeroString := estraiCifre(riga, p2)
		numeroInt, _ := strconv.Atoi(numeroString)
		s += numeroInt
	}

	return s
}

func p1() int {
	return sol(false)
}

func p2() int {
	return sol(true)
}

func estraiCifre(s string, scritte bool) string {
	var dig1, dig2 string

	var mappaCifre map[string]string = map[string]string{"one" : "1", "two" : "2", "three": "3", "four" : "4", "five" : "5", "six" : "6", "seven" : "7", "eight" : "8", "nine" : "9"}

	for i := 0; i < len(s); i++ {
		var cifra string

		if unicode.IsDigit(rune(s[i])) {
			cifra = string(s[i])
		}

		if cifra == "" && scritte {
			for j := 3; j < 6 && i + j <= len(s) && cifra == ""; j++ {
				sub := string(s[i:i+j])
				cifra = mappaCifre[sub]

				if cifra != "" {
					Println(i + j, cifra, s)
					i += j - 2
				}
			}
		}

		if cifra != "" {
			if dig1 == "" {
				dig1 = cifra
			}

			dig2 = cifra
		}
	}

	return dig1 + dig2
}

func main() {
	Println(p1())
	Println(p2())
}
