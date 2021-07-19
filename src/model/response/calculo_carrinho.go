package response

import "encoding/json"

func UnmarshalCalculoCarrinho(data []byte) (CalculoCarrinho, error) {
	var r CalculoCarrinho
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CalculoCarrinho) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CalculoCarrinho struct {
	Data  CalculoCarrinhoData `json:"data"`
	Error Error               `json:"error"`
}

type CalculoCarrinhoData struct {
	ValorFrete            float64                  `json:"valorFrete"`
	ValorImpostos         float64                  `json:"valorImpostos"`
	ValorTotaldosProdutos float64                  `json:"valorTotaldosProdutos"`
	ValorTotaldoPedido    float64                  `json:"valorTotaldoPedido"`
	Produtos              []ProdutoCalculoCarrinho `json:"produtos"`
}

type ProdutoCalculoCarrinho struct {
	IDSku              int64   `json:"idSku"`
	PrevisaoEntrega    string  `json:"previsaoEntrega"`
	ValorUnitario      float64 `json:"valorUnitario"`
	ValorTotal         float64 `json:"valorTotal"`
	ValorTotalFrete    float64 `json:"valorTotalFrete"`
	ValorTotalImpostos float64 `json:"valorTotalImpostos"`
	Erro               bool    `json:"erro"`
	MensagemDeErro     string  `json:"mensagemDeErro"`
	CodigoDoErro       string  `json:"codigoDoErro"`
}
