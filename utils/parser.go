package utils

import (
	"errors"
	"mutant-detector/model"
	"regexp"
	"strings"
)

func ParseInputData(data model.InputData) (model.DNA, error) {

	var dna model.DNA
	dna = make([][]string, len(data.DNA))

	for i, fila := range data.DNA {

		if err := validate(fila, data); err != nil {
			return dna, err
		}

		dna[i] = strings.Split(fila, "")
	}

	return dna, nil
}

func validate(fila string, data model.InputData) error {

	if len(data.DNA) != len(fila) {
		return errors.New("Dimensiones invalidas!")
	}

	if matched, err := regexp.MatchString("^[ACGT]*$", fila); err != nil || !matched {
		return errors.New("Cadena de ADN invalida!")
	}

	return nil
}