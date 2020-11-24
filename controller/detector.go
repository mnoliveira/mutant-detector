package controller

import (
	"github.com/gin-gonic/gin"
	"mutant-detector/dna"
	"mutant-detector/model"
	"mutant-detector/utils"
	"net/http"
)

func DetectMutant(c *gin.Context) {

	var input model.InputData
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: "json invalido"})
		return
	}

	dnaData, err := utils.ParseInputData(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Error: err.Error()})
		return
	}

	if dna.IsMutant(dnaData) {
		c.JSON(http.StatusOK, model.ResponseOk{Message:"Mutante detectado!"})
	} else {
		c.JSON(http.StatusForbidden, model.ResponseError{Error:"Es solo otro humano."})
	}
}
