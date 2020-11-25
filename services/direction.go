package services

import (
	"mutant-detector/model"
)

/*La idea general es que al recorrer cada direccion, primero valide que todavia se pueda formar la secuencia a partir de la posicion actual
y saltar a dicha posicion volviendo hacia atras, asi ahorro recorrer uno por uno todos los elementos todo el tiempo.
Un dato que se tuvo en cuenta es poder cambiar la longitud de la secuencia que se busca solo cambiando el valor de SECUENCE_SIZE
*/

//Posible mejora: Generalizar un poco algunas busquedas ya que solo cambia como incrementan o decrementan los indices

func GetDirections() []model.Direction {
	return []model.Direction{findHorizontalMatchs, findVerticalMatchs, findObliqueRightToLeftMatchs, findObliqueLeftToRightMatchs}
}

func findHorizontalMatchs(dna model.DNA, i int, matchs *int) {

	//Busco secuencias en la i-esima fila
	j := 0
	for j+IndexChange < len(dna) && *matchs <= MaxHumanSecuence {

		k := j + IndexChange
		for k > j && dna[i][k-1] == dna[i][k] {
			k--
		}

		if k == j {
			*matchs++
			k += SecuenceSize
		}

		j = k
	}
}

func findVerticalMatchs(dna model.DNA, j int, matchs *int) {

	//Busco secuencias en la j-esima columna
	i := 0
	for i+IndexChange < len(dna) && *matchs <= MaxHumanSecuence {

		k := i + IndexChange
		for k > i && dna[k-1][j] == dna[k][j] {
			k--
		}

		if k == i {
			*matchs++
			k += SecuenceSize
		}

		i = k
	}
}

//TODO Cuando hablamos de diagonales siempre asumimos que se recorren de arriba hacia abajo, aclarando solo si es de izquiera a derecha "\" o de derecha a izquierda "/"

func findObliqueLeftToRightMatchs(dna model.DNA, indexDiagonal int, matchs *int) {

	//A partir del indice obtengo 2 diagonales de izquierda a derecha y busco secuencias
	//Al ir aumentando el indexDiagonal empieza desde las diagonales mas cercanas a la principal, hasta las mas lejanas
	indexRow := 0
	indexColumn := indexDiagonal

	searchMatchsObliqueLR(dna, indexRow, indexColumn, matchs)

	//Si esta en la primer posicion, ya recorri la unica diagonal posible
	if indexDiagonal > 0 && *matchs <= MaxHumanSecuence {
		indexRow = indexDiagonal
		indexColumn = 0

		searchMatchsObliqueLR(dna, indexRow, indexColumn, matchs)
	}
}

func searchMatchsObliqueLR(dna model.DNA, i int, j int, matchs *int) {

	//Recorre una diagonal de izquierda a derecha buscando secuencias
	for j+IndexChange < len(dna) && i+IndexChange < len(dna) && *matchs <= MaxHumanSecuence {

		k := i + IndexChange
		l := j + IndexChange
		for k > i && l > j && dna[k-1][l-1] == dna[k][l] {
			k--
			l--
		}

		if k == i && l == j {
			*matchs++
			k += SecuenceSize
			l += SecuenceSize
		}

		i = k
		j = l
	}
}

func findObliqueRightToLeftMatchs(dna model.DNA, indexDiagonal int, matchs *int) {

	//A partir del indice obtengo 2 diagonales de derecha a izquierda y busco secuencias
	//Al ir aumentando el indexDiagonal empieza desde las diagonales mas cercanas a la secundaria, hasta las mas lejanas
	indexRow := 0
	indexColumn := len(dna) - 1 - indexDiagonal

	searchMatchsObliqueRL(dna, indexRow, indexColumn, matchs)

	//Si esta en la primer posicion, ya recorri la unica diagonal posible
	if indexDiagonal > 0 && *matchs <= MaxHumanSecuence {
		indexRow = indexDiagonal
		indexColumn = len(dna) - 1

		searchMatchsObliqueRL(dna, indexRow, indexColumn, matchs)
	}
}

func searchMatchsObliqueRL(dna model.DNA, i int, j int, matchs *int) {

	//Recorre una diagonal de derecha a izquierda buscando secuencias
	for j-IndexChange >= 0 && i+IndexChange < len(dna) && *matchs <= MaxHumanSecuence {

		k := i + IndexChange
		l := j - IndexChange
		for k > i && l < j && dna[k-1][l+1] == dna[k][l] {
			k--
			l++
		}

		if k == i && l == j {
			*matchs++
			k += SecuenceSize
			l -= SecuenceSize
		}

		i = k
		j = l
	}
}