package response

type ChaveDTO struct {
	Data  Chave `json:"data"`
	Error Error `json:"error"`
}

type Chave struct {
	ChavePublica    string     `json:"chavePublica"`
	DataCadastro    CustomTime `json:"dataCadastro"`
	DataExpiracao   CustomTime `json:"dataExpiracao"`
	DataAtualizacao CustomTime `json:"dataAtualizacao"`
	Ativo           bool       `json:"ativo"`
}
