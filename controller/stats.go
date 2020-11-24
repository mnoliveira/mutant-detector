package controller

import (
	"github.com/gin-gonic/gin"
	"mutant-detector/dna"
	"mutant-detector/model"
	"net/http"
)

func Stats(c *gin.Context) {

	stats, err := dna.GetStats()

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{Error: err.Error() })
	} else {
		c.JSON(http.StatusOK, stats)
	}
}
