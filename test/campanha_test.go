package main

import (
	"fmt"
	"testing"
	campanhaApi "api"
)


func TestGetCampanhas(t *testing.T){
	campanhas := campanhaApi.GetCampanhas("2019-08-04", "2100-08-04")
	fmt.Println("Response:")
	fmt.Println(campanhas)

	expected := "57.822.975/0001-12"
	actual := campanhas
  	if strings.Contains(actual, expected) {
		t.Error("Test failed, Deveria retornar a campanha - esperado: '%s, retornou:  '%s", expected, actual)
  	}
}