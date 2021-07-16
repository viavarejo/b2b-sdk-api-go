// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    criacaoPedidoResp, err := UnmarshalCriacaoPedidoResp(bytes)
//    bytes, err = criacaoPedidoResp.Marshal()

package response

import "encoding/json"

func UnmarshalCriacaoPedidoResp(data []byte) (CriacaoPedidoResp, error) {
	var r CriacaoPedidoResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CriacaoPedidoResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CriacaoPedidoResp struct {
	Data  Data  `json:"data"`
	Error Error `json:"error"`
}

type Data struct {
	ValorProduto               float64                    `json:"valorProduto"`
	ValorTotalPedido           float64                    `json:"valorTotalPedido"`
	CodigoPedido               int64                      `json:"codigoPedido"`
	PedidoParceiro             int64                      `json:"pedidoParceiro"`
	IDPedidoMktplc             string                     `json:"idPedidoMktplc"`
	Produtos                   []Produto                  `json:"produtos"`
	ParametrosExtras           string                     `json:"parametrosExtras"`
	AguardandoConfirmacao      bool                       `json:"aguardandoConfirmacao"`
	DadosEntrega               DadosEntrega               `json:"dadosEntrega"`
	DadosPagamentoComplementar DadosPagamentoComplementar `json:"dadosPagamentoComplementar"`
}

type DadosEntrega struct {
	PrevisaoDeEntrega    string  `json:"previsaoDeEntrega"`
	ValorFrete           float64 `json:"valorFrete"`
	IDEntregaTipo        int64   `json:"idEntregaTipo"`
	IDEnderecoLojaFisica int64   `json:"idEnderecoLojaFisica"`
	IDUnidadeNegocio     int64   `json:"idUnidadeNegocio"`
}

type DadosPagamentoComplementar struct {
	Pagamentos                     []Pagamento `json:"pagamentos"`
	ValorTotalComplementar         float64     `json:"valorTotalComplementar"`
	ValorTotalComplementarCOMJuros float64     `json:"valorTotalComplementarComJuros"`
}

type Pagamento struct {
	CodigoDoErro       string  `json:"codigoDoErro"`
	ValorComplementar  float64 `json:"valorComplementar"`
	QuantidadeParcelas int64   `json:"quantidadeParcelas"`
	ValorParcela       float64 `json:"valorParcela"`
	IDFormaPagamento   int64   `json:"idFormaPagamento"`
	Erro               bool    `json:"erro"`
	MensagemDeErro     string  `json:"mensagemDeErro"`
	URL                string  `json:"url"`
}

type Produto struct {
	IDLojista  int64   `json:"idLojista"`
	Codigo     int64   `json:"codigo"`
	Quantidade int64   `json:"quantidade"`
	Premio     float64 `json:"premio"`
	PrecoVenda float64 `json:"precoVenda"`
}
