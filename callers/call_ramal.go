package callers

import (
	"fmt"
	"net/http"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/util"
)

var (

	// Cliente_ramal domain.Cliente
	ramal            = domain.Ramal{}
	RamalSelecionado = ramal.Sip
)

// HandleRamais é o manipulador para a rota /ramais
func HandleRamais(w http.ResponseWriter, r *http.Request) {

	// Fazer a requisição para obter os dados JSON

	ramal.InUse = true
	ramais, err := ramal.RequestJsonRamal(Cliente.Link + "/status_central")

	if err != nil {
		http.Error(w, "Erro ao obter os ramais", http.StatusInternalServerError)
		return
	}

	fmt.Println(ramais)

	util.RenderTemplate(w, ramais, "./util/assets/ramais.html")

}

func HandleSelecionarRamal(w http.ResponseWriter, r *http.Request) {
	// Obter o SIP da query string
	RamalSelecionado = r.URL.Query().Get("sip")

	// Implementar a lógica para lidar com o SIP selecionado (por exemplo, armazenar em uma variável global)
	fmt.Println("SIP Selecionado:", RamalSelecionado)

	// Responder ao cliente
	w.WriteHeader(http.StatusOK)

	util.RenderTemplate(w, RamalSelecionado, "util/assets/ramalSelecionado.html")
}
