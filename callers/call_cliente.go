package callers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/util"
)

var (
	Cliente *domain.Cliente

	url_padrao = "https://root:agner102030@basesip.makesystem.com.br/clientes?documento="
)

func HandleClient(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Renderizar o formulário HTML
		util.RenderTemplate(w, nil, filepath.Join(os.TempDir(), "/clientes.html"))
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

		util.RenderTemplate(w, cliente, filepath.Join(os.TempDir(), "/clientes.html"))

		return

	}

	// Método não suportado
	http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
}
