package callers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/makesystem/list_ramais/domain"
	"github.com/makesystem/list_ramais/util"
)

// var RamalSelecionado *ramal.Ramal

// HandleRamais é o manipulador para a rota /ramais
func HandleRamais(w http.ResponseWriter, r *http.Request) {

	// Fazer a requisição para obter os dados JSON
	ramal := domain.Ramal{}
	ramal.InUse = true
	ramais, err := ramal.RequestJsonRamal("link" + "/status_central")

	if err != nil {
		http.Error(w, "Erro ao obter os ramais", http.StatusInternalServerError)
		return
	}

	fmt.Println(ramais)

	// util.RenderTemplate(w, ramais, "main/assets/ramais.html")

}

func HandleSelecionarRamal(w http.ResponseWriter, r *http.Request) {
	// Obter o SIP da query string
	ramalSelecionado := r.URL.Query().Get("sip")

	// RamalSelecionado.Sip = RamalSelect

	// Implementar a lógica para lidar com o SIP selecionado (por exemplo, armazenar em uma variável global)
	fmt.Println("SIP Selecionado:", ramalSelecionado)

	// Responder ao cliente
	w.WriteHeader(http.StatusOK)

	// url := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"

	// destDown := filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")

	destFile := filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")

	// destFile := filepath.Join("main", "assets", "teste.ini")

	// BAIXAR O MICROSIP
	// err = execute.DownloadMicrosip(url, destDown)
	// if err != nil {
	// 	log.Fatal("Erro ao baixar o Arquivo.", err)
	// }

	err := util.AdicionarConfiguracao(destFile)
	if err != nil {
		log.Fatal("Erro ao inserir o Account. \n", err)
	}

	// fmt.Println(cliente.Cliente.RamaisRegistrados)

	// // EDIÇÃO DO ARQUIVO
	// err = util.ReadFile(destFile, "accountId=1", 2)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o AccountId. \n", err)
	// }

	// err = util.ReadFile(destFile, "videoBitrate=256", 24)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o videoBitrate. \n", err)
	// }

	// err = util.ReadFile(destFile, filepath.Join(util.UserCurrent().HomeDir, "Desktop"), 33)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o recordingPath. \n", err)
	// }

	// err = util.ReadFile(destFile, "recordingFormat=mp3", 34)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o recordingFormat. \n", err)
	// }

	// err = util.ReadFile(destFile, "autoAnswer=all", 38)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o autoAnswer. \n", err)
	// }

	// err = util.ReadFile(destFile, "denyIncoming=", 44)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o denyIncoming. \n", err)
	// }

	// err = util.ReadFile(destFile, "label="+ramalSelecionado, 108)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o Sip no Label. \n", err)
	// }

	// err = util.ReadFile(destFile, "server="+cliente.Cliente.Link_sip, 109)
	// if err != nil {
	// 	log.Fatalf("Erro para setar o link do cliente %s. \n %s", cliente.Cliente.Cliente, err)
	// }

	// err = util.ReadFile(destFile, "proxy="+cliente.Cliente.Link_sip, 110)
	// if err != nil {
	// 	log.Fatalf("Erro para setar o link do cliente %s. \n %s", cliente.Cliente.Cliente, err)
	// }

	// err = util.ReadFile(destFile, "domain="+cliente.Cliente.Link_sip, 111)
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente.Cliente, err)
	// }

	// err = util.ReadFile(destFile, "username="+ramalSelecionado, 112)
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente.Cliente, err)
	// }

	// err = util.ReadFile(destFile, "password="+RamalSelecionado.Sip+"@abc", 113)
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente.Cliente, err)
	// }

	// err = util.ReadFile(destFile, "authID="+RamalSelecionado.Sip, 114)
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente.Cliente, err)
	// }

	util.RenderTemplate(w, ramalSelecionado, "main/assets/ramalSelecionado.html")
}
