package main

import (
	campanhaApi "api"
	"model/response"
	"strings"
	"testing"
)

func TestGetCampanhas(t *testing.T) {
	dtFim := "2100-08-04"
	campanhas := campanhaApi.GetCampanhas("2019-08-04", &dtFim)
	//fmt.Println("Response:")
	//fmt.Println(campanhas)

	expected := "57.822.975/0001-12"
	actual := campanhas
	if strings.Contains(actual.Data.CnpjContrato, expected) {
		t.Errorf("Test failed, Deveria retornar a campanha - esperado: '%s, retornou:  '%s", expected, actual.Data.CnpjContrato)
	}
}

func testGetCampanhaFail(t *testing.T) {
	campanhas := campanhaApi.GetCampanhas("2019-08-04", nil)

	if &campanhas.Data != (*response.Campanha)(nil) || campanhas.Error.Code != "400" || campanhas.Error.Message != "Request inválido\\r\\nA dataFim é um parâmetro obrigatório." {
		t.Error("Test failed")
	}
}

func TestGetFormasPagamentoSucess(t *testing.T) {
	dto := campanhaApi.GetFormasPagamento("5940", "57.822.975/0001-12")
	if &dto.Data == (*[]response.FormasPagamento)(nil) || dto.Data[0].IdFormaPagamento != 1 || dto.Data[0].Nome != "Cartão de Crédito Visa " {
		t.Error("Test failed")
	}
}

func TestGetFormasPagamentoFail(t *testing.T) {
	dto := campanhaApi.GetFormasPagamento("590", "57.822.975/0001-12")
	if &dto.Data == (*[]response.FormasPagamento)(nil) ||  &dto.Data != []response.FormasPagamento|| &dto.Error.Code != (*string)(nil) {
		t.Error("Test failed")
	}
}
