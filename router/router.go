package router

import (
	_ "embed"

	"github.com/agnerft/ListRamais/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	// Configuração do middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8182"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	// Adiciona o middleware CORS ao roteador Gin
	r.Use(cors.New(config))

	r.GET("/", handler.HandleHomeClient)
	r.POST("/cliente", handler.HandlePostClient)
	r.GET("/ramal", handler.HandleSelecionarRamal)
	r.GET("/t", handler.HandleFileConfig)
	// r.Static("/assets", "./handler/assets/assets")

	r.Run(":8080")
}
