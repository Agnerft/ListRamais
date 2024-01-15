package util

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"runtime"
)

// var UsrCurr *user.User

func ExecuteUnistall(filePath string) error {
	// C:\Users\USER\AppData\Local\MicroSIP\Uninstall.exe

	// filePath := filepath.Join(UserCurrent().HomeDir,
	// 	"AppData",
	// 	"Local",
	// 	"MicroSIP",
	// 	"Uninstall.exe")

	cmd := exec.Command(filePath, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s ", err)
		return err
	}

	fmt.Println("Removido")

	return nil
}

func ExecuteInstall(filePath string) error {
	cmd := exec.Command(filePath, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s ", err)
		return err
	}

	fmt.Println("Instalado")

	return nil
}

func UserCurrent() user.User {
	// Obter o diretório do usuário
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Erro ao obter o diretório do usuário:", err)
	}

	return *usr
}

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		fmt.Printf("Não foi possível detectar o sistema operacional para abrir o navegador automaticamente.")
		return
	}

	if err != nil {
		fmt.Printf("Erro ao abrir o navegador: %v\n", err)
	}
}
