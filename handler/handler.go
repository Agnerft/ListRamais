package handler

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/agnerft/ListRamais/domain"
	"github.com/gin-gonic/gin"
)

var (
	Cliente *domain.Cliente
	//go:embed resources
	staticFile embed.FS

	url_padrao = "https://root:agner102030@basesip.makesystem.com.br/clientes?documento="
)

func HandleHomeClient(c *gin.Context) {

	data, err := staticFile.ReadFile("resources/template/clientes.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro para contectar no %s")
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

	// if r.Method == http.MethodPost {
	// 	// Obter a informação do formulário
	// 	cnpj := r.FormValue("cnpj")

	// 	cli := domain.Cliente{}
	// 	cliente, err := cli.RequestJsonCliente(url_padrao + cnpj)

	// 	Cliente = &cliente[0]

	// 	if err != nil {
	// 		log.Fatal("Erro ao buscar os ramais.", err)
	// 	}

	// 	fmt.Println(cliente)
	// 	fmt.Println("ta aqui")

	// util.RenderTemplate(w, cliente, filepath.Join(os.TempDir(), "/clientes.html"))
	// 	util.RenderTemplate(w, cliente, "./main/html/clientes.html")
	// 	return

	// }

	// // Método não suportado
	// http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
}

func HandlePostClient(c *gin.Context) {
	cnpj := c.Request.FormValue("cnpj")

	cli := domain.Cliente{}
	cliente, err := cli.RequestJsonCliente(url_padrao + cnpj)
	if err != nil {
		log.Fatal("Erro ao buscar os ramais.", err)
	}
	Cliente = &cliente[0]

	data, err := staticFile.ReadFile("template/clientes.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro para contectar no %s")
		return
	}

	tmpl, err := template.New("cliente").Parse(string(data))
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro para carregar o template.")
		return
	}
	fmt.Println(cliente)
	err = tmpl.Execute(c.Writer, cliente)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro para executar o template.")
	}

}
