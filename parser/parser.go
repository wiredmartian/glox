package parser

import (
	"glox/expr"
	"glox/scanner"
)

type Parser struct {
	tokens  []scanner.Token
	current int
}

func (p *Parser) Parse(tokens []scanner.Token) {
	p.tokens = tokens
	p.current = 0
}

func expression() expr.Expr {
	return nil
}
