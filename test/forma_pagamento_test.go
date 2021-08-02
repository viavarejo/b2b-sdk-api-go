package test

import (
	"testing"

	"github.com/viavarejo/b2b-sdk-api-go/src/api"
	"github.com/viavarejo/b2b-sdk-api-go/src/model/response"
)

func TestFormaPagamento(t *testing.T) {
	t.Run("Deveria retornar as Opcoes de Parcelamento",
		func(t *testing.T) {
			const VALOR_PARCELA_EXPECTED = 1000.0
			var dto response.OpcoesParcelamentoDTO = api.GetOpcoesParcelamento("1", "5940", "57.822.975/0001-12", "1000")
			if len(dto.Data) == 0 {
				t.Error("Test failed-1")
				t.Fail()
			}
			if dto.Data[0].ValorParcela != VALOR_PARCELA_EXPECTED {
				t.Error("Test failed-2")
				t.Fail()
			}
		})

	t.Run("Deveria falhar o retorno das Opcoes de Parcelamento",
		func(t *testing.T) {
			var dto response.OpcoesParcelamentoDTO = api.GetOpcoesParcelamento("8", "5940", "57.822.975/0001-12", "1000")
			if len(dto.Data) > 0 {
				t.Error("Test failed-1")
			}
		})
}
