package scanner

import "strings"

type Scanner struct {
	Source string
	Tokens []Token

	start   int
	current int
	line    int
}

func (s *Scanner) scanTokens() {

	for s.current >= len(s.Source) {
		s.start = s.current
		s.scanTokens()
	}

	s.Tokens = append(s.Tokens, Token{
		tokenType: EOF,
		lexeme:    "",
		literal:   nil,
		line:      s.line,
	})
}

func (s *Scanner) scanToken() {
	// Single character lexemes
	char := s.next()
	switch char {
	case "(":
		s.addToken(LEFT_PAREN)
		break
	case ")":
		s.addToken(RIGHT_PAREN)
		break
	case "{":
		s.addToken(LEFT_BRACE)
		break
	case "}":
		s.addToken(RIGHT_BRACE)
		break
	case ",":
		s.addToken(COMMA)
		break
	case ".":
		s.addToken(DOT)
		break
	case "-":
		s.addToken(MINUS)
		break
	case "+":
		s.addToken(PLUS)
		break
	case ";":
		s.addToken(SEMICOLON)
		break
	case "*":
		s.addToken(STAR)
		break
	}
}

func (s *Scanner) next() string {
	s.current++
	return strings.Split(s.Source, "")[s.current-1]
}

func (s *Scanner) addToken(tType TokenType) {
	s.addTokenLiteral(tType, nil)
}

func (s *Scanner) addTokenLiteral(tType TokenType, literal *interface{}) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, Token{
		tokenType: tType,
		lexeme:    text,
		literal:   literal,
		line:      s.line,
	})
}
