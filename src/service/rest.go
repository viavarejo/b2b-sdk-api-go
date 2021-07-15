package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

//var basePath string = "http://api-integracao-extra.hlg-b2b.net"
//var token string = "H9xO4+R8GUy+18nUCgPOlg=="

func doRequest(jsonData []byte, path string, method string, query map[string]interface{}) string {
	basePath := GetVariable("HOST_BANDEIRA")
	token := GetVariable("TOKEN_PARCEIRO")
	//fmt.Printf("godotenv : %s = %s \n", "HOST_BANDEIRA", basePath)
	//fmt.Printf("godotenv : %s = %s \n", "TOKEN_PARCEIRO", token)

	var reqbody *bytes.Buffer = nil
	if jsonData != nil {
		reqbody = bytes.NewBuffer(jsonData)
	}
	httpurl := basePath + path
	if len(query) > 0 {
		httpurl += "?" + BuildHttpQuery(query)
	}

	fmt.Println("HTTP JSON POST URL:", httpurl)
	request, error := http.NewRequest(method, httpurl, reqbody)
	if error != nil {
		panic(error)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body)
}

func Post(jsonData []byte, path string) string {
	return doRequest(jsonData, path, "POST", nil)
}

func Get(path string, query map[string]interface{}) string {
	return doRequest(nil, path, "GET", query)
}

func Patch(path string, query map[string]interface{}) string {
	return doRequest(nil, path, "PATCH", query)
}

func BuildHttpQuery(data map[string]interface{}) string {
	query := make([]string, 0, len(data))
	for k, v := range data {
		query = append(query, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(query, "&")
}

func GetVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro carregando arquivo .env")
	}
	return os.Getenv(key)
}
