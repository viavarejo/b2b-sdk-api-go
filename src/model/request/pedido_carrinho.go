package request

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
	IDCampanha           int64                   `json:"idCampanha,omitempty"`
	Cnpj                 string                  `json:"cnpj,omitempty"`
	Cep                  string                  `json:"cep,omitempty"`
	IDEntregaTipo        int64                   `json:"idEntregaTipo,omitempty"`
	IDEnderecoLojaFisica int64                   `json:"idEnderecoLojaFisica,omitempty"`
	IDUnidadeNegocio     int64                   `json:"idUnidadeNegocio,omitempty"`
	Produtos             []ProdutoPedidoCarrinho `json:"produtos,omitempty"`
}

type ProdutoPedidoCarrinho struct {
	Codigo     int64 `json:"codigo,omitempty"`
	Quantidade int64 `json:"quantidade,omitempty"`
	IDLojista  int64 `json:"idLojista,omitempty"`
}
