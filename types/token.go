package types

//go:generate stringer -type=TokenCode
type TokenCode uint

const (
	EOF TokenCode = iota
	ILLEGAL
	NIL
  NEWLINE

	IDENT
	STRING

	NUMBER
	FLOAT

	DOT
  COMMA
	COLON
	SEMICOLON
	SINGLEQUOTE

  // (
	LPAREN
  // )
	RPAREN
  // [
	LBRACKET
  // ]
	RBRACKET
  // {
	LBRACE
  // }
	RBRACE

	EQ
	PLUS
	MINUS
	MULTI

	BACKSLASH
	SLASH
	DOUBLESLASH

	// Comment style like Nim
	COMMENT       // #
	COMMENT_BLOCK // #[ ... ]#

	BINARY

	// Keywords
	KEY_MAIN

	KEY_FUNCTION
	KEY_LET
	KEY_STRUCT
	KEY_IMPL
	KEY_RANGE

	KEY_PANIC
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

var KEYWORDS = map[string]TokenCode{
	"master": KEY_MAIN, // it's the same thing
	"main":   KEY_MAIN,

	"fn":     KEY_FUNCTION,
	"let":    KEY_LET,
	"struct": KEY_STRUCT,
	"impl":   KEY_IMPL,
	"range":  KEY_RANGE,

	"fuck": KEY_PANIC, // cause it's fun
  
}

func CheckIfKeyword(str string) TokenCode {
	val, ok := KEYWORDS[str]
	if ok {
		return val
	}

	return IDENT
}
