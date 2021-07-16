// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pedidoParceiroDados, err := UnmarshalPedidoParceiroDados(bytes)
//    bytes, err = pedidoParceiroDados.Marshal()

package response

import "encoding/json"

func UnmarshalPedidoParceiroDados(data []byte) (PedidoParceiroDados, error) {
	var r PedidoParceiroDados
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PedidoParceiroDados) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PedidoParceiroDados struct {
	Data  Data  `json:"data"`
	Error Error `json:"error"`
}

type Data struct {
	Pedido       Pedido       `json:"pedido"`
	Endereco     Endereco     `json:"endereco"`
	Destinatario Destinatario `json:"destinatario"`
	Entregas     []Entregas   `json:"entregas"`
}

type Destinatario struct {
	Nome              string `json:"nome"`
	CpfCnpj           string `json:"cpfCnpj"`
	InscricaoEstadual string `json:"inscricaoEstadual"`
	Email             string `json:"email"`
}

type Endereco struct {
	Cep         string `json:"cep"`
	Estado      string `json:"estado"`
	Logradouro  string `json:"logradouro"`
	Cidade      string `json:"cidade"`
	Numero      int64  `json:"numero"`
	Referencia  string `json:"referencia"`
	Bairro      string `json:"bairro"`
	Complemento string `json:"complemento"`
	Telefone    string `json:"telefone"`
	Telefone2   string `json:"telefone2"`
	Telefone3   string `json:"telefone3"`
}

type Entregas struct {
	CodigoEntrega         float64           `json:"codigoEntrega"`
	PrevisaoEntrega       string            `json:"previsaoEntrega"`
	DataEntrega           string            `json:"dataEntrega"`
	DataPrevisao          string            `json:"dataPrevisao"`
	DataEmissaoNotaFiscal string            `json:"dataEmissaoNotaFiscal"`
	IDNotaFiscal          int64             `json:"idNotaFiscal"`
	SerieNotaFiscal       string            `json:"serieNotaFiscal"`
	ChaveAcesso           string            `json:"chaveAcesso"`
	TrackingEntrega       []TrackingEntrega `json:"trackingEntrega"`
	ProdutoEntrega        []ProdutoEntrega  `json:"produtoEntrega"`
	RastreioEntrega       string            `json:"rastreioEntrega"`
	NomeTransportadora    string            `json:"nomeTransportadora"`
	LinkNotaFiscalPDF     string            `json:"linkNotaFiscalPDF"`
	ListNotaFiscalXML     string            `json:"listNotaFiscalXML"`
	Estorno               bool              `json:"estorno"`
	Origem                string            `json:"origem"`
	Motivo                Motivo            `json:"motivo"`
}

type Motivo struct {
	Categoria  string `json:"categoria"`
	Assunto    string `json:"assunto"`
	Motivo     string `json:"motivo"`
	Observacao string `json:"observacao"`
}

type ProdutoEntrega struct {
	Codigo         int64   `json:"codigo"`
	Nome           string  `json:"nome"`
	Quantidade     int64   `json:"quantidade"`
	Valor          float64 `json:"valor"`
	Frete          float64 `json:"frete"`
	ValorAdicional float64 `json:"valorAdicional"`
	ValorTotal     float64 `json:"valorTotal"`
	IDLojista      float64 `json:"idLojista"`
}

type TrackingEntrega struct {
	CodDescricao string `json:"codDescricao"`
	Data         string `json:"data"`
	DataEntrega  string `json:"dataEntrega"`
	Descricao    string `json:"descricao"`
}

type Pedido struct {
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
	Premio     int64   `json:"premio"`
	PrecoVenda float64 `json:"precoVenda"`
}
