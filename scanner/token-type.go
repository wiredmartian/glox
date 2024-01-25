package scanner

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

// Maps iota token to strings
var StringTokenType = map[int]string{
    int(LEFT_PAREN):    "LEFT_PAREN",
    int(RIGHT_PAREN):   "RIGHT_PAREN",
    int(LEFT_BRACE):    "LEFT_BRACE",
    int(RIGHT_BRACE):   "RIGHT_BRACE",
    int(COMMA):         "COMMA",
    int(DOT):           "DOT",
    int(MINUS):         "MINUS",
    int(PLUS):          "PLUS",
    int(SEMICOLON):     "SEMICOLON",
    int(SLASH):         "SLASH",
    int(STAR):          "STAR",
    int(BANG):          "BANG",
    int(BANG_EQUAL):    "BANG_EQUAL",
    int(EQUAL):         "EQUAL",
    int(EQUAL_EQUAL):   "EQUAL_EQUAL",
    int(GREATER):       "GREATER",
    int(GREATER_EQUAL): "GREATER_EQUAL",
    int(LESS):          "LESS",
    int(LESS_EQUAL):    "LESS_EQUAL",
    int(IDENTIFIER):    "IDENTIFIER",
    int(STRING):        "STRING",
    int(NUMBER):        "NUMBER",
    int(AND):           "AND",
    int(CLASS):         "CLASS",
    int(ELSE):          "ELSE",
    int(FALSE):         "FALSE",
    int(FUN):           "FUN",
    int(FOR):           "FOR",
    int(IF):            "IF",
    int(NIL):           "NIL",
    int(OR):            "OR",
    int(PRINT):         "PRINT",
    int(RETURN):        "RETURN",
    int(SUPER):         "SUPER",
    int(THIS):          "THIS",
    int(TRUE):          "TRUE",
    int(VAR):           "VAR",
    int(WHILE):         "WHILE",
    int(EOF):           "EOF",
}