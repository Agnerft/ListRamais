package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/makesystem/list_ramais/callers"
)

func main() {
	// http://workstation.gvctelecom.com.br:1139/painel.php ok
	// http://pires.gvctelecom.com.br:1155/painel.php ok
	// http://mscelular.gvctelecom.com.br:1133/painel.php ok
	// http://msb2b.gvctelecom.com.br:1127/painel.php ok
	// http://rs.gvctelecom.com.br:1079/painel.php ok
	// http://vitale.gvctelecom.com.br:1191/painel.php ok
	// http://jrsolution.gvctelecom.com.br:5131/painel.php ok
	// http://umusul.gvctelecom.com.br:1135/painel.php ok
	// http://clik.gvctelecom.com.br:1137/painel.php ok

	// url := "http://clik.gvctelecom.com.br:1137/status_central"

	http.HandleFunc("/cliente", callers.HandleClient)
	// http.HandleFunc("/ramais", callers.HandleRamais)
	// http.HandleFunc("/selecionar-sip", callers.HandleSelecionarRamal)
	// http.HandleFunc("/acao", config.HandleConfigInstall)

	fmt.Println("Servidor Rodando")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
