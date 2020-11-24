package model

type Stats struct {
	CountMutant int64   `json:"count_mutant_dna"`
	CountHuman  int64   `json:"count_human_dna"`
	Ratio       float64 `json:"ratio"`
}
