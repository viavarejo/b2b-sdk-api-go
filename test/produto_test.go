package test

import (
	"api"
	"testing"
)

func TestGetDadosProdutosSucess(t *testing.T) {
	dto := api.GetDadosProduto("15", "5880205")

	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Data.Nome != "Bola de Natal Santini Christmas 10cm Transparente - 3 Unidades." {
		t.Error("Test failed-2")
	}
	if dto.Data.Imagem != "http://imagens.extra.com.br/Control/ArquivoExibir.aspx?IdArquivo=253172122" {
		t.Error("Test failed-3")
	}
	if dto.Data.Categoria != 2868 {
		t.Error("Test failed-4")
	}
	if dto.Data.Valor != 29.9 {
		t.Error("Test failed-5")
	}
}

func TestGetListaDadosProdutosSucess(t *testing.T) {
	dto := api.GetListaDadosProdutos("15", []string{"5880205", "5880206"})
	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Data[0].Nome != "Bola de Natal Santini Christmas 10cm Transparente - 3 Unidades." {
		t.Error("Test failed-2")
	}
	if dto.Data[0].Descricao != "Sua decora&amp#231&amp#227o de Natal vai ficar moderna e cheia de estilo com o Conjunto de Bolas Santini Christmas - Branco/Transparente. Elas contam com 10 cent&amp#237metros de di&amp#226metro e s&amp#227o feitas em material pl&amp#225stico resistente com detalhes em glitter branco.&ampltbr&ampgt&amp#13&amp#10Um conjunto para decorar &amp#225rvores de Natal, ambientes e outros arranjos. Com os produtos Santini Christmas sua decora&amp#231&amp#227o de Natal vai ganhar elogios dos convidados!" {
		t.Error("Test failed-3")
	}
	if dto.Data[0].Imagem != "http://imagens.extra.com.br/Control/ArquivoExibir.aspx?IdArquivo=253172122" {
		t.Error("Test failed-4")
	}
	if dto.Data[0].Categoria != 2868 {
		t.Error("Test failed-5")
	}
	if dto.Data[0].Valor != 29.9 {
		t.Error("Test failed-5")
	}

}

func TestGetDadosProdutoCampanhaSucess(t *testing.T) {
	dto := api.GetDadosProdutoCampanha("5940", "5880205", "57.822.975/0001-12", "15")

	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Data.Nome != "Bola de Natal Santini Christmas 10cm Transparente - 3 Unidades." {
		t.Error("Test failed-2")
	}
	if dto.Data.Imagem != "http://imagens.extra.com.br/Control/ArquivoExibir.aspx?IdArquivo=253172122" {
		t.Error("Test failed-3")
	}
	if dto.Data.Categoria != 2868 {
		t.Error("Test failed-4")
	}
	if dto.Data.Valor != 29.9 {
		t.Error("Test failed-5")
	}
	if dto.Data.ValorDe != 29.9 {
		t.Error("Test failed-6")
	}
}

func TestGetDadosProdutosFail(t *testing.T) {
	dto := api.GetDadosProduto("15", "595959")

	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Error.Code != "ProdutoNaoEncontrado" {
		t.Error("Test failed-2")
	}
	if dto.Error.Message != "Produto ou Sku 595959 não encontrado" {
		t.Error("Test failed-3")
	}
}

func TestGetListaDadosProdutosFail(t *testing.T) {
	ids := []string{"595959"}
	dto := api.GetListaDadosProdutos("15", ids)
	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Error.Message != "Nenhum SKU foi encontrado.(ErroValidacao)" {
		t.Error("Test failed-2")
	}
	if dto.Error.Code != "NaoEncontrado" {
		t.Error("Test failed-3")
	}
}

func TestGetDadosProdutoCampanhaFail(t *testing.T) {
	dto := api.GetDadosProdutoCampanha("5940", "595959", "2", "2")
	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Error.Message != "A campanha não pertence ao cliente informado.(ErroValidacao)" {
		t.Error("Test failed-2")
	}
	if dto.Error.Code != "400" {
		t.Error("Test failed-3")
	}
}
