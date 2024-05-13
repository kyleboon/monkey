package lexer

import (
	"monkey/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input :=
		`let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		};

		+-/*^%
		<>

		==
		!=

   		let result = add(five, ten);
	    if return true
		else false
	    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.POWER, "^"},
		{token.MODULUS, "%"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EQ, "=="},
		{token.NE, "!="},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.ELSE, "else"},
		{token.FALSE, "false"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)

		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}