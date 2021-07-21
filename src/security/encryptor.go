package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func Encrypt(publicKeyString string, secretMessage string) string {
	//to tests: publicKeyString := "MIIENTCCAx2gAwIBAgIJAJ5ApEGl2oaIMA0GCSqGSIb3DQEBBQUAMIGwMQswCQYDVQQGEwJCUjELMAkGA1UECAwCU1AxFDASBgNVBAcMC1NBTyBDQUVUQU5PMRMwEQYDVQQKDApWSUEgVkFSRUpPMSAwHgYDVQQLDBdTRUdVUkFOQ0EgREEgSU5GT1JNQUNBTzEOMAwGA1UEAwwFUFJPWFkxNzA1BgkqhkiG9w0BCQEWKHRpLnNlZ3VyYW5jYS5pbmZvcm1hY2FvQHZpYXZhcmVqby5jb20uYnIwHhcNMTgwODE2MTIzNjQ2WhcNMjEwODE1MTIzNjQ2WjCBsDELMAkGA1UEBhMCQlIxCzAJBgNVBAgMAlNQMRQwEgYDVQQHDAtTQU8gQ0FFVEFOTzETMBEGA1UECgwKVklBIFZBUkVKTzEgMB4GA1UECwwXU0VHVVJBTkNBIERBIElORk9STUFDQU8xDjAMBgNVBAMMBVBST1hZMTcwNQYJKoZIhvcNAQkBFih0aS5zZWd1cmFuY2EuaW5mb3JtYWNhb0B2aWF2YXJlam8uY29tLmJyMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqObNb7KAP09WsV9h76Dw3tj2qa3l97K+slfzLkOBvi0xjacuKCnvsMSGEBosvWY/qNmSLE1YaoyFt7ZaeOiALKh2AFckJRM+/zvQzqi6cPnW0cGsEE/9WO48Fgh894pKjHpukATFb9tBYGTBEW46AH2WiAR735KEnDfFAHG//pkLKriPWEZBr9tf4gdNvyJ/ybs5JrBRU1RKE9MM7qnMkCouKTPwY/lS/2Xb1IYkyZulCf3Uyl7zpB6hQUhprS1R5meRocpGgHJCFfiWD/uXa5nREuGuQxcImwzvf+enwT6CooRoM2rN6IQWSY+uQ64dhSt4FMajZFmHVpLfUIOjEwIDAQABo1AwTjAdBgNVHQ4EFgQUZ22K62aMm/lI5LfblgINPvz8ae8wHwYDVR0jBBgwFoAUZ22K62aMm/lI5LfblgINPvz8ae8wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEAj23IDXLPkQpFDbgAtgKuO9N66o61edbJ1+BMjdSsfO0vMVpmBDlKdinxlh509/qJm/WLYswKkKOi7VHojBSV5HyrO5YGCSJFvVGJqF4JUxy7GrWTHqgwcylmX5B5lNd5aMIxwG6AF4o2cp6IPe+Uwaroa8kLTrtM0eRgAInHbQA7MXbvOZY+pzE4s6jFbA1O321zVg4C4Y3C4e30yf9YJNK5XjUP26duvwGqQrZg49ZU3W/t6GYY1kQhSeBG0FPg2GOIHX03WPZpaJ7i1uCv6Ial07pxDxqcT8oCJalY9tW9sv7zBJRaJgTIf5oz5jElb9kWd2D6XwaGB5PJfD6CTQ=="

	//privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	//CheckError(err)
	//publicKey := privateKey.PublicKey

	publicKeyPem := ConvertToPem(publicKeyString)
	publicKeyBytes := []byte(publicKeyPem)
	var publicKey rsa.PublicKey = BytesToPublicKey(publicKeyBytes)

	encryptedMessage := PKCS1_Encrypt(secretMessage, publicKey)

	fmt.Println("Cipher Text:", encryptedMessage)
	return encryptedMessage
}

func PKCS1_Encrypt(secretMessage string, key rsa.PublicKey) string {
	//label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptPKCS1v15(rng, &key, []byte(secretMessage))
	CheckError(err)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) rsa.PublicKey {
	block, _ := pem.Decode([]byte(pub))
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
	fmt.Println(rsaPublicKey.N)
	fmt.Println(rsaPublicKey.E)
	return *rsaPublicKey
}

func CheckError(e error) {
	if e != nil {
		fmt.Println("Erro: ")
		fmt.Println(e.Error())
	}
}

func ConvertToPem(key string) string {
	var split = 64
	var k = 1
	var resultado string
	for i := 0; i < len([]rune(key)); i++ {
		resultado += key[i : i+1]
		if k == split {
			resultado += "\n"
			k = 1
		} else {
			k++
		}
	}
	resultado = "-----BEGIN CERTIFICATE-----\n" + resultado + "\n-----END CERTIFICATE-----"
	fmt.Println(resultado)
	return resultado
}
