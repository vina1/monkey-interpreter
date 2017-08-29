package token

type TokenType string

type Token struct { 
	Type TokenType 
	Literal string
 } 

const ( 
	ILLEGAL = "ILLEGAL" 
	EOF = "EOF"
	
	// Identifiers + literals 
	IDENT = "IDENT" // add, foobar, x, y, ... 
	INT = "INT" // 1343456
	
	// Operators
 	ASSIGN = "="
	PLUS = "+"
	MINUS = "-" 
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LESSTHAN = "<"
	GREATERTHAN = ">"
	
	// Delimiters
 	COMMA = "," 
 	SEMICOLON = ";"
	LPAREN = "(" 
	RPAREN = ")"
 	LBRACE = "{" 
 	RBRACE = "}"
	
	// Keywords 
	FUNCTION = "FUNCTION" 
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	// Booleans
	TRUE = "TRUE"
	FALSE = "FALSE"
) 

var keywords = map[string]TokenType{ 
	"fn": FUNCTION,
	"let": LET, 
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

func NewTokenFromChar(tokenType TokenType, char byte) Token { 
	return Token{Type: tokenType, Literal: string(char)} 
} 

func NewToken(tokenType TokenType, str string) Token { 
	return Token{Type: tokenType, Literal: str} 
} 

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok { 
		return tok 
	} 
	return IDENT 
} 