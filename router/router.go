package router

import (
	_ "embed"

	"github.com/agnerft/ListRamais/handler"
	"github.com/gin-gonic/gin"
)

var ()

func InitRouter() {
	r := gin.Default()

	r.GET("/cliente", handler.HandleClient)
	r.Static("/assets", "./handler/assets/assets")
	r.Run(":8080")
}