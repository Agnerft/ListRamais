package callers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/util"
)

var (
	url                  = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	destDownMicroSIP     = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
)

func HandleInstallMicrosip(w http.ResponseWriter, r *http.Request) {

	if _, err := os.Stat(destDeleleteMicroSIP); err == nil {
		// Se o caminho existe, execute algo
		fmt.Println("o caminho existe")
		err = util.ExecuteUnistall(destDeleleteMicroSIP)
		if err != nil {
			fmt.Printf("Erro ou executar o Desinstalador.")
		}

		err := execute.DownloadGeneric(url, destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao baixar o Arquivo.", err)
		}

		err = util.ExecuteInstall(destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao instalar o %s", destDownMicroSIP)
		}

	} else if os.IsNotExist(err) {
		fmt.Println("o caminho não existe")
		// Se o caminho não existe, faça algo diferente
		// BAIXAR O MICROSIP
		err := execute.DownloadGeneric(url, destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao baixar o Arquivo.", err)
		}

		err = util.ExecuteInstall(destDownMicroSIP)
		if err != nil {
			fmt.Printf("Erro ou executar o Instalador.")
		}

		fmt.Print("Aguardando")

		for {
			for _, r := range `...` {
				time.Sleep(500 * time.Millisecond)
				fmt.Print(string(r))
			}

			// Limpar os pontos
			fmt.Print("\b\b\b   \b\b\b")

			// Aguardar um pouco antes de recomeçar
			time.Sleep(50 * time.Millisecond)
		}

		fmt.Printf("Chamando configuração")
		HandleFileConfig(w, r)

	} else {
		// Algum erro ocorreu ao verificar o caminho
		fmt.Printf("Erro ao verificar o caminho: %v\n", err)
		// Adicione aqui o código para lidar com o erro, se necessário
	}

	fmt.Println("teste")
}

func HandleFileConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ta clicando aqui?")

	fmt.Println(RamalSelecionado)

	destFileConfigMicrosip := filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")

	err := util.AdicionarConfiguracao(destFileConfigMicrosip)
	if err != nil {
		log.Fatal("Erro ao Adicionar a Configuração. \n", err)
	}

	fmt.Println(string(Cliente.Cliente))
	// EDIÇÃO DO ARQUIVO
	err = util.ReadFile(destFileConfigMicrosip, "accountId=0", "accountId=1", 1)
	if err != nil {
		log.Fatal("Erro para modificar o AccountId. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "videoBitrate=0", "videoBitrate=256", 23)
	if err != nil {
		log.Fatal("Erro para modificar o videoBitrate. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "recordingPath=", "recordingPath="+filepath.Join(util.UserCurrent().HomeDir, "Desktop"), 32)
	if err != nil {
		log.Fatal("Erro para modificar o recordingPath. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "recordingFormat=", "recordingFormat=mp3", 33)
	if err != nil {
		log.Fatal("Erro para modificar o recordingFormat. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "autoAnswer=button", "autoAnswer=all", 37)
	if err != nil {
		log.Fatal("Erro para modificar o autoAnswer. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "denyIncoming=button", "denyIncoming=", 43)
	if err != nil {
		log.Fatal("Erro para modificar o denyIncoming. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "label=", "label="+RamalSelecionado, 106)
	if err != nil {
		log.Fatal("Erro para modificar o Sip no Label. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "server=", "server="+string(Cliente.Link_sip), 107)
	if err != nil {
		log.Fatalf("Erro para setar o link do cliente %s. \n %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "proxy=", "proxy="+string(Cliente.Link_sip), 108)
	if err != nil {
		log.Fatalf("Erro para setar o link do cliente %s. \n %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "domain=", "domain="+string(Cliente.Link_sip), 109)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "username=", "username="+RamalSelecionado, 110)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "password=", "password="+RamalSelecionado+"@abc", 111)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "authID=", "authID="+RamalSelecionado, 112)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

}

// for {
// 	for _, r := range `...` {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Print(string(r))
// 	}

// 	// Limpar os pontos
// 	fmt.Print("\b\b\b   \b\b\b")

// 	// Aguardar um pouco antes de recomeçar
// 	time.Sleep(50 * time.Millisecond)
// }
