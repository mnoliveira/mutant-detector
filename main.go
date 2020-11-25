package main

import (
	"github.com/gin-gonic/gin"
	"mutant-detector/config"
	"mutant-detector/routes"
	"strconv"
)

func main() {

	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.Register(router)

	listenPort, _ := config.Config.Int("listenPort")
	port := ":" + strconv.Itoa(listenPort)

	router.Run(port)
}