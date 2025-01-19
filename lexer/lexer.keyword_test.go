package lexer

import (
	"testing"

	ty "github.com/amstrups/nao/types"
)

func TestFuncKeyword1(t *testing.T) {
	input := "func 2 + 2"

	expected := TList{
		ty.KEY_FUNCTION,
		ty.NUMBER,
		ty.PLUS,
		ty.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestTypeKeyword1(t *testing.T) {
	input := "int float string"

	expected := TList{
		ty.KEY_INT,
		ty.KEY_FLOAT,
		ty.KEY_STRING,
	}

	assertNoError(input, expected, t)
}

func TestTypeKeyword2(t *testing.T) {
	input := "intfloatstring"

	expected := TList{
		ty.KEY_INT,
		ty.KEY_FLOAT,
		ty.KEY_STRING,
	}

	assert(input, expected, ELEMENT_INEQUALITY, t)
}

func TestTypeKeyword3(t *testing.T) {
	input := "intfloat string"

	expected := TList{
		ty.IDENT,
		ty.KEY_STRING,
	}

	assertNoError(input, expected, t)
}
