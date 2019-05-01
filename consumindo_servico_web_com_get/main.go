package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("Erro na pesquisa", err.Error())
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteudo", err.Error())
		}
		fmt.Println(string(corpo))
	}

	//Para chamada que precise de altenticação usar o codigo seguinte

	req, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("Erro ao criar um requeste", err.Error())
	}
	req.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(req)

	if err != nil {
		fmt.Println("Erro na pesquisa", err.Error())
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteudo", err.Error())
		}
		fmt.Println("  ------------------------------------------------------------------ ")
		fmt.Println(string(corpo))
	}

}
