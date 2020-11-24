package dna

import (
	"mutant-detector/model"
)

/*La idea general es que al recorrer primero validar que todavia se pueda formar la secuencia a partir de la posicion actual
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
	for j+INDEX_CHANGE < len(dna) && *matchs <= LIMIT_SECUENCE {

		k := j + INDEX_CHANGE
		for k > j && dna[i][k-1] == dna[i][k] {
			k--
		}

		if k == j {
			//fmt.Println("Encontre un match horizontal.")
			*matchs++
			k += SECUENCE_SIZE
		}

		j = k
	}
}

func findVerticalMatchs(dna model.DNA, j int, matchs *int) {

	//Busco secuencias en la j-esima columna
	i := 0
	for i+INDEX_CHANGE < len(dna) && *matchs <= LIMIT_SECUENCE {

		k := i + INDEX_CHANGE
		for k > i && dna[k-1][j] == dna[k][j] {
			k--
		}

		if k == i {
			//fmt.Println("Encontre un match vertical.")
			*matchs++
			k += SECUENCE_SIZE
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
	if indexDiagonal > 0 && *matchs <= LIMIT_SECUENCE {
		indexRow = indexDiagonal
		indexColumn = 0

		searchMatchsObliqueLR(dna, indexRow, indexColumn, matchs)
	}
}

func searchMatchsObliqueLR(dna model.DNA, i int, j int, matchs *int) {

	//Recorre una diagonal de izquierda a derecha buscando secuencias
	for j+INDEX_CHANGE < len(dna) && i+INDEX_CHANGE < len(dna) && *matchs <= LIMIT_SECUENCE {

		k := i + INDEX_CHANGE
		l := j + INDEX_CHANGE
		for k > i && l > j && dna[k-1][l-1] == dna[k][l] {
			k--
			l--
		}

		if k == i && l == j {
			//fmt.Println("Encontre un match oblicuo de izquierda a derecha.")
			*matchs++
			k += SECUENCE_SIZE
			l += SECUENCE_SIZE
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
	if indexDiagonal > 0 && *matchs <= LIMIT_SECUENCE {
		indexRow = indexDiagonal
		indexColumn = len(dna) - 1

		searchMatchsObliqueRL(dna, indexRow, indexColumn, matchs)
	}
}

func searchMatchsObliqueRL(dna model.DNA, i int, j int, matchs *int) {

	//Recorre una diagonal de derecha a izquierda buscando secuencias
	for j-INDEX_CHANGE >= 0 && i+INDEX_CHANGE < len(dna) && *matchs <= LIMIT_SECUENCE {

		k := i + INDEX_CHANGE
		l := j - INDEX_CHANGE
		for k > i && l < j && dna[k-1][l+1] == dna[k][l] {
			k--
			l++
		}

		if k == i && l == j {
			//fmt.Println("Encontre un match oblicuo de derecha a izquierda.")
			*matchs++
			k += SECUENCE_SIZE
			l -= SECUENCE_SIZE
		}

		i = k
		j = l
	}
}