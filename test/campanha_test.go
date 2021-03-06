package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/viavarejo/b2b-sdk-api-go/src/api"
	"github.com/viavarejo/b2b-sdk-api-go/src/model/response"
)

func TestCampanhas(t *testing.T) {
	t.Run("Deveria retornar as Campanhas",
		func(t *testing.T) {
			var campanhas response.CampanhasDTO = api.GetCampanhas("2019-08-04", "2100-08-04")
			b, err := json.MarshalIndent(campanhas, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			expected := "57.822.975/0001-12"

			if campanhas.Data[0].CnpjContrato != expected {
				t.Errorf("Test failed, Deveria retornar a campanha - esperado: %s, retornou: %s", expected, campanhas.Data[0].CnpjContrato)
			}
		})

	t.Run("Falha ao retornar as Campanhas",
		func(t *testing.T) {
			var campanhas response.CampanhasDTO = api.GetCampanhas("2019-08-04", "")
			b, err := json.MarshalIndent(campanhas, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			if campanhas.Data != nil {
				t.Error("Test failed-1")
			}
			if campanhas.Error.Code != "400" {
				t.Error("Test failed-2")
			}
			if campanhas.Error.Message != "Request inválido\r\nA dataFim é um parâmetro obrigatório." {
				t.Error("Test failed-3")
			}
		})

	t.Run("Deveria retornar as formas de pagamento",
		func(t *testing.T) {
			var dto response.FormasPagamentoDTO = api.GetFormasPagamento("5940", "57.822.975/0001-12")
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			if len(dto.Data) == 0 {
				t.Error("Test failed-1")
			}
			if dto.Data[0].IdFormaPagamento != 1 {
				t.Error("Test failed-2")
			}
			if dto.Data[0].Nome != "Cartão de Crédito Visa " {
				t.Error("Test failed-3")
			}
		})

	t.Run("Falha ao retornar Formas de pagamento",
		func(t *testing.T) {
			var dto response.FormasPagamentoDTO = api.GetFormasPagamento("590", "57.822.975/0001-12")
			b, err := json.MarshalIndent(dto, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))

			if len(dto.Data) > 0 {
				t.Error("Test failed-1")
			}
			if len(dto.Error.Code) > 0 {
				t.Error("Test failed-2")
			}
		})
}
