package run

import (
	"fmt"
	"github.com/amstrups/nao/ast"
	"github.com/amstrups/nao/lexer"
	"github.com/amstrups/nao/parser"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func Run(input string) ast.Stmt {
	l := lexer.New(input)
	p := parser.New(l)

	if err := p.ParseFile(); err != nil {
		fmt.Println(err)
		return nil
	}

  fmt.Println("\nAST:")
	spew.Dump(p.Root)
	return p.Root
}

func RunFile(relativePath string) ast.Stmt {
	fmt.Printf("Reading file: \"%s\"\n", relativePath)
	f, err := os.ReadFile(relativePath)
	if err != nil {
		panic(err)
	}

	return Run(string(f))
}

func RunDump(relativePath string) {
	context := RunFile(relativePath)

	outPath := relativePath + ".out"

	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Writing dump to file: \"%s\"\n", outPath)
	spew.Fdump(out, context)
}
