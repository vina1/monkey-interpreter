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

func (l *Lexer) peekChar() byte { 
	if l.readPosition >= len(l.input) { 
		return 0 
	} else { 
		return l.input[l.readPosition] 
	} 
} 

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentChar {
		case '=':
			if l.peekChar() == '=' { 
				ch := l.currentChar 
				l.readChar() 
				literal := string(ch) + string(l.currentChar) 
				tok = token.NewToken(token.EQUAL, literal) 
			} else { 
				tok = token.NewTokenFromChar(token.ASSIGN, l.currentChar)
			}
		case ';':
			tok = token.NewTokenFromChar(token.SEMICOLON, l.currentChar)
		case '(':
			tok = token.NewTokenFromChar(token.LPAREN, l.currentChar)
		case ')':
			tok = token.NewTokenFromChar(token.RPAREN, l.currentChar)
		case ',':
			tok = token.NewTokenFromChar(token.COMMA, l.currentChar)
		case '+':
			tok = token.NewTokenFromChar(token.PLUS, l.currentChar)
		case '-':
			tok = token.NewTokenFromChar(token.MINUS, l.currentChar)
		case '!':
			if l.peekChar() == '=' { 
				ch := l.currentChar 
				l.readChar() 
				literal := string(ch) + string(l.currentChar) 
				tok = token.NewToken(token.NOT_EQUAL, literal) 
			} else { 
				tok = token.NewTokenFromChar(token.BANG, l.currentChar)
			}
		case '*':
			tok = token.NewTokenFromChar(token.ASTERISK, l.currentChar)
		case '/':
			tok = token.NewTokenFromChar(token.SLASH, l.currentChar)
		case '<':
			tok = token.NewTokenFromChar(token.LESS_THAN, l.currentChar)
		case '>':
			tok = token.NewTokenFromChar(token.GREATER_THAN, l.currentChar)
		case '{':
			tok = token.NewTokenFromChar(token.LBRACE, l.currentChar)
		case '}':
			tok = token.NewTokenFromChar(token.RBRACE, l.currentChar)
		case 0:
			tok = token.NewToken(token.EOF, "")
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
				tok = token.NewTokenFromChar(token.ILLEGAL, l.currentChar)
			}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position 
	
	for isLetter(l.currentChar) { 
		l.readChar() 
	} 
	
	return l.input[startPosition:l.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' 
} 

func (l *Lexer) skipWhitespace() { 
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' { 
		l.readChar()
	} 
} 

func (l *Lexer) readNumber() string { 
	startPosition := l.position 
	
	for isDigit(l.currentChar) { 
		l.readChar() 
	} 
	return l.input[startPosition:l.position] 
}

func isDigit(char byte) bool { 
	return '0' <= char && char <= '9' 
} 