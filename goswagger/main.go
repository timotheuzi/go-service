package main

import "fmt"

func main() {
	final := ""
	skuTemp := "011112KLSDJ32"
	firstCharacter := skuTemp[0:1]
	//arrayTest := []string(skuTemp)
	if firstCharacter != "0" {
		final := skuTemp
		fmt.Print(final)
	} else {
		length := len(skuTemp)
		final := skuTemp[1:length]
		fmt.Print(final)
	}

	fmt.Println(final)

	server := NewServer()
	server.ListenAndServe(":8080")

}
