package response

type ProdutosDTO struct {
	Data  []ListaProdutos `json:"data"`
	Error Error           `json:"error"`
}
