package api

import (
	"service"
)

func GetOpcoesParcelamento(idFormaPagamento string, idCampanha string, cnpj string, valorParcela string) string {
	var queryParams = make(map[string]interface{})
	queryParams["idCampanha"] = idCampanha
	queryParams["cnpj"] = cnpj
	queryParams["valorParcelar"] = valorParcela
	return service.Get("/formas-pagamento/"+idFormaPagamento+"/opcoes-parcelamento", queryParams)
}
