package test

import (
	"api"
	"testing"
)

func TestGetOpcoesParcelamentoSucess(t *testing.T) {
	const VALOR_PARCELA_EXPECTED = 1000.0
	dto := api.GetOpcoesParcelamento("1", "5940", "57.822.975/0001-12", "1000")
	if &dto.Data == nil && dto.Data[0].ValorParcela != VALOR_PARCELA_EXPECTED {
		t.Error("Test failed")
	}
}

//ERRO FORA DO PADRÃO(objeto vem todo vazio)
func TestGetOpcoesParcelamentoFail(t *testing.T) {
	dto := api.GetOpcoesParcelamento("8", "5940", "57.822.975/0001-12", "1000")
	if &dto.Data == nil && len(dto.Data) > 0 {
		t.Error("Test failed")
	}
}
