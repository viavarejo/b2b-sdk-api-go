package response

import "time"

type ChaveDTO struct {
	Data  Chave `json:"chave"`
	Error Error `json:"error"`
}

type Chave struct {
	ChavePublica    string    `json:"chavePublica"`
	DataCadastro    time.Time `json:"dataCadastro"`
	DataExpiracao   time.Time `json:"dataExpiracao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
	Ativo           bool      `json:"ativo"`
}
