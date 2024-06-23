package lexer

import "igorracki/sledzscript/token"

type Lexer struct {
	input                    string
	position                 int
	currentCharacter         byte
	currentCharacterPosition int
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readNextCharacter()
	return lexer
}

func (lexer *Lexer) readNextCharacter() {
	if lexer.currentCharacterPosition >= len(lexer.input) {
		lexer.currentCharacter = 0
	} else {
		lexer.currentCharacter = lexer.input[lexer.currentCharacterPosition]
	}
	lexer.position = lexer.currentCharacterPosition
	lexer.currentCharacterPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var matchedToken token.Token
	lexer.ignoreWhitespace()
	switch lexer.currentCharacter {
	case '=':
		if lexer.peekNextCharacter() == '=' {
			matchedEqualsCharacter := lexer.currentCharacter
			lexer.readNextCharacter()
			literal := string(matchedEqualsCharacter) + string(lexer.currentCharacter)
			matchedToken = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			matchedToken = token.Token{Type: token.ASSIGNMENT, Literal: string(lexer.currentCharacter)}
		}
	case ';':
		matchedToken = token.Token{Type: token.SEMICOLON, Literal: string(lexer.currentCharacter)}
	case '(':
		matchedToken = token.Token{Type: token.LEFT_PARENTHESIS, Literal: string(lexer.currentCharacter)}
	case ')':
		matchedToken = token.Token{Type: token.RIGHT_PARENTHESIS, Literal: string(lexer.currentCharacter)}
	case ',':
		matchedToken = token.Token{Type: token.COMMA, Literal: string(lexer.currentCharacter)}
	case '+':
		matchedToken = token.Token{Type: token.PLUS, Literal: string(lexer.currentCharacter)}
	case '{':
		matchedToken = token.Token{Type: token.LEFT_BRACE, Literal: string(lexer.currentCharacter)}
	case '}':
		matchedToken = token.Token{Type: token.RIGHT_BRACE, Literal: string(lexer.currentCharacter)}
	case '-':
		matchedToken = token.Token{Type: token.MINUS, Literal: string(lexer.currentCharacter)}
	case '!':
		if lexer.peekNextCharacter() == '=' {
			matchedBangCharacter := lexer.currentCharacter
			lexer.readNextCharacter()
			literal := string(matchedBangCharacter) + string(lexer.currentCharacter)
			matchedToken = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			matchedToken = token.Token{Type: token.BANG, Literal: string(lexer.currentCharacter)}
		}
	case '/':
		matchedToken = token.Token{Type: token.SLASH, Literal: string(lexer.currentCharacter)}
	case '*':
		matchedToken = token.Token{Type: token.ASTERISK, Literal: string(lexer.currentCharacter)}
	case '<':
		matchedToken = token.Token{Type: token.LESS_THAN, Literal: string(lexer.currentCharacter)}
	case '>':
		matchedToken = token.Token{Type: token.GREATER_THAN, Literal: string(lexer.currentCharacter)}
	case 0:
		matchedToken = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(lexer.currentCharacter) {
			identifier := lexer.readIdentifier()
			tokenType := token.LookupIdentifier(identifier)
			matchedToken = token.Token{Type: tokenType, Literal: identifier}
			return matchedToken
		} else if isDigit(lexer.currentCharacter) {
			matchedToken = token.Token{Type: token.INT, Literal: lexer.readNumber()}
			return matchedToken
		} else {
			matchedToken = token.Token{Type: token.ILLEGAL, Literal: string(lexer.currentCharacter)}
		}
	}
	lexer.readNextCharacter()
	return matchedToken
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.currentCharacter) {
		lexer.readNextCharacter()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func (lexer *Lexer) ignoreWhitespace() {
	for lexer.currentCharacter == ' ' || lexer.currentCharacter == '\t' || lexer.currentCharacter == '\r' || lexer.currentCharacter == '\n' {
		lexer.readNextCharacter()
	}
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.currentCharacter) {
		lexer.readNextCharacter()
	}
	return lexer.input[position:lexer.position]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) peekNextCharacter() byte {
	if lexer.currentCharacterPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.currentCharacterPosition]
	}
}
