package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Ramal struct {
	Sip   string `json:"sip"`
	Ip    string `json:"ip"`
	InUse bool   `json:"inuse"`
	// Empresa string `json:"empresa"`
}

type RamaisRegistrados struct {
	RamaisRegistrados []Ramal `json:"ramais_registrados"`
}

func (r *Ramal) RequestJsonRamal(url string) (*RamaisRegistrados, error) {
	// Fazer uma requisição HTTP para obter os dados JSON
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição HTTP:", err)
		return nil, err
	}

	defer response.Body.Close()

	// fmt.Println(response.Body)

	var ramais RamaisRegistrados

	err = json.NewDecoder(response.Body).Decode(&ramais)
	if err != nil {
		log.Fatal("Erro ao decodificar o JSON:", err)
		return nil, err
	}
	count := 0
	// Imprimir os dados dos ramais registrados
	for _, ramal := range ramais.RamaisRegistrados {
		// fmt.Printf("SIP: %s, IP: %s\n", ramal.Sip, ramal.Ip)
		if ramal.Sip != "" {
			count++
		} else {
			fmt.Println("Vazio")
		}

	}

	fmt.Println(count)

	return &ramais, nil
}
