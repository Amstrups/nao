package semantics

import (
	"fmt"

	"github.com/amstrups/nao/parser"
	ty "github.com/amstrups/nao/types"
)

var TtT map[ty.TokenCode]T = map[ty.TokenCode]T{
	ty.NUMBER: INT,
	ty.FLOAT:  FLOAT,
	ty.BINARY: INT,
	ty.STRING: STRING,
}

type Semantics struct {
	*parser.Parser
	Root Stmt
}

func New(p *parser.Parser) *Semantics {
	return &Semantics{Parser: p}
}

func (s *Semantics) Eval() Stmt {
	s.Root = s.evalStmt(s.Context.Program.Root)
	return s.Root
}

func (s *Semantics) evalStmt(stmt parser.Stmt) Stmt {
	switch stm := stmt.(type) {
	case *parser.SeqStmt:
		seq := &SeqStmt{Pos: stm.Start(), X: make([]Expr, len(stm.X))}
		for i, exp := range stm.X {
			e, t := s.evalExpr(exp)
			seq.X[i] = e
			seq.T = t
		}
		return seq
	default:
		panic("Only SeqStmt supported at current date")
	}

	return nil
}

func (s *Semantics) evalExpr(exp parser.Expr) (Expr, T) {
	switch expT := exp.(type) {
	case *parser.UnaryExpr:
		e, t := s.evalExpr(expT.A)
		return &UnaryExpr{OP: expT.OP, A: e, T: t}, t
	case *parser.BasicLit:
		t, ok := TtT[expT.Tok]
		if ok {
			return Basic(expT, t), t
		}
		return Basic(expT, ILLEGAL), ILLEGAL
	case *parser.BinaryExpr:
		e1, t1 := s.evalExpr(expT.A)
		e2, t2 := s.evalExpr(expT.B)
		if t1 != t2 {
			panic(fmt.Sprintf("binary op between %d and %d not allowed", t1, t2))
		}
		return &BinaryExpr{A: e1, B: e2, OP: expT.OP, T: t1}, t1
	case *parser.ParenExpr:
		e, t := s.evalExpr(expT.A)
		return &ParenExpr{A: e, T: t}, t
	case *parser.Ident:
		panic("Semantic of IDENT nyi")
	default:
		panic("hit default in semantic analysis")
	}
}
