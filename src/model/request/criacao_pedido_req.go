package request

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
	Produtos                       []ProdutoCriacaoPedido      `json:"produtos,omitempty"`
	EnderecoEntrega                Endereco                    `json:"enderecoEntrega,omitempty"`
	Destinatario                   Destinatario                `json:"destinatario,omitempty"`
	Campanha                       int64                       `json:"campanha,omitempty"`
	Cnpj                           string                      `json:"cnpj,omitempty"`
	PedidoParceiro                 int64                       `json:"pedidoParceiro,omitempty"`
	IDPedidoMktplc                 string                      `json:"idPedidoMktplc,omitempty"`
	SenhaAtendimento               string                      `json:"senhaAtendimento,omitempty"`
	Apolice                        string                      `json:"apolice,omitempty"`
	Administrador                  int64                       `json:"administrador,omitempty"`
	ParametrosExtras               string                      `json:"parametrosExtras,omitempty"`
	ValorFrete                     float64                     `json:"valorFrete,omitempty"`
	AguardarConfirmacao            bool                        `json:"aguardarConfirmacao,omitempty"`
	OptantePeloSimples             bool                        `json:"optantePeloSimples,omitempty"`
	PossuiPagtoComplementar        bool                        `json:"possuiPagtoComplementar,omitempty"`
	PagtosComplementares           []PagtosComplementare       `json:"pagtosComplementares,omitempty"`
	DadosEntrega                   DadosEntrega                `json:"dadosEntrega,omitempty"`
	EnderecoCobranca               Endereco                    `json:"enderecoCobranca,omitempty"`
	ValorTotalPedido               float64                     `json:"valorTotalPedido,omitempty"`
	ValorTotalComplementar         float64                     `json:"valorTotalComplementar,omitempty"`
	ValorTotalComplementarCOMJuros float64                     `json:"valorTotalComplementarComJuros,omitempty"`
	IntermediadoresFinanceiros     []IntermediadoresFinanceiro `json:"intermediadoresFinanceiros,omitempty"`
}

type DadosEntrega struct {
	PrevisaoDeEntrega    string  `json:"previsaoDeEntrega,omitempty"`
	ValorFrete           float64 `json:"valorFrete,omitempty"`
	IDEntregaTipo        int64   `json:"idEntregaTipo,omitempty"`
	IDEnderecoLojaFisica int64   `json:"idEnderecoLojaFisica,omitempty"`
	IDUnidadeNegocio     int64   `json:"idUnidadeNegocio,omitempty"`
}

type Destinatario struct {
	Nome              string `json:"nome,omitempty"`
	CpfCnpj           string `json:"cpfCnpj,omitempty"`
	InscricaoEstadual string `json:"inscricaoEstadual,omitempty"`
	Email             string `json:"email,omitempty"`
}

type Endereco struct {
	Cep         string `json:"cep,omitempty"`
	Estado      string `json:"estado,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	Numero      int64  `json:"numero,omitempty"`
	Referencia  string `json:"referencia,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Telefone    string `json:"telefone,omitempty"`
	Telefone2   string `json:"telefone2,omitempty"`
	Telefone3   string `json:"telefone3,omitempty"`
}

type IntermediadoresFinanceiro struct {
	FormaPagamento           int64   `json:"formaPagamento,omitempty"`
	MeioPagamento            string  `json:"meioPagamento,omitempty"`
	ValorPagamento           float64 `json:"valorPagamento,omitempty"`
	TipoIntegracaoPagamento  int64   `json:"tipoIntegracaoPagamento,omitempty"`
	CnpjIntermediador        string  `json:"cnpjIntermediador,omitempty"`
	RazaoSocialIntermediador string  `json:"razaoSocialIntermediador,omitempty"`
	BandeiraOperadoraCartao  string  `json:"bandeiraOperadoraCartao,omitempty"`
	NumAutorizacaoCartao     string  `json:"numAutorizacaoCartao,omitempty"`
}

type PagtosComplementare struct {
	IDFormaPagamento            int64                       `json:"idFormaPagamento,omitempty"`
	DadosCartaoCredito          DadosCartaoCredito          `json:"dadosCartaoCredito,omitempty"`
	DadosCartaoCreditoValidacao DadosCartaoCreditoValidacao `json:"dadosCartaoCreditoValidacao,omitempty"`
	ValorComplementarCOMJuros   float64                     `json:"valorComplementarComJuros,omitempty"`
	ValorComplementar           float64                     `json:"valorComplementar,omitempty"`
}

type DadosCartaoCredito struct {
	Nome               string `json:"nome,omitempty"`
	Numero             string `json:"numero,omitempty"`
	CodigoVerificador  string `json:"codigoVerificador,omitempty"`
	ValidadeAno        string `json:"validadeAno,omitempty"`
	ValidadeMes        string `json:"validadeMes,omitempty"`
	ValidadeAnoMes     string `json:"validadeAnoMes,omitempty"`
	QuantidadeParcelas int64  `json:"quantidadeParcelas,omitempty"`
}

type DadosCartaoCreditoValidacao struct {
	Nome                          string `json:"nome,omitempty"`
	NumeroMascarado               string `json:"numeroMascarado,omitempty"`
	QtCaracteresCodigoVerificador string `json:"qtCaracteresCodigoVerificador,omitempty"`
	ValidadeAno                   string `json:"validadeAno,omitempty"`
	ValidadeMes                   string `json:"validadeMes,omitempty"`
}

type ProdutoCriacaoPedido struct {
	IDLojista  int64   `json:"idLojista,omitempty"`
	Codigo     int64   `json:"codigo,omitempty"`
	Quantidade int64   `json:"quantidade,omitempty"`
	Premio     float64 `json:"premio,omitempty"`
	PrecoVenda float64 `json:"precoVenda,omitempty"`
}
