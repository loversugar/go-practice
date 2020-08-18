package main

import (
	"github.com/gin-gonic/gin"
	"go-practice/GinProject/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterApiRouter(router)

	router.Run()
}
