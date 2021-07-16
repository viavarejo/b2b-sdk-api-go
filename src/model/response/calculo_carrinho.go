// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    calculoCarrinho, err := UnmarshalCalculoCarrinho(bytes)
//    bytes, err = calculoCarrinho.Marshal()

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
	Data  Data  `json:"data"`
	Error Error `json:"error"`
}

type Data struct {
	ValorFrete            int64     `json:"valorFrete"`
	ValorImpostos         int64     `json:"valorImpostos"`
	ValorTotaldosProdutos int64     `json:"valorTotaldosProdutos"`
	ValorTotaldoPedido    int64     `json:"valorTotaldoPedido"`
	Produtos              []Produto `json:"produtos"`
}

type Produto struct {
	IDSku              int64  `json:"idSku"`
	PrevisaoEntrega    string `json:"previsaoEntrega"`
	ValorUnitario      int64  `json:"valorUnitario"`
	ValorTotal         int64  `json:"valorTotal"`
	ValorTotalFrete    int64  `json:"valorTotalFrete"`
	ValorTotalImpostos int64  `json:"valorTotalImpostos"`
	Erro               bool   `json:"erro"`
	MensagemDeErro     string `json:"mensagemDeErro"`
	CodigoDoErro       string `json:"codigoDoErro"`
}
