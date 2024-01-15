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

	} else if os.IsNotExist(err) {
		fmt.Println("o caminho não existe")
		// Se o caminho não existe, faça algo diferente
		// BAIXAR O MICROSIP
		err := execute.DownloadMicrosip(url, destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao baixar o Arquivo.", err)
		}

		err = util.ExecuteInstall(destDeleleteMicroSIP)
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
			time.Sleep(500 * time.Millisecond)
		}

		// fmt.Printf("Executando a configuração no caminho %s", destFileConfigMicrosip)

	} else {
		// Algum erro ocorreu ao verificar o caminho
		fmt.Printf("Erro ao verificar o caminho: %v\n", err)
		// Adicione aqui o código para lidar com o erro, se necessário
	}

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
	err = util.ReadFile(destFileConfigMicrosip, "accountId=1", 2)
	if err != nil {
		log.Fatal("Erro para modificar o AccountId. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "videoBitrate=256", 24)
	if err != nil {
		log.Fatal("Erro para modificar o videoBitrate. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "recordingPath="+filepath.Join(util.UserCurrent().HomeDir, "Desktop"), 33)
	if err != nil {
		log.Fatal("Erro para modificar o recordingPath. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "recordingFormat=mp3", 34)
	if err != nil {
		log.Fatal("Erro para modificar o recordingFormat. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "autoAnswer=all", 38)
	if err != nil {
		log.Fatal("Erro para modificar o autoAnswer. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "denyIncoming=", 44)
	if err != nil {
		log.Fatal("Erro para modificar o denyIncoming. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "label="+RamalSelecionado, 108)
	if err != nil {
		log.Fatal("Erro para modificar o Sip no Label. \n", err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "server="+string(Cliente.Link_sip), 109)
	if err != nil {
		log.Fatalf("Erro para setar o link do cliente %s. \n %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "proxy="+string(Cliente.Link_sip), 110)
	if err != nil {
		log.Fatalf("Erro para setar o link do cliente %s. \n %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "domain="+string(Cliente.Link_sip), 111)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "username="+RamalSelecionado, 112)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "password="+RamalSelecionado+"@abc", 113)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "authID="+RamalSelecionado, 114)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

}
