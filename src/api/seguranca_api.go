package api

import (
	"encoding/json"
	"fmt"
	model "model/response"
	"service"
)

func GetChave() model.ChaveDTO {
	var bodyString string = service.Get("/seguranca/chaves", make(map[string]interface{}))
	bodyBytes := []byte(bodyString)

	// Convert response body to struct
	var dtoStruct model.ChaveDTO
	error := json.Unmarshal(bodyBytes, &dtoStruct)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\n%+v\n", dtoStruct)

	return dtoStruct
}
