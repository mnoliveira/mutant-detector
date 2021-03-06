package routes

import (
	"github.com/gin-gonic/gin"
	"mutant-detector/controller"
	"net/http"
)

func Register(router *gin.Engine){

	router.GET("/ping", ping)

	router.POST("/mutant", controller.DetectMutant)
	router.GET("/stats", controller.Stats)

}

func ping(c *gin.Context){
	c.String(http.StatusOK, "ok")
}