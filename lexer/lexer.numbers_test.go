package lexer

import (
	"fmt"
	"testing"

	ty "github.com/amstrups/nao/types"
)

func TestNumbers(t *testing.T) {
	for i := 0; i < 20; i++ {
		input := fmt.Sprint(i)
		assertNoError(input, TList{ty.NUMBER}, t)
	}
}

func TestFloats(t *testing.T) {
	assertNoError("0.1", TList{ty.FLOAT}, t)
	assertNoError("0.1", TList{ty.FLOAT}, t)
	assertNoError("1.1", TList{ty.FLOAT}, t)
	assertNoError("11.1", TList{ty.FLOAT}, t)
	assertNoError("1.11", TList{ty.FLOAT}, t)
	assertNoError("11.11", TList{ty.FLOAT}, t)
}

func TestBinaries(t *testing.T) {
	assertNoError("0b0", TList{ty.BINARY}, t)
	assertNoError("0b1", TList{ty.BINARY}, t)
	assertNoError("0b01", TList{ty.BINARY}, t)
}
