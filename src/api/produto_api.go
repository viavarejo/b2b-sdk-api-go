package api

import (
	"service"
)

func GetDadosProduto(idLogista string, idSKu string) string {
	return service.Get("/lojistas/"+idLogista+"/produtos/"+idSKu, make(map[string]interface{}))
}

func GetListaDadosProdutos(idLogista string, idSKus ...string) string {
	var queryParams = make(map[string]interface{})
	for _, idSku := range idSKus {
		queryParams["idSku"] = idSku
	}
	return service.Get("/lojistas/"+idLogista+"/produtos", queryParams)
}

func GetDadosProdutoCampanha(idCampanha string, idSKu string, cnpj string, idLojista string) string {
	var queryParams = make(map[string]interface{})
	queryParams["idLojista"] = idLojista
	queryParams["cnpj"] = cnpj
	return service.Get("/campanhas/"+idCampanha+"/produtos/"+idSKu, queryParams)
}
