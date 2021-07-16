package response

type ProdutoDTO struct {
	Data  ListaProdutos `json:"data"`
	Error Error         `json:"error"`
}
