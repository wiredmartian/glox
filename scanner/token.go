package scanner

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   *interface{}
	line      int
}

func (t *Token) toString() string {
	return fmt.Sprintf("%v %v %v", t.tokenType, t.lexeme, t.literal)
}
