package util

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
)

// var UsrCurr *user.User

func ExecuteUnistall() {
	// C:\Users\USER\AppData\Local\MicroSIP\Uninstall.exe

	filePath := filepath.Join(UserCurrent().HomeDir,
		"AppData",
		"Local",
		"MicroSIP",
		"Uninstall.exe")

	cmd := exec.Command(filePath, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s ", err)

	}

	fmt.Println("Removido")
}

func UserCurrent() user.User {
	// Obter o diretório do usuário
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Erro ao obter o diretório do usuário:", err)
	}

	return *usr
}
