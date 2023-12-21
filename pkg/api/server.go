package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	initialiseRoutes(router)
	err := router.Run(":8081")
	if err != nil {
		fmt.Println("Failed top start the API server ")
		return
	}
}
