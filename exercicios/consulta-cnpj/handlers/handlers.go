package handlers

import (
	"consulta-cnpj/empresa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const receitaAPIURL = "https://www.receitaws.com.br/v1/cnpj/"

func ConsultarCNPJHandler(w http.ResponseWriter, r *http.Request) {
	cnpj := r.URL.Query().Get("cnpj")
	if cnpj == "" {
		http.Error(w, "O cnpj é obrigatório.", http.StatusBadRequest)
		return
	}

	cnpjLimpo := empresa.LimparCNPJ(cnpj)
	empresaInfo, err := consultarCNPJ(cnpjLimpo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Dados da empresa
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empresaInfo)
}

func consultarCNPJ(cnpj string) (*empresa.Empresa, error) {

	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	// Solicitação HTTP para a API da Receita Federal
	resp, err := http.Get(receitaAPIURL + cnpj)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("não foi possível obter informações da Receita Federal. Status code: %d", resp.StatusCode)
	}

	// Leia o corpo da resposta JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Analise a resposta JSON
	var empresa empresa.Empresa
	err = json.Unmarshal(body, &empresa)
	if err != nil {
		return nil, err
	}

	return &empresa, nil
}
