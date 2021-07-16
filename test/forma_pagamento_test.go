package test

import (
	"api"
	"testing"
)

func TestGetOpcoesParcelamentoSucess(t *testing.T) {
	const VALOR_PARCELA_EXPECTED = 1000.0
	dto := api.GetOpcoesParcelamento("1", "5940", "57.822.975/0001-12", "1000")
	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if dto.Data[0].ValorParcela != VALOR_PARCELA_EXPECTED {
		t.Error("Test failed-2")
	}
}

//ERRO FORA DO PADRÃO(objeto vem todo vazio)
func TestGetOpcoesParcelamentoFail(t *testing.T) {
	dto := api.GetOpcoesParcelamento("8", "5940", "57.822.975/0001-12", "1000")
	if &dto.Data == nil {
		t.Error("Test failed-1")
	}
	if len(dto.Data) > 0 {
		t.Error("Test failed-2")
	}
}
