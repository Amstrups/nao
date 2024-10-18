package lexer

import (
	"fmt"
	ty "github.com/amstrups/nao/types"
	"strings"
	"testing"
)

type ASSERT_ERROR_CODE int

const (
	NO_ERROR = iota
	ELEMENT_INEQUALITY
	LENGTH_INEQAULITY
)

type AssertError struct {
	code ASSERT_ERROR_CODE
	msg  string
}

func toString[T ty.Token](ts []T, fn func(T) string) []string {
	result := make([]string, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func concat(xs []ty.Token) string {
	ss := toString(xs, func(t ty.Token) string { return t.S })

	return "[" + strings.Join(ss, ",") + "]"
}

func concat2(xs []ty.TokenCode) string {
	ss := make([]string, len(xs))
	for i, x := range xs {
		ss[i] = fmt.Sprintf("%d", x)
	}

	return "[" + strings.Join(ss, ",") + "]"
}

func prepareMessage(ts []ty.Token, tcs []ty.TokenCode) string {

	return fmt.Sprintf("\n%10s %s\n%10s %s", "Found:", concat(ts), "Expected:", concat2(tcs))
}

func lexAndAssertEquality(input string, expected []ty.TokenCode) AssertError {

	lexer := New(input)

	i := 0

	lexed := make([]ty.Token, len(expected))

	for {
		tok := lexer.Lex()
		if tok.T == ty.EOF {
			if i == len(expected) {

				return AssertError{NO_ERROR, ""}
			}
			return AssertError{LENGTH_INEQAULITY, prepareMessage(lexed, expected)}
		}

		if i >= len(expected) {
			return AssertError{LENGTH_INEQAULITY, prepareMessage(lexed, expected)}
		}

		lexed[i] = tok

		if expected[i] != tok.T {
			return AssertError{ELEMENT_INEQUALITY, prepareMessage(lexed, expected)}
		}

		i++
	}

}
func assert(input string, expected []ty.TokenCode, expRes ASSERT_ERROR_CODE, t *testing.T) {
	result := lexAndAssertEquality(input, expected)

	if result.code != expRes {
		t.Fatal(result.msg)
	}
}

func assertNoError(input string, expected []ty.TokenCode, t *testing.T) {
	assert(input, expected, NO_ERROR, t)
}

func TestSymbols(t *testing.T) {
	assertNoError(".", []ty.TokenCode{ty.DOT}, t)
	assertNoError("\"", []ty.TokenCode{ty.DOUBLEQUOTE}, t)
	assertNoError("'", []ty.TokenCode{ty.SINGLEQUOTE}, t)

	assertNoError("(", []ty.TokenCode{ty.LPAREN}, t)
	assertNoError(")", []ty.TokenCode{ty.RPAREN}, t)

	assertNoError("=", []ty.TokenCode{ty.EQ}, t)
	assertNoError("+", []ty.TokenCode{ty.PLUS}, t)
	assertNoError("-", []ty.TokenCode{ty.MINUS}, t)
	assertNoError("*", []ty.TokenCode{ty.MULTI}, t)

	assertNoError("\\", []ty.TokenCode{ty.BACKSLASH}, t)
	assertNoError("/", []ty.TokenCode{ty.SLASH}, t)
}

func TestIdent1(t *testing.T) {
	input := "Foo"

	expected := []ty.TokenCode{
		ty.IDENT,
	}

	assertNoError(input, expected, t)
}

func TestIdent2(t *testing.T) {
	input := "Foo2"

	expected := []ty.TokenCode{
		ty.IDENT,
	}

	assertNoError(input, expected, t)
}

func TestIdent3(t *testing.T) {
	input := "Foo+"

	expected := []ty.TokenCode{
		ty.IDENT,
		ty.PLUS,
	}

	assertNoError(input, expected, t)
}

func TestSimpleBinaryExpression1(t *testing.T) {
	input := "2+2"

	expected := []ty.TokenCode{
		ty.NUMBER,
		ty.PLUS,
		ty.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSimpleBinaryExpression2(t *testing.T) {
	input := "22+2"

	expected := []ty.TokenCode{
		ty.NUMBER,
		ty.PLUS,
		ty.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSimpleBinaryExpression3(t *testing.T) {
	input := "22+111111"

	expected := []ty.TokenCode{
		ty.NUMBER,
		ty.PLUS,
		ty.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSpacedTokens1(t *testing.T) {
	input := "2 2"

	expected := []ty.TokenCode{
		ty.NUMBER,
		ty.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSpacedTokens2(t *testing.T) {
	input := "Foo Baa"

	expected := []ty.TokenCode{
		ty.IDENT,
		ty.IDENT,
	}

	assertNoError(input, expected, t)
}

func TestPlusSequence(t *testing.T) {
	input := "++"

	expected := []ty.TokenCode{
		ty.PLUS,
		ty.PLUS,
	}

	assertNoError(input, expected, t)
}

func TestEmptyInput(t *testing.T) {
	input := ""

	expected := []ty.TokenCode{
		ty.NUMBER,
	}

	assert(input, expected, LENGTH_INEQAULITY, t)
}

func TestEmptyExpectation(t *testing.T) {
	input := "input"

	expected := []ty.TokenCode{}

	assert(input, expected, LENGTH_INEQAULITY, t)
}
