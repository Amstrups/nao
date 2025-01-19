package parser

import (
	"fmt"

	"github.com/amstrups/nao/ast"
	ty "github.com/amstrups/nao/types"
	"github.com/davecgh/go-spew/spew"
)

type Ass struct {
	K string
	V uint
}

type Semantics struct {
	*Parser
	Root  ast.Stmt
	Store []Ass
}

func NewSem(p *Parser) *Semantics {
	return &Semantics{Parser: p}
}

func (s *Semantics) Eval() ast.Stmt {
	s.Root = s.evalStmt(s.Program.Root)
	return s.Root
}

func (s *Semantics) evalStmt(stmt ast.Stmt) ast.Stmt {
	switch stm := stmt.(type) {
	case *ast.ExprStmt, *ast.AssignStmt:
		return s.evalSimpleStmt(stm)
	case *ast.SeqStmt:
		seq := &ast.SeqStmt{Pos: stm.Start(), X: make([]ast.Stmt, len(stm.X))}
		for i, exp := range stm.X {
			seq.X[i] = s.evalSimpleStmt(exp)
		}
		return seq
	default:
		panic("Semantics: evalStmt")
	}
}

func (s *Semantics) evalSimpleStmt(stmt ast.Stmt) ast.Stmt {
	switch st := stmt.(type) {
	case *ast.ExprStmt:
		e, t := s.evalExpr(st.A)
		return &ast.ExprStmt{A: e, T: t}
	case *ast.AssignStmt:
		e, t := s.evalExpr(st.A)
		return &ast.ExprStmt{A: e, T: t}
	default:
		panic("evalSimpleStmt")
	}
}

func (s *Semantics) evalExpr(exp ast.Expr) (ast.Expr, ty.T) {
	switch expT := exp.(type) {
	case *ast.UnaryExpr:
		e, t := s.evalExpr(expT.A)
		return &ast.UnaryExpr{OP: expT.OP, A: e, T: t}, t
	case *ast.BasicLit:
		t, ok := ty.TtT[expT.Tok]
		if ok {
			expT.T = t
			return expT, t
		}
		expT.T = ty.T_ILLEGAL
		return expT, ty.T_ILLEGAL
	case *ast.BinaryExpr:
		e1, t1 := s.evalExpr(expT.A)
		e2, t2 := s.evalExpr(expT.B)
		if t1 != t2 {
			panic(fmt.Sprintf("binary op between %s and %s not allowed", t1, t2))
		}
		return &ast.BinaryExpr{A: e1, B: e2, OP: expT.OP, T: t1}, t1
	case *ast.ParenExpr:
		e, t := s.evalExpr(expT.A)
		return &ast.ParenExpr{A: e, T: t}, t
	default:
		spew.Dump(exp)
		panic("hit default in semantic analysis")
	}
}

/*
func (s *Semantics) evalIdent(i ast.Ident) (ast.Ident, ty.T) {


}
*/
