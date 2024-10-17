package main

import (
	"fmt"
	"nao/lexer"
	"nao/parser"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	filename := "./examples/Seq/ex_seq3"
	filename += ".nao"
	fmt.Printf("Reading file: \"%s\"\n", filename)
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(f))
	p := parser.New(l)

	if err := p.ParseExpr(); err != nil {
		fmt.Println(err)
	}

	out, err := os.Create(filename + ".out")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Writing dump to file: \"%s\"\n", filename + ".out")
	spew.Fdump(out, p.Context)
}
