package api

const SECUENCE_SIZE = 4
const INDEX_CHANGE = SECUENCE_SIZE - 1
const LIMIT_SECUENCE = 1

//TODO Asumo que una cadena de 8 digitos iguales son 2 matchs!!

func IsMutant(dna [][]string) bool {

	matchs := 0

	//Recorro una vez, se podria ver como que recorro por la diagonal principal
	i := 0
	for i < len(dna) && matchs <= LIMIT_SECUENCE {

		//Busco en todas las direcciones posibles a partir de ese punto de la diagonal y otras calculadas a partir de ese punto
		//con la idea de poder revisar todas las posibilidades dentro de la matriz.
		searchInAllDirections(dna, i, &matchs)
		i++
	}

	return matchs > LIMIT_SECUENCE
}

func searchInAllDirections(dna [][]string, indexDiagonal int, matchs *int) {

	directions := GetDirections()

	i := 0
	for i < len(directions) && *matchs <= LIMIT_SECUENCE {
		directions[i](dna, indexDiagonal, matchs)
		i++
	}
}
