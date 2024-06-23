package lexer

import (
	"igorracki/sledzscript/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.ASSIGNMENT, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESIS, "("},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range tests {
		token := lexer.NextToken()
		if token.Type != expected.Type {
			t.Fatalf("i = %d - Expected type [%q], got [%q]", i, expected.Type, token.Type)
		}
		if token.Literal != expected.Literal {
			t.Fatalf("i = %d - Expected Literal [%q], got [%q]", i, expected.Literal, token.Literal)
		}
	}
}

func TestBasicSnippet(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		return x + y;
	};

	let result = add(five, ten);
	`
	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.LET, "let"}, {token.IDENTIFIER, "five"}, {token.ASSIGNMENT, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENTIFIER, "ten"}, {token.ASSIGNMENT, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENTIFIER, "add"}, {token.ASSIGNMENT, "="}, {token.FUNCTION, "fn"},
		{token.LEFT_PARENTHESIS, "("}, {token.IDENTIFIER, "x"}, {token.COMMA, ","}, {token.IDENTIFIER, "y"}, {token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RETURN, "return"}, {token.IDENTIFIER, "x"}, {token.PLUS, "+"}, {token.IDENTIFIER, "y"}, {token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENTIFIER, "result"}, {token.ASSIGNMENT, "="}, {token.IDENTIFIER, "add"},
		{token.LEFT_PARENTHESIS, "("}, {token.IDENTIFIER, "five"}, {token.COMMA, ","}, {token.IDENTIFIER, "ten"}, {token.RIGHT_PARENTHESIS, ")"}, {token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range tests {
		token := lexer.NextToken()
		if token.Type != expected.Type {
			t.Fatalf("i = %d - Expected type [%q], got [%q]", i, expected.Type, token.Type)
		}
		if token.Literal != expected.Literal {
			t.Fatalf("i = %d - Expected Literal [%q], got [%q]", i, expected.Literal, token.Literal)
		}
	}
}

func TestBasicConditions(t *testing.T) {
	input := `if (5 < 10) {
		return true;
	} else {
		return false;
	}

	return 10 == 10;
	return 5 != 10;
	`
	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.IF, "if"}, {token.LEFT_PARENTHESIS, "("}, {token.INT, "5"}, {token.LESS_THAN, "<"}, {token.INT, "10"}, {token.RIGHT_PARENTHESIS, ")"}, {token.LEFT_BRACE, "{"},
		{token.RETURN, "return"}, {token.TRUE, "true"}, {token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"}, {token.ELSE, "else"}, {token.LEFT_BRACE, "{"},
		{token.RETURN, "return"}, {token.FALSE, "false"}, {token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.RETURN, "return"}, {token.INT, "10"}, {token.EQUAL, "=="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.RETURN, "return"}, {token.INT, "5"}, {token.NOT_EQUAL, "!="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range tests {
		token := lexer.NextToken()
		if token.Type != expected.Type {
			t.Fatalf("i = %d - Expected type [%q], got [%q]", i, expected.Type, token.Type)
		}
		if token.Literal != expected.Literal {
			t.Fatalf("i = %d - Expected Literal [%q], got [%q]", i, expected.Literal, token.Literal)
		}
	}
}

func TestOperators(t *testing.T) {
	input := `let result = 5 + 10;
	!-/*5;
	5 < 10 > 5;
	`
	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.LET, "let"}, {token.IDENTIFIER, "result"}, {token.ASSIGNMENT, "="}, {token.INT, "5"}, {token.PLUS, "+"}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.BANG, "!"}, {token.MINUS, "-"}, {token.SLASH, "/"}, {token.ASTERISK, "*"}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.INT, "5"}, {token.LESS_THAN, "<"}, {token.INT, "10"}, {token.GREATER_THAN, ">"}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range tests {
		token := lexer.NextToken()
		if token.Type != expected.Type {
			t.Fatalf("i = %d - Expected type [%q], got [%q]", i, expected.Type, token.Type)
		}
		if token.Literal != expected.Literal {
			t.Fatalf("i = %d - Expected Literal [%q], got [%q]", i, expected.Literal, token.Literal)
		}
	}
}

