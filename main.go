package main

import (
	"fmt"
	"glox/scanner"
)

func main() {
	source := "() {} "
	scanner := scanner.Scanner{
		Source: source,
	}

	scanner.ScanTokens()

	fmt.Printf("Tokens: %v", scanner.Tokens)
}
