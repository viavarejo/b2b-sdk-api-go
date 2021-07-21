package main

import "fmt"

func main() {
	key := "ABCDEFGHIJKLM"
	// Omit start index, this is the same as zero.
	// ... Take 2-char substring.
	//substring := key[0:0]
	fmt.Println(key[3:4])

	var resultado string
	var k int32 = 1
	for i := 0; i < len([]rune(key)); i++ {
		resultado += key[i : i+1]
		if k == 2 {
			resultado += "\n"
			k = 1
		} else {
			k++
		}
	}

	fmt.Println(resultado)

}
