package api

import (
	"service"
	"fmt"
	"encoding/json"
)

func GetCampanhas(dtInicio string, dtFim string) Order {
	var queryParams = make(map[string]interface{})
	queryParams["dataInicio"] = dtInicio
	queryParams["dataFim"] = dtFim

	bodyString := service.Get("/campanhas", queryParams)
	bodyBytes := []byte (bodyString)

	// Convert response body to Order struct
	var orderStruct Order
	json.Unmarshal(bodyBytes, &orderStruct)
	fmt.Printf("\n%+v\n", orderStruct)

	return orderStruct
}

func GetFormasPagamento(idCampanha string, cnpj string) string {
	var queryParams = make(map[string]interface{})
	queryParams["cnpj"] = cnpj
	return service.Get("/campanhas/"+idCampanha+"/formas-pagamento/opcoes-parcelamento", queryParams)
}
