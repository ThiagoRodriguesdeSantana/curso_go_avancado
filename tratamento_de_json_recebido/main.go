package main

import (
	"curso_go_avancado/tratamento_de_json_recebido/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Erro ao criar um requeste para servidor", err.Error())
	}
	req.SetBasicAuth("teste", "teste")
	resposta, err := cliente.Do(req)

	if err != nil {
		fmt.Println("Erro na pesquisa", err.Error())
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteudo do servidor", err.Error())
		}

		fmt.Println("  ------------------------------------------------------------------ ")

		post := model.BlogPost{}

		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("erro ao fazer o unmarshal", err.Error())
			return
		}

		fmt.Printf("Poste é: %+v\r\n", post)
	}
}
