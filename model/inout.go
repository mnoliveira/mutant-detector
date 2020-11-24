package model

type InputData struct {
	DNA []string `json:"dna"`
}

type ResponseOk struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Error string `json:"error"`
}