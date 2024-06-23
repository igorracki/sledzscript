package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL           = "ILLEGAL"
	EOF               = "EOF"
	IDENTIFIER        = "IDENTIFIER"
	INT               = "INT"
	ASSIGNMENT        = "="
	PLUS              = "+"
	MINUS             = "-"
	BANG              = "!"
	ASTERISK          = "*"
	SLASH             = "/"
	GREATER_THAN      = ">"
	LESS_THAN         = "<"
	COMMA             = ","
	SEMICOLON         = ";"
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"
	FUNCTION          = "FUNCTION"
	LET               = "LET"
	RETURN            = "RETURN"
	IF                = "IF"
	ELSE              = "ELSE"
	TRUE              = "TRUE"
	FALSE             = "FALSE"
	EQUAL             = "=="
	NOT_EQUAL         = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(identifier string) TokenType {
	if matchedToken, ok := keywords[identifier]; ok {
		return matchedToken
	}
	return IDENTIFIER
}
