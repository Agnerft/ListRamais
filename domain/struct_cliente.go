package domain

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/makesystem/list_ramais/domain"
)

type Cliente struct {
	Cliente           string `json:"cliente"`
	Documento         string `json:"documento"`
	Quantidade_ramal  int    `json:"quantidade_ramal"`
	Link              string `json:"link"`
	Id                int    `json:"id"`
	Link_sip          string `json:"link_sip"`
	RamaisRegistrados []domain.Ramal
}

type ClientesRegistrados struct {
	ClientesRegistrados []Cliente
}

// type ClientesRegistrados []Cliente

func (c *Cliente) RequestJsonCliente(url string) ([]Cliente, error) {
	// Fazer uma requisição HTTP para obter os dados JSON
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição HTTP:", err)
		return nil, err
	}
	defer response.Body.Close()

	// Imprimir o conteúdo do corpo da resposta
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Erro ao ler o corpo da resposta:", err)
		return nil, err
	}

	// fmt.Println(string(responseBody))

	var clientes []Cliente

	// err = json.NewDecoder(response.Body).Decode(&cliente)
	// if err != nil {
	// 	log.Fatal("Erro ao decodificar o JSON:", err)
	// 	return nil, err
	// }

	err = json.Unmarshal(responseBody, &clientes)
	if err != nil {
		log.Fatal("Erro ao decodificar o JSON:", err)
		return nil, err
	}

	for i, cliente := range clientes {
		// fmt.Println(cliente.Link)

		ramais, err := domain.RequestJsonRamal(cliente.Link + "/status_central")
		if err != nil {
			log.Fatalln("Usuário não encontrado.", err)
			return nil, err
		} else {
			// Atualiza a lista de ramais no cliente
			clientes[i].RamaisRegistrados = ramais.RamaisRegistrados
		}

	}

	return clientes, nil
}
