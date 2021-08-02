package api

import (
	"encoding/json"

	model "github.com/viavarejo/b2b-sdk-api-go/src/model/response"
	"github.com/viavarejo/b2b-sdk-api-go/src/service"
)

func GetOpcoesParcelamento(idFormaPagamento string, idCampanha string, cnpj string, valorParcela string) model.OpcoesParcelamentoDTO {
	var queryParams = make(map[string]interface{})
	queryParams["idCampanha"] = idCampanha
	queryParams["cnpj"] = cnpj
	queryParams["valorParcelar"] = valorParcela
	bodyString := service.Get("/formas-pagamento/"+idFormaPagamento+"/opcoes-parcelamento", queryParams)
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.OpcoesParcelamentoDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	//fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}
