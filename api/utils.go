package api

import "strings"

func MakeMatrix(filas []string) [][]string {

	matrix := make([][]string, len(filas))

	for i, fila := range filas {
		matrix[i] = strings.Split(fila, "")
	}

	return matrix
}