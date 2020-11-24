package dna

import (
	"mutant-detector/database"
	"mutant-detector/model"
)

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
