package router

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed *
var static embed.FS

func InitRouter() {
	r := gin.Default()

	r.GET("/cliente", func(c *gin.Context) {
		data, err := static.ReadFile("main/templates/clientes.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Erro para contectar no clientes.html")
			return
		}

		tmpl, err := template.New("cliente").Parse(string(data))
		if err != nil {
			c.String(http.StatusInternalServerError, "Erro para carregar o template.")
			return
		}

		err = tmpl.Execute(c.Writer, nil)
		if err != nil {
			c.String(http.StatusInternalServerError, "Erro para executar o template.")
		}
	})
	r.Run(":8080")
}
