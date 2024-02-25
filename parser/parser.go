package parser

import (
	"glox/expr"
	"glox/scanner"
)

// Presidency and association seems to be another key concept here
type Parser struct {
	tokens  []scanner.Token
	current int
}

func (p *Parser) Parse(tokens []scanner.Token) {
	p.tokens = tokens
	p.current = 0
}

func (p *Parser) expression() expr.Expr {
	expression := p.comparison()

	for p.match(scanner.BANG_EQUAL, scanner.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expression = &expr.Binary{
			Left:     expression, // Need to revisit
			Operator: operator,
			Right:    right,
		}
	}
	return expression
}

func (p *Parser) equality() expr.Expr {
	exp := p.comparison()
	for p.match(scanner.BANG, scanner.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		exp = &expr.Binary{
			Left:     exp,
			Operator: operator,
			Right:    right,
		}
	}
	return exp
}

func (p *Parser) comparison() expr.Expr {
	exp := p.term()
	for p.match(scanner.GREATER, scanner.GREATER_EQUAL, scanner.LESS, scanner.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		exp = &expr.Binary{
			Left:     exp,
			Operator: operator,
			Right:    right,
		}
	}
	return exp
}

func (p *Parser) factor() expr.Expr {
	return nil
}

func (p *Parser) term() expr.Expr {
	exp := p.factor()
	for p.match(scanner.MINUS, scanner.PLUS) {
		operator := p.previous()
		right := p.factor()
		exp = &expr.Binary{
			Left:     exp,
			Operator: operator,
			Right:    right,
		}
	}
	return exp
}

// Helper functions

func (p *Parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

// The match() method returns true if the current token is of the given type.
//
// It also consumes the token
func (p *Parser) match(tokenTypes ...scanner.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == scanner.EOF
}

// The check() method returns true if the current token is of the given type
func (p *Parser) check(tokenType scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}
