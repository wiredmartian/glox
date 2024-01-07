package scanner

type Keywords map[string]TokenType

// reserved for the language
var keywords = Keywords{
	"var":    VAR,
	"class":  CLASS,
	"nil":    NIL,
	"fun":    FUN,
	"or":     OR,
	"and":    AND,
	"if":     IF,
	"else":   ELSE,
	"false":  FALSE,
	"true":   TRUE,
	"for":    FOR,
	"while":  WHILE,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
}
