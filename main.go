package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agnerft/ListRamais/callers"
	"github.com/agnerft/ListRamais/util"
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

	// err := execute.DownloadGeneric("https://raw.githubusercontent.com/Agnerft/ListRamais/main/util/assets/clientes.html", os.TempDir())
	// if err != nil {
	// 	fmt.Println("Erro para baixar o cliente.html")
	// }
	// err = execute.DownloadGeneric("https://raw.githubusercontent.com/Agnerft/ListRamais/main/util/assets/ramais.html", os.TempDir())
	// if err != nil {
	// 	fmt.Println("Erro para baixar o cliente.html")
	// }
	// err = execute.DownloadGeneric("https://raw.githubusercontent.com/Agnerft/ListRamais/main/util/assets/ramalSelecionado.html", os.TempDir())
	// if err != nil {
	// 	fmt.Println("Erro para baixar o cliente.html")
	// }

	http.HandleFunc("/cliente", callers.HandleClient)
	http.HandleFunc("/ramais", callers.HandleRamais)
	http.HandleFunc("/selecionar-sip", callers.HandleSelecionarRamal)
	http.HandleFunc("/acaoyes", callers.HandleFileConfig)
	http.HandleFunc("/acaono", callers.HandleInstallMicrosip)

	fmt.Println("Servidor Rodando")

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Abrir o navegador padrão automaticamente
	util.OpenBrowser("http://localhost:8080/cliente")
	// Manter o programa em execução
	select {}
}
