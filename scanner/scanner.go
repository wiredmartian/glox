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

	for !s.isAtEnd() {
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
	// Single character lexemes (e.g. ( ) { } , . - + ; *)

	// Two character lexemes (e.g. != == >= <=)

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
	case "!":
		if s.match("=") {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
		break
	case "=":
		if s.match("=") {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
		break
	case "<":
		if s.match("=") {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
		break
	case ">":
		if s.match("=") {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
		break
	}

}

// next character in source to scan
func (s *Scanner) next() string {
	s.current++
	return strings.Split(s.Source, "")[s.current-1]
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if strings.Split(s.Source, "")[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\\0"
	}
	return strings.Split(s.Source, "")[s.current]
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
