package types

import "testing"

func defaultToken(t TokenCode) Token {
	return Token{
		T:   t,
		Pos: Default(),
		S:   "",
	}

}

func TestEqualTokens(t *testing.T) {
	// arrange
	tokens := []Token{
		defaultToken(LPAREN),
		defaultToken(NUMBER),
		defaultToken(PLUS),
		defaultToken(NUMBER),
		defaultToken(RPAREN),
	}

	// assume
	codes := []TokenCode{LPAREN, NUMBER, PLUS, NUMBER, RPAREN}

	// assert
	if IsEqual(tokens, codes) == false {
		t.Fail()
	}
}

func TestNonEqualTokens(t *testing.T) {
	// arrange
	tokens := []Token{
		defaultToken(LPAREN),
		defaultToken(NUMBER),
		defaultToken(PLUS),
		defaultToken(NUMBER),
		defaultToken(RPAREN),
	}

	// assume
	codes := []TokenCode{LPAREN, NUMBER, NUMBER, PLUS, RPAREN}

	// assert
	if IsEqual(tokens, codes) == true {
		t.Fail()
	}
}
