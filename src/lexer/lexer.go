package lexer

import ( 
	"token"
)

type Lexer struct { 
	input string 
	position int // current position in input (points to current char) 
	readPosition int // current reading position in input (after current char) 
	currentChar byte // current char under examination 
}

func New(input string) *Lexer { 
	l := &Lexer{input: input}
	l.readChar() 
	return l 
} 

func (l *Lexer) readChar() { 
	if l.readPosition >= len(l.input) { 
		l.currentChar = 0 
	} else { 
		l.currentChar = l.input[l.readPosition] 
	} 
	
	l.position = l.readPosition

	l.readPosition += 1
} 

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentChar {
		case '=':
			tok = newToken(token.ASSIGN, l.currentChar)
		case ';':
			tok = newToken(token.SEMICOLON, l.currentChar)
		case '(':
			tok = newToken(token.LPAREN, l.currentChar)
		case ')':
			tok = newToken(token.RPAREN, l.currentChar)
		case ',':
			tok = newToken(token.COMMA, l.currentChar)
		case '+':
			tok = newToken(token.PLUS, l.currentChar)
		case '{':
			tok = newToken(token.LBRACE, l.currentChar)
		case '}':
			tok = newToken(token.RBRACE, l.currentChar)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.currentChar) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if isDigit(l.currentChar) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.currentChar)
			}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token { 
	return token.Token{Type: tokenType, Literal: string(ch)} 
} 

func (l *Lexer) readIdentifier() string {
	position := l.position 
	
	for isLetter(l.currentChar) { 
		l.readChar() 
	} 
	
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' 
} 

func (l *Lexer) skipWhitespace() { 
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' { 
		l.readChar()
	} 
} 

func (l *Lexer) readNumber() string { 
	position := l.position 
	
	for isDigit(l.currentChar) { 
		l.readChar() 
	} 
	return l.input[position:l.position] 
}

func isDigit(ch byte) bool { 
	return '0' <= ch && ch <= '9' 
} 