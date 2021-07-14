// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    criacaoPedidoReq, err := UnmarshalCriacaoPedidoReq(bytes)
//    bytes, err = criacaoPedidoReq.Marshal()

package main

import "encoding/json"

func UnmarshalCriacaoPedidoReq(data []byte) (CriacaoPedidoReq, error) {
	var r CriacaoPedidoReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CriacaoPedidoReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CriacaoPedidoReq struct {
	Produtos                       []Produto                   `json:"produtos"`
	EnderecoEntrega                Endereco                    `json:"enderecoEntrega"`
	Destinatario                   Destinatario                `json:"destinatario"`
	Campanha                       int64                       `json:"campanha"`
	Cnpj                           string                      `json:"cnpj"`
	PedidoParceiro                 int64                       `json:"pedidoParceiro"`
	IDPedidoMktplc                 string                      `json:"idPedidoMktplc"`
	SenhaAtendimento               string                      `json:"senhaAtendimento"`
	Apolice                        string                      `json:"apolice"`
	Administrador                  int64                       `json:"administrador"`
	ParametrosExtras               string                      `json:"parametrosExtras"`
	ValorFrete                     float64                     `json:"valorFrete"`
	AguardarConfirmacao            bool                        `json:"aguardarConfirmacao"`
	OptantePeloSimples             bool                        `json:"optantePeloSimples"`
	PossuiPagtoComplementar        bool                        `json:"possuiPagtoComplementar"`
	PagtosComplementares           []PagtosComplementare       `json:"pagtosComplementares"`
	DadosEntrega                   DadosEntrega                `json:"dadosEntrega"`
	EnderecoCobranca               Endereco                    `json:"enderecoCobranca"`
	ValorTotalPedido               float64                     `json:"valorTotalPedido"`
	ValorTotalComplementar         float64                     `json:"valorTotalComplementar"`
	ValorTotalComplementarCOMJuros float64                     `json:"valorTotalComplementarComJuros"`
	IntermediadoresFinanceiros     []IntermediadoresFinanceiro `json:"intermediadoresFinanceiros"`
}

type DadosEntrega struct {
	PrevisaoDeEntrega    string  `json:"previsaoDeEntrega"`
	ValorFrete           float64 `json:"valorFrete"`
	IDEntregaTipo        int64   `json:"idEntregaTipo"`
	IDEnderecoLojaFisica int64   `json:"idEnderecoLojaFisica"`
	IDUnidadeNegocio     int64   `json:"idUnidadeNegocio"`
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

type IntermediadoresFinanceiro struct {
	FormaPagamento           int64   `json:"formaPagamento"`
	MeioPagamento            string  `json:"meioPagamento"`
	ValorPagamento           float64 `json:"valorPagamento"`
	TipoIntegracaoPagamento  int64   `json:"tipoIntegracaoPagamento"`
	CnpjIntermediador        string  `json:"cnpjIntermediador"`
	RazaoSocialIntermediador string  `json:"razaoSocialIntermediador"`
	BandeiraOperadoraCartao  string  `json:"bandeiraOperadoraCartao"`
	NumAutorizacaoCartao     string  `json:"numAutorizacaoCartao"`
}

type PagtosComplementare struct {
	IDFormaPagamento            int64                       `json:"idFormaPagamento"`
	DadosCartaoCredito          DadosCartaoCredito          `json:"dadosCartaoCredito"`
	DadosCartaoCreditoValidacao DadosCartaoCreditoValidacao `json:"dadosCartaoCreditoValidacao"`
	ValorComplementarCOMJuros   float64                     `json:"valorComplementarComJuros"`
	ValorComplementar           float64                     `json:"valorComplementar"`
}

type DadosCartaoCredito struct {
	Nome               string `json:"nome"`
	Numero             string `json:"numero"`
	CodigoVerificador  string `json:"codigoVerificador"`
	ValidadeAno        string `json:"validadeAno"`
	ValidadeMes        string `json:"validadeMes"`
	ValidadeAnoMes     string `json:"validadeAnoMes"`
	QuantidadeParcelas int64  `json:"quantidadeParcelas"`
}

type DadosCartaoCreditoValidacao struct {
	Nome                          string `json:"nome"`
	NumeroMascarado               string `json:"numeroMascarado"`
	QtCaracteresCodigoVerificador string `json:"qtCaracteresCodigoVerificador"`
	ValidadeAno                   string `json:"validadeAno"`
	ValidadeMes                   string `json:"validadeMes"`
}

type Produto struct {
	IDLojista  int64   `json:"idLojista"`
	Codigo     int64   `json:"codigo"`
	Quantidade int64   `json:"quantidade"`
	Premio     float64 `json:"premio"`
	PrecoVenda float64 `json:"precoVenda"`
}
