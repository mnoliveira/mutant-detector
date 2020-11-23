package main

import (
	"fmt"
	"mutant-detector/api"
)

func main() {

	mutanteInput := []string{	"ATGCGA",
								"CAGTGC",
								"TTATGT",
								"AGAAGG",
								"CCCCTA",
								"TCACTG"}

	mutanteDiagRLInput := []string{ "TTGCCA",
									"CAGTGC",
									"TTATCT",
									"AGACTG",
									"CACGTA",
									"TTTTGG"}

	mutanteDiagLRInput := []string{ "TAGCCA",
		  							"CAATGC",
		  							"TTAAGT",
		  							"AGAAAG",
		  							"CCGCTT",
		  							"TTTTGG"}

	humanoInput := []string{"TTGCCA","CAGTGC","TTATGT","AGAAGG","CCGCTA","TCACTG"}

	mutanteDNA := api.MakeMatrix(mutanteInput)

	mutanteDiagRLDNA := api.MakeMatrix(mutanteDiagRLInput)

	mutanteDiagLRDNA := api.MakeMatrix(mutanteDiagLRInput)

	humanoDNA := api.MakeMatrix(humanoInput)

	fmt.Printf("El mutante es mutante...? %v \n", api.IsMutant(mutanteDNA))

	fmt.Printf("El mutante diagonal RL es mutante...? %v \n", api.IsMutant(mutanteDiagRLDNA))

	fmt.Printf("El mutante diagonal LR es mutante...? %v \n", api.IsMutant(mutanteDiagLRDNA))

	fmt.Printf("El humano es mutante...? %v \n", api.IsMutant(humanoDNA))
}

func printMatrix(matrix [][]string) {
	fmt.Println("---------------------------------------------------")

	for _, fila := range matrix {
		for _, item := range fila {
			fmt.Print(item + " ")
		}
		fmt.Println()
	}
}