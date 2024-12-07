package utils

import (
	"bufio"
	"os"
)

func ReadInput(fileName string) ([]string, bool) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, false
	}
	r := bufio.NewScanner(f)

	righe := []string{}
	for r.Scan() {
		linea := r.Text()
		righe = append(righe, linea)

	}

	return righe, true
}
