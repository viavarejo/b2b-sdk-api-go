// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pedidoCarrinho, err := UnmarshalPedidoCarrinho(bytes)
//    bytes, err = pedidoCarrinho.Marshal()

package main

import "encoding/json"

func UnmarshalPedidoCarrinho(data []byte) (PedidoCarrinho, error) {
	var r PedidoCarrinho
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PedidoCarrinho) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PedidoCarrinho struct {
	IDCampanha           int64     `json:"idCampanha"`
	Cnpj                 string    `json:"cnpj"`
	Cep                  string    `json:"cep"`
	IDEntregaTipo        int64     `json:"idEntregaTipo"`
	IDEnderecoLojaFisica int64     `json:"idEnderecoLojaFisica"`
	IDUnidadeNegocio     int64     `json:"idUnidadeNegocio"`
	Produtos             []Produto `json:"produtos"`
}

type Produto struct {
	Codigo     int64 `json:"codigo"`
	Quantidade int64 `json:"quantidade"`
	IDLojista  int64 `json:"idLojista"`
}
