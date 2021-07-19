package main

import "fmt"

func main() {
	bodyString := (*string)(nil)
	bodyBytes := []byte(bodyString)
	fmt.Println(bodyBytes)
}
