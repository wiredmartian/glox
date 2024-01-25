package main

import (
	"fmt"
	"glox/scanner"
)

func main() {
	source := "var num = 1;"

	scanner := scanner.Scanner{}
	scanner.Source = source

	scanner.ScanTokens()
	fmt.Println("Tokens:")
	for _, token := range scanner.Tokens {
		fmt.Println(token.ToString())
	}
}
