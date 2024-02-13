package main

import (
	"fmt"
	"glox/expression"
	"glox/printer"
	"glox/scanner"
)

func main() {
	source := "var num = 1;"

	scan := scanner.Scanner{}
	scan.Source = source

	scan.ScanTokens()
	fmt.Println("Tokens:")
	for _, token := range scan.Tokens {
		fmt.Println(token.ToString())
	}

	// Test expression parsing
	// Expr expression = new Expr.Binary(
	// 	new Expr.Unary(
	// 		new Token(TokenType.MINUS, "-", null, 1),
	// 		new Expr.Literal(123)),
	// 	new Token(TokenType.STAR, "*", null, 1),
	// 	new Expr.Grouping(
	// 		new Expr.Literal(45.67)));

	expr := expression.Binary{
		Left: &expression.Unary{
			Right: &expression.Literal{
				Value: "123",
			},
			Operator: scanner.Token{
				TokenType: scanner.MINUS,
				Lexeme:    "-",
				Literal:   nil,
				Line:      1,
			},
		},
		Operator: scanner.Token{
			TokenType: scanner.STAR,
			Lexeme:    "*",
			Literal:   nil,
			Line:      1,
		},
		Right: &expression.Grouping{
			Expression: &expression.Literal{
				Value: "45.67",
			},
		},
	}

	print := printer.Printer{}
	fmt.Println(print.Print(&expr))
}
