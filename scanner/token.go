package scanner

import "fmt"

type Token struct {
	tokenType TokenType
	Lexeme    string
	Literal   interface{}
	line      int
}

func (t *Token) ToString() string {
	return fmt.Sprintf("%v %v %v", StringTokenType[int(t.tokenType)], t.Lexeme, t.Literal)
}
