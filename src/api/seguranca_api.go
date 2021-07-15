package api

import (
	"service"
)

func GetChave() string {
	return service.Get("/seguranca/chaves", make(map[string]interface{}))
}
