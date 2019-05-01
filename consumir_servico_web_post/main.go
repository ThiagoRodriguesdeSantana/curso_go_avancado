package main

import (
	"bytes"
	"curso_go_avancado/consumir_servico_web_post/model"
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

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Thiago"

	conteudoParaEnviar, err := json.Marshal(usuario)

	if err != nil {
		fmt.Println("erro ao gerar json", err.Error())
	}
	req, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/1e1q8w91", bytes.NewBuffer(conteudoParaEnviar))
	if err != nil {
		fmt.Println("Erro ao chamar servico requeste", err.Error())
	}
	req.SetBasicAuth("fizz", "buzz")
	req.Header.Set("content-type", "application/json; charset=utf-8")

	resposta, err := cliente.Do(req)

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
