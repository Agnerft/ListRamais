package callers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/makesystem/list_ramais/domain"
	"github.com/makesystem/list_ramais/util"
)

var Cliente *domain.Cliente

var url_padrao = "https://root:agner102030@basesip.makesystem.com.br/clientes?documento="

func HandleClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Renderizar o formulário HTML
		util.RenderTemplate(w, nil, "main/assets/clientes.html")
		return

	}

	if r.Method == http.MethodPost {
		// Obter a informação do formulário
		cnpj := r.FormValue("cnpj")

		cli := domain.Cliente{}
		cliente, err := cli.RequestJsonCliente(url_padrao + cnpj)

		Cliente = &cliente[0]

		if err != nil {
			log.Fatal("Erro ao buscar os ramais.", err)
		}

		fmt.Println(cliente)
		fmt.Println("ta aqui")

		util.RenderTemplate(w, cliente, "main/assets/clientes.html")

		return

	}

	// Método não suportado
	http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
}
