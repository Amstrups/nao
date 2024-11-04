package types

//go:generate stringer -type=TokenCode
type TokenCode uint

const (
	EOF TokenCode = iota
	ILLEGAL

	IDENT
	STRING

	NUMBER
	FLOAT

	DOT
	SEMICOLON
	DOUBLEQUOTE
	SINGLEQUOTE

	LPAREN
	RPAREN

	EQ
	PLUS
	MINUS
	MULTI

	BACKSLASH
	SLASH

	BINARY
	P2

	// Keywords
	FUNC
)

type Token struct {
	T   TokenCode
	Pos Position
	S   string
}

func IsEqual(value []Token, expected []TokenCode) bool {
	if len(value) != len(expected) {
		return false
	}
	for i := range value {
		if value[i].T != expected[i] {
			return false
		}

	}
	return true
}

func CheckIfKeyword(str string) TokenCode {
	switch str {
	case "func":
		return FUNC
	default:
		return IDENT
	}
}
