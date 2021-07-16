package response

type FormasPagamentoDTO struct {
	Data  []FormasPagamento `json:"data"`
	Error Error             `json:"error"`
}

type FormasPagamento struct {
	IdFormaPagamento int32  `json:"idFormaPagamento"`
	Nome             string `json:"nome"`
}
