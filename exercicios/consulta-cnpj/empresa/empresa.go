package empresa

import "strings"

type AtividadePrincipal struct {
	Text string `json:"text"`
}

type Empresa struct {
	Nome       string               `json:"nome"`
	Situacao   string               `json:"situacao"`
	Atividades []AtividadePrincipal `json:"atividade_principal"`
}

func LimparCNPJ(cnpj string) string {
	// Remover formatação do CNPJ
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	return cnpj
}
