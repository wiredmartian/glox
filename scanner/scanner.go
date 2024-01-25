package scanner

import (
	"fmt"
	"strconv"
	"strings"
)

type Scanner struct {
	Source string
	Tokens []Token

	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() {

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
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
	case "/":
		// A slash could be a start of a comment or a division sign
		if s.match("/") {
			// A comment goes until the end of the line.
			// Single line comment
			for s.peek() != "\n" && !s.isAtEnd() {
				s.next()
			}
		} else if s.match("*") {
			// multi-line comment support
			for s.peek() != "/" && s.peekAt(s.current-1) != "*" && !s.isAtEnd() {
				s.next()
			}

		} else {
			s.addToken(SLASH)
		}
		break
	case " ":
	case "\r":
	case "\t":
		// Ignore whitespace.
		break
	case "\n":
		s.line++
		break
	case "\"":
		s.string()
		break
	default:
		if s.isDigit(char) {
			s.number()
		} else if s.isAlpha(char) {
			s.identifier()
		} else {
			fmt.Printf("Unexpected character at line: %v", s.line)
		}
		break
	}

}

func (s *Scanner) addToken(tType TokenType) {
	s.addTokenLiteral(tType, nil)
}

func (s *Scanner) addTokenLiteral(tType TokenType, literal interface{}) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, Token{
		tokenType: tType,
		lexeme:    text,
		literal:   literal,
		line:      s.line,
	})
}

// Helper functions
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

func (s *Scanner) isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

func (s *Scanner) isAlpha(char string) bool {
	return (char >= "a" && char <= "z") ||
		(char >= "A" && char <= "Z") || char == "_"
}

func (s *Scanner) isAlphanumeric(char string) bool {
	return s.isAlpha(char) || s.isDigit(char)
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\\0"
	}
	return strings.Split(s.Source, "")[s.current]
}

func (s *Scanner) peekAt(position int) string {
	return strings.Split(s.Source, "")[position]
}

func (s *Scanner) peekNext() string {
	if s.current+1 > len(s.Source) {
		return "\\0"
	}
	return s.Source[s.current+1:]
}

func (s *Scanner) string() {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() != "\n" {
			s.line++
		}
		s.next()
	}
	if s.isAtEnd() {
		fmt.Printf("Unterminated string at line: %v", s.line)
		return
	}
	// The closing ".
	s.next()

	// Trim the surrounding quotes.
	value := s.Source[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, value)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.next()
	}
	// Look for a fractional part.
	if s.peek() == "." && s.isDigit(s.peekNext()) {
		// Consume the "."
		s.next()

		for s.isDigit(s.peek()) {
			s.next()
		}
	}
	value, err := strconv.ParseFloat(s.Source[s.start:s.current], 32)
	if err != nil {
		fmt.Printf("Unable to convert %v to a numerical literal value", value)
	}
	s.addTokenLiteral(NUMBER, value)
}

func (s *Scanner) identifier() {
	for s.isAlphanumeric(s.peek()) {
		s.next()
	}
	text := s.Source[s.start:s.current]
	var tokenType TokenType = keywords[text]
	if &tokenType == nil {
		tokenType = IDENTIFIER
	}
	s.addToken(IDENTIFIER)
}
