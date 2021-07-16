package response

type ListaProdutos struct {
	Nome            string                `json:"nome"`
	Descricao       string                `json:"descricao"`
	Imagem          string                `json:"imagem"`
	Categoria       int32                 `json:"categoria"`
	Valor           float32               `json:"valor"`
	ValorDe         float32               `json:"valorDe"`
	Disponibilidade bool                  `json:"disponibilidade"`
	ForaDeLinha     bool                  `json:"foraDeLinha"`
	IdLojista       int32                 `json:"idLojista"`
	DadosEntrega    []ProdutoDadosEntrega `json:"dadosEntrega"`
}

type ProdutoDadosEntrega struct {
	IdEntregaTipo   int32   `json:"idEntregaTipo"`
	Disponibilidade bool    `json:"disponibilidade"`
	Fretes          []Frete `json:"fretes"`
}

type Frete struct {
	Estado        string `json:"estado"`
	PrecoCapital  string `json:"precoCapital"`
	PrecoInterior string `json:"precoInterior"`
}
