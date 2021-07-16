package response

type OpcoesParcelamentoDTO struct {
	Data  []OpcoesParcelamento `json:"data"`
	Error Error                `json:"error"`
}

type OpcoesParcelamento struct {
	IdFormaPagamento   int32   `json:"idFormaPagamento"`
	QuantidadeParcelas int32   `json:"quantidadeParcelas"`
	TaxaJurosAoMes     float32 `json:"taxaJurosAoMes"`
	ValorParcela       float32 `json:"valorParcela"`
	ValorTotal         float32 `json:"valorTotal"`
}
