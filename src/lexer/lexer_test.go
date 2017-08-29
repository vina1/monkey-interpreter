package lexer
import ( 
	"testing"

	"token"
)

// go test ./lexer

func TestOperators(t *testing.T) {
	input := `=+`

	tests := []token.Token { 
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.PLUS, "+"),
		token.NewToken(token.EOF, ""),
		}

	AssertExpectedVersusActual(t, input, tests)
}

func TestAdditionalOperators(t *testing.T) {
	input := `-!*/<>`

	tests := []token.Token { 
		token.NewToken(token.MINUS, "-"),
		token.NewToken(token.BANG, "!"),
	 	token.NewToken(token.ASTERISK, "*"),
		token.NewToken(token.SLASH, "/"),
		token.NewToken(token.LESSTHAN, "<"),
		token.NewToken(token.GREATERTHAN, ">"),
		token.NewToken(token.EOF, ""),
		}

	AssertExpectedVersusActual(t, input, tests)
}

func TestDelimeters(t *testing.T) {
	input := `(){},;`

	tests := []token.Token { 
		token.NewToken(token.LPAREN, "("),
		token.NewToken(token.RPAREN, ")"),
		token.NewToken(token.LBRACE, "{"),
		token.NewToken(token.RBRACE, "}"),
		token.NewToken(token.COMMA, ","),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.EOF, ""),
		}

	AssertExpectedVersusActual(t, input, tests)
}

func TestConditionalBranchAndBooleans(t *testing.T) {
	input := `if (5 < 10) { 
		return true; 
	} else { 
		return false; 
	}`

	tests := []token.Token { 
		token.NewToken(token.IF, "if"),
		token.NewToken(token.LPAREN, "("),
		token.NewToken(token.INT, "5"),
		token.NewToken(token.LESSTHAN, "<"),
		token.NewToken(token.INT, "10"),
		token.NewToken(token.RPAREN, ")"),
		token.NewToken(token.LBRACE, "{"),
		token.NewToken(token.RETURN, "return"),
  		token.NewToken(token.TRUE, "true"),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.RBRACE, "}"),
		token.NewToken(token.ELSE, "else"),
		token.NewToken(token.LBRACE, "{"),
		token.NewToken(token.RETURN, "return"),
  		token.NewToken(token.FALSE, "false"),
	    token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.RBRACE, "}"),
		token.NewToken(token.EOF, ""),
		}

	AssertExpectedVersusActual(t, input, tests)
}

func TestComplex(t *testing.T) {
	input := `let five = 5; let ten = 10;
	let add = fn(x, y) { x + y; };
	let result = add(five, ten); `

	tests := []token.Token{
		token.NewToken(token.LET, "let"), 
		token.NewToken(token.IDENT, "five"), 
		token.NewToken(token.ASSIGN, "="), 
		token.NewToken(token.INT, "5"), 
		token.NewToken(token.SEMICOLON, ";"), 
		token.NewToken(token.LET, "let"), 
		token.NewToken(token.IDENT, "ten"), 
		token.NewToken(token.ASSIGN, "="), 
		token.NewToken(token.INT, "10"), 
		token.NewToken(token.SEMICOLON, ";"), 
		token.NewToken(token.LET, "let"), 
		token.NewToken(token.IDENT, "add"), 
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.FUNCTION, "fn"),
		token.NewToken(token.LPAREN, "("),
		token.NewToken(token.IDENT, "x"),
		token.NewToken(token.COMMA, ","),
		token.NewToken(token.IDENT, "y"),
		token.NewToken(token.RPAREN, ")"),
		token.NewToken(token.LBRACE, "{"),
		token.NewToken(token.IDENT, "x"),
		token.NewToken(token.PLUS, "+"),
		token.NewToken(token.IDENT, "y"),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.RBRACE, "}"), 
		token.NewToken(token.SEMICOLON, ";"), 
		token.NewToken(token.LET, "let"), 
		token.NewToken(token.IDENT, "result"), 
		token.NewToken(token.ASSIGN, "="), 
		token.NewToken(token.IDENT, "add"), 
		token.NewToken(token.LPAREN, "("), 
		token.NewToken(token.IDENT, "five"), 
		token.NewToken(token.COMMA, ","), 
		token.NewToken(token.IDENT, "ten"),
		token.NewToken(token.RPAREN, ")"),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.EOF, ""),
	} 

	AssertExpectedVersusActual(t, input, tests)
}

func AssertExpectedVersusActual(t *testing.T, input string, expectedTokens []token.Token) {
	l := New(input)

	for i, expectedToken := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expectedToken.Type, tok.Type) 
		}

		if tok.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expectedToken.Literal, tok.Literal) 
		}
	}
}