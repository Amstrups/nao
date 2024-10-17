package types

import (
	e "errors"
)

var InvalidArgument = e.New("lists of unequal lengths")

type TokenCode int

const (
	EOF = iota
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
