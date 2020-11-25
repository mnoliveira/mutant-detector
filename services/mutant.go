package services

import (
	"mutant-detector/config"
	"mutant-detector/model"
	database "mutant-detector/persistence"
)

var MaxHumanSecuence int
var SecuenceSize int
var IndexChange int

//TODO Asumo que una cadena de 8 digitos iguales son 2 matchs!!
func IsMutant(dna [][]string) bool {

	matchs := 0

	//Seteo las configuraciones en variables una sola vez
	MaxHumanSecuence, _ = config.Config.Int("detector.max_human_secuence")
	SecuenceSize, _ = config.Config.Int("detector.secuence_size")
	IndexChange = SecuenceSize - 1

	//Recorro una vez, se podria ver como que recorro por la diagonal principal
	i := 0
	for i < len(dna) && matchs <= MaxHumanSecuence {

		//Busco en todas las direcciones posibles a partir de ese punto de la diagonal y otras calculadas a partir de ese punto
		//con la idea de poder revisar todas las posibilidades dentro de la matriz.
		searchInAllDirections(dna, i, &matchs)
		i++
	}

	return matchs > MaxHumanSecuence
}

func GetStats() (*model.Stats, error){

	countMutant, err := database.GetMutantCount()
	if err != nil {
		return nil, err
	}

	countHumans, err := database.GetHumanCount()
	if err != nil {
		return nil, err
	}

	var ratio float64
	if countMutant > 0 {
		ratio = float64(countHumans) / float64(countMutant)
	}

	stats := model.Stats{
		CountMutant: countMutant,
		CountHuman: countHumans,
		Ratio: ratio,
	}

	return &stats, nil
}

func searchInAllDirections(dna [][]string, indexDiagonal int, matchs *int) {

	directions := GetDirections()

	i := 0
	for i < len(directions) && *matchs <= MaxHumanSecuence {
		directions[i](dna, indexDiagonal, matchs)
		i++
	}
}
