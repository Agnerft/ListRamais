package handler

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/util"
	"github.com/gin-gonic/gin"
)

var (
	Cliente              *domain.Cliente
	url_padrao           = "https://basesip.makesystem.com.br/clientes?documento="
	url                  = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	// destRunningMicroSIP    = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "microsip.exe")
	destDownMicroSIP       = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
	destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")
	ramalAtual             string
	processName            = "microsip.exe"

	ramaisDoCliente []string
)

func HandleHomeClient(c *gin.Context) {

	if c.Request.Method == http.MethodGet {
		// header := c.GetHeader("meu-Header")
		// header := c.Query("cnpj")

		// fmt.Printf("Esse é meu Header -> %s", header)

		fmt.Println("Bem-vindo! Por favor, digite algo:")
		reader := bufio.NewReader(os.Stdin)

		entradaCNPJ, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}

		// if entradaCNPJ ==

		entradaCNPJ = strings.TrimRight(entradaCNPJ, "\n")
		// cnpj := r.FormValue("cnpj")

		cli := &domain.Cliente{}
		cliente, err := cli.RequestJsonCliente(url_padrao + entradaCNPJ)
		if err != nil {
			http.Error(c.Writer, "Erro ao encontrar cliente.", http.StatusBadRequest)
		} else {
			Cliente = &cliente[0]
			c.JSON(http.StatusOK, Cliente)
			fmt.Printf("Encontrei a empresa %s na base. . .\n", Cliente.Cliente)

			// fmt.Println(Cliente.RamaisRegistrados)

			fmt.Println("Os ramais que temos são:")
			ramaisDoCliente = make([]string, 0)

			for _, ramais := range Cliente.RamaisRegistrados {
				fmt.Printf("Ramal %s, ", ramais.Sip)
				ramaisDoCliente = append(ramaisDoCliente, ramais.Sip)

			}

			fmt.Println(ramaisDoCliente)
			fmt.Println("Escolha um ramal, por favor:")
			reader = bufio.NewReader(os.Stdin)

			entradaRAMAL, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Erro ao ler a entrada:", err)
				return
			}
			entradaRAMAL = strings.TrimRight(entradaRAMAL, "\n")

			for _, sip := range ramaisDoCliente {
				if sip == entradaRAMAL {
					fmt.Printf("Achamos o seu ramal %s", sip)
				}
			}
		}

		return
	}

	// Método não suportado
	http.Error(c.Writer, "Método HTTP não suportado", http.StatusMethodNotAllowed)
}

func HandlePostClient(c *gin.Context) {
	// cnpj := c.Request.FormValue("cnpj")

	// cli := domain.Cliente{}
	// cliente, err := cli.RequestJsonCliente(url_padrao + cnpj)
	// if err != nil {
	// 	log.Fatal("Erro ao buscar os ramais.", err)
	// }
	// Cliente = &cliente[0]

	// fmt.Println(c.Request.Body)

}

func HandleSelecionarRamal(c *gin.Context) {
	// Obter o SIP da query string
	RamalSelecionado := c.Query("sip")

	fmt.Printf("SIP Selecionado: %s \n", RamalSelecionado)

	for _, ramal := range Cliente.RamaisRegistrados {

		if RamalSelecionado == ramal.Sip {

			if !ramal.InUse {

				fmt.Printf("ramal %s tem na base e está liberado \n", ramal.Sip)

			}

			fmt.Printf("ramal %s tem na base mas não está liberado \n", ramal.Sip)
		}

	}
	// Responder ao cliente
	c.String(http.StatusOK, "Ramais selecionados %s \n", RamalSelecionado)
	c.String(http.StatusOK, "Começando a instalação . .. \n")

	ramalAtual = RamalSelecionado

	fmt.Println(ramalAtual)

	pid, err := util.GetPIDbyName(processName)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(pid)

	err = util.TaskkillExecute(pid)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Não tem esse processo rodando %s", []any{processName}...)
	}

	// HandleInstallMicrosip(c)

}

func HandleInstallMicrosip(c *gin.Context) {

	if _, err := os.Stat(destDeleleteMicroSIP); err == nil {
		// Se o caminho existe, execute algo

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
			log.Printf("Erro ao instalar o %s", destDownMicroSIP)
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

	} else {
		// Algum erro ocorreu ao verificar o caminho
		fmt.Printf("Erro ao verificar o caminho: %v\n", err)
		// Adicione aqui o código para lidar com o erro, se necessário
	}

	fmt.Printf("Chamando configuração")

	// err := HandleFileConfig(c)
	// if err != nil {
	// 	fmt.Printf("Erro ao editar o Arquivo: %s", err)
	// }

	fmt.Println("teste")
}

func HandleFileConfig(c *gin.Context) {

	fmt.Println("Ta clicando aqui?")

	// fmt.Println(RamalSelecionado)

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

	err = util.ReadFile(destFileConfigMicrosip, "label=", "label="+ramalAtual, 106)
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

	err = util.ReadFile(destFileConfigMicrosip, "username=", "username="+ramalAtual, 110)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "password=", "password="+ramalAtual+"@abc", 111)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

	err = util.ReadFile(destFileConfigMicrosip, "authID=", "authID="+ramalAtual, 112)
	if err != nil {
		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
	}

}

func HandleTeste(c *gin.Context) {
	cmd := exec.Command("ps", "-p", fmt.Sprint(890))
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Deu erro: %s", err)
	}
}
