package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mutant-detector/config"
	"mutant-detector/model"
	database "mutant-detector/persistence"
	"mutant-detector/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	config.TestMode = true
	config.LoadConfig()

	router = gin.New()
	routes.Register(router)

	code := m.Run()
	database.TestDropDBTest()
	os.Exit(code)
}

func TestPing(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Ping no responde")
	}
}

func TestMutant(t *testing.T) {

	input := model.InputData{ DNA: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}

	w := doMutantPost(input)

	if w.Code != http.StatusOK {
		t.Errorf("Es mutante pero detecta que es humano.")
	}
}

func TestMutantVerticalHorizontal(t *testing.T) {

	input := model.InputData{ DNA: []string{"TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}

	w := doMutantPost(input)

	if w.Code != http.StatusOK {
		t.Errorf("Es mutante pero detecta que es humano.")
	}
}

func TestMutantDiagonal(t *testing.T) {

	input := model.InputData{ DNA: []string{"TTGCCA", "CAGTAC", "TATAGT", "AGATGG", "CCGCTA", "TCACTT"}}

	w := doMutantPost(input)

	if w.Code != http.StatusOK {
		t.Errorf("Es mutante pero detecta que es humano.")
	}
}

func TestMutantInvalidMatrixValues(t *testing.T) {

	input := model.InputData{ DNA: []string{"PTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}

	w := doMutantPost(input)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Tiene digitos invalidos y no los detecto.")
	}
}

func TestMutantInvalidMatrixDimensions(t *testing.T) {

	input := model.InputData{ DNA: []string{"ATGCGAA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}

	w := doMutantPost(input)

	if w.Code != http.StatusBadRequest {
		t.Errorf("No es cuadrada y no lo detecto.")
	}
}

func TestHuman(t *testing.T) {

	input := model.InputData{ DNA: []string{"TTGCCA","CAGTGC","TTATGT","AGAAGG","CCGCTA","TCACTG"}}

	w := doMutantPost(input)

	if w.Code != http.StatusForbidden {
		t.Errorf("Es humano pero detecta que es mutante.")
	}
}

func TestStats(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/stats", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Stats no responde")
	}
}

func doMutantPost(data model.InputData) *httptest.ResponseRecorder {

	byte, _ := json.Marshal(data)
	body := bytes.NewReader(byte)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/mutant", body)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	return w
}