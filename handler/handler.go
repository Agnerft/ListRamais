package handler

import (
	"net/http"

	"github.com/agnerft/ListRamais/domain"
	"github.com/gin-gonic/gin"
)

var (
	Cliente *domain.Cliente

	// url_padrao = "https://root:agner102030@basesip.makesystem.com.br/clientes?documento="
)

func HandleClient(c *gin.Context) {

	c.HTML(http.StatusOK, "clientes.html", nil)

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

	// 	// util.RenderTemplate(w, cliente, filepath.Join(os.TempDir(), "/clientes.html"))
	// 	util.RenderTemplate(w, cliente, "./main/html/clientes.html")
	// 	return

	// }

	// // Método não suportado
	// http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
}
