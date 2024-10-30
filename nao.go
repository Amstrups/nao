package nao

import (
	"fmt"
	"github.com/amstrups/nao/lexer"
	"github.com/amstrups/nao/parser"
	"github.com/amstrups/nao/semantics"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func Run(input string) semantics.Stmt {
	l := lexer.New(input)
	p := parser.New(l)

	if err := p.ParseExpr(); err != nil {
		fmt.Println(err)
		return nil
	}

	s := semantics.New(p)

	s_ := s.Eval()
	spew.Dump(s)
	return s_
}

func RunFile(path string) semantics.Stmt {
	fmt.Printf("Reading file: \"%s\"\n", path)
	f, err := os.ReadFile(path)
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
