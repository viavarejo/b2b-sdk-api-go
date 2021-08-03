package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/viavarejo/b2b-sdk-api-go/src/api"

	"github.com/viavarejo/b2b-sdk-api-go/src/model/response"
)

func TestProdutos(t *testing.T) {
	t.Run("Deveria retornar os Produtos",
		func(t *testing.T) {
			var dto response.ProdutoDTO = api.GetDadosProduto("15", "5880205")
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			/*if &dto.Data == nil {
				t.Error("Test failed-1")
			}*/
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
		})

	t.Run("Deveria retornar uma lista de dados dos Produtos",
		func(t *testing.T) {
			var dto response.ProdutosDTO = api.GetListaDadosProdutos("15", []string{"5880205", "5880206"})
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			/*if &dto.Data == nil {
				t.Error("Test failed-1")
			}*/
			if dto.Data[0].Nome != "Bola de Natal Santini Christmas 8cm Vermelha/Dourada - 6 Unidades." {
				t.Error("Test failed-2")
			}
			if dto.Data[0].Descricao != "As cores vermelha e dourada s&amp;#227;o as que mais representam o Natal e fazem parte da decora&amp;#231;&amp;#227;o deste conjunto de bolas Santini Christmas. Composto por seis pe&amp;#231;as desenvolvidas em material pl&amp;#225;stico e com acabamento em brilho para agregar ainda mais sofistica&amp;#231;&amp;#227;o ao ambiente.&amp;lt;br&amp;gt;&amp;#13;&amp;#10;Ideal para adornar arranjos, &amp;#225;rvores de Natal e at&amp;#233; mesmo para decorar ambientes maiores, como sal&amp;#245;es e escrit&amp;#243;rios." {
				t.Error("Test failed-3")
			}
			if dto.Data[0].Imagem != "http://imagens.extra.com.br/Control/ArquivoExibir.aspx?IdArquivo=253172225" {
				t.Error("Test failed-4")
			}
			if dto.Data[0].Categoria != 2867 {
				t.Error("Test failed-5")
			}
			if dto.Data[0].Valor != 22.9 {
				t.Error("Test failed-5")
			}

		})

	t.Run("Deveria retornar os dados do Produto da Campanha",
		func(t *testing.T) {
			var dto response.ProdutoDTO = api.GetDadosProdutoCampanha("5940", "5880205", "57.822.975/0001-12", "15")
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			/*if &dto.Data == nil {
				t.Error("Test failed-1")
			}*/
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
		})
	t.Run("Falhar ao buscar dados do Produto",
		func(t *testing.T) {
			var dto response.ProdutoDTO = api.GetDadosProduto("15", "595959")
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			/*if &dto.Data == nil {
				t.Error("Test failed-1")
			}*/
			if dto.Error.Code != "ProdutoNaoEncontrado" {
				t.Error("Test failed-2")
			}
			if dto.Error.Message != "Produto ou Sku 595959 não encontrado" {
				t.Error("Test failed-3")
			}
		})

	t.Run("Falhar ao buscar lista de dados dos Produtos",
		func(t *testing.T) {
			ids := []string{"595959"}
			var dto response.ProdutosDTO = api.GetListaDadosProdutos("15", ids)
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			/*if &dto.Data == nil {
				t.Error("Test failed-1")
			}*/
			if dto.Error.Message != "Nenhum SKU foi encontrado.(ErroValidacao)" {
				t.Error("Test failed-2")
			}
			if dto.Error.Code != "NaoEncontrado" {
				t.Error("Test failed-3")
			}
		})

	t.Run("Falhar ao buscar dados do Produto da Campanha",
		func(t *testing.T) {
			var dto response.ProdutoDTO = api.GetDadosProdutoCampanha("5940", "595959", "2", "2")
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			/*if &dto.Data == nil {
				t.Error("Test failed-1")
			}*/
			if dto.Error.Message != "A campanha não pertence ao cliente informado.(ErroValidacao)" {
				t.Error("Test failed-2")
			}
			if dto.Error.Code != "400" {
				t.Error("Test failed-3")
			}
		})
}
