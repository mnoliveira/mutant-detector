package main

import (
	"github.com/gin-gonic/gin"
	"mutant-detector/routes"
)

func main() {

	router := gin.Default()
	gin.SetMode("debug")

	routes.Register(router)

	router.Run(":5000")
}