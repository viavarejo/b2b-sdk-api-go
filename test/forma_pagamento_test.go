package main

import (
	"api"
	model "model/response"
	"strings"
	"testing"
)

func TestGetOpcoesParcelamentoSucess(){
	const VALOR_PARCELA_EXPECTED := 1000.0
	 dto =  api.GetOpcoesParcelamentoAsync("1", "5940", "57.822.975/0001-12", "1000")
	 if &dto.Data == (*response.OpcoesParcelamento)(nil) {
		t.Error("Test failed")
	 }
	 if dto.Data[0].ValorParcela != VALOR_PARCELA_EXPECTED{
		t.Error("Test failed")
	 }
}
//ERRO FORA DO PADRÃO(objeto vem todo vazio)
func TestGetOpcoesParcelamentoFail(){
	dto :=  api.GetOpcoesParcelamento("8", "5940", "57.822.975/0001-12", "1000")
	if &dto.Data == (*response.OpcoesParcelamento)(nil) || len(dto.Data) > 0{
		t.Error("Test failed")
	 }
}