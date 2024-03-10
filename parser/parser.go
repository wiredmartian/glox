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

// / equality -> comparison ( ( "!=" | "==" ) comparison )*
func (p *Parser) equality() expr.Expr {
	exp := p.comparison()
	for p.match(scanner.BANG, scanner.BANG_EQUAL) {
		// Right, so if we are inside the while loop in equality() , then we know we have found a
		// != or == operator and must be parsing an equality expression.

		// We grab the matched operator token so we can track which kind of equality expression we have. Then we call comparison() again to parse the right-hand operand. We combine the operator and its two operands into a new
		// Expr.Binary syntax tree node, and then loop around. Each iteration, we store the resulting expression back in the same expr local variable. As we zip through a sequence of equality expressions, that creates a left-associative nested tree of binary operator nodes.
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

// / comparison -> term ( ( ">" | ">=" | "<" | "<=" ) term )*
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

// / factor -> unary ( ( "/" | "*" ) unary )*
func (p *Parser) factor() expr.Expr {
	exp := p.unary()
	for p.match(scanner.SLASH, scanner.STAR) {
		operator := p.previous()
		right := p.unary()
		exp = &expr.Binary{
			Left:     exp,
			Operator: operator,
			Right:    right,
		}
	}
	return exp
}

// / unary -> ( "!" | "-" ) unary | primary
func (p *Parser) unary() expr.Expr {
	if p.match(scanner.BANG, scanner.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &expr.Unary{
			Operator: operator,
			Right:    right,
		}
	}
	return p.primary()
}

// / primary -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")"
func (p *Parser) primary() expr.Expr {
	if p.match(scanner.FALSE) {
		return &expr.Literal{Value: false}
	}
	if p.match(scanner.TRUE) {
		return &expr.Literal{Value: true}
	}
	if p.match(scanner.NIL) {
		return &expr.Literal{Value: nil}
	}
	if p.match(scanner.NUMBER, scanner.STRING) {
		return &expr.Literal{Value: p.previous().Literal}
	}
	if p.match(scanner.LEFT_PAREN) {
		expression := p.expression()
		p.match(scanner.RIGHT_PAREN) // FIXME: Handle error
		return &expr.Grouping{Expression: expression}
	}
	return nil
}

// / term -> factor ( ( "-" | "+" ) factor )*
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
