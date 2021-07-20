package api

import (
	"encoding/json"
	"fmt"

	model "github.com/viavarejo/b2b-sdk-api-go/src/model/response"
	"github.com/viavarejo/b2b-sdk-api-go/src/service"
)

func GetCampanhas(dtInicio string, dtFim *string) model.CampanhasDTO {
	var queryParams = make(map[string]interface{})
	queryParams["dataInicio"] = dtInicio
	queryParams["dataFim"] = *dtFim

	bodyString := service.Get("/campanhas", queryParams)
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.CampanhasDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}

func GetFormasPagamento(idCampanha string, cnpj string) model.FormasPagamentoDTO {
	var queryParams = make(map[string]interface{})
	queryParams["cnpj"] = cnpj
	bodyString := service.Get("/campanhas/"+idCampanha+"/formas-pagamento/opcoes-parcelamento", queryParams)
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.FormasPagamentoDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}
