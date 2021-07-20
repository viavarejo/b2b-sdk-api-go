package api

import (
	"encoding/json"
	"fmt"

	model "github.com/viavarejo/b2b-sdk-api-go/src/model/response"
	"github.com/viavarejo/b2b-sdk-api-go/src/service"
)

func GetDadosProduto(idLogista string, idSKu string) model.ProdutoDTO {
	bodyString := service.Get("/lojistas/"+idLogista+"/produtos/"+idSKu, make(map[string]interface{}))
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.ProdutoDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}

func GetListaDadosProdutos(idLogista string, idSKus []string) model.ProdutosDTO {
	var queryParams = make(map[string]interface{})
	for _, idSku := range idSKus {
		queryParams["idSku"] = idSku
	}
	bodyString := service.Get("/lojistas/"+idLogista+"/produtos", queryParams)
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.ProdutosDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}

func GetDadosProdutoCampanha(idCampanha string, idSKu string, cnpj string, idLojista string) model.ProdutoDTO {
	var queryParams = make(map[string]interface{})
	queryParams["idLojista"] = idLojista
	queryParams["cnpj"] = cnpj
	bodyString := service.Get("/campanhas/"+idCampanha+"/produtos/"+idSKu, queryParams)
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.ProdutoDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}
