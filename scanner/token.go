package scanner

import "fmt"

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
	Line      int
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%v %v %v", StringTokenType[int(t.TokenType)], t.Lexeme, t.Literal)
}
