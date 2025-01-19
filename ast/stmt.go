package ast

import (
	"fmt"

	ty "github.com/amstrups/nao/types"
)

type (
	Stmt interface {
		Node
		stmtNode()
	}
)

// Kallax Statement
type KallaxStmt struct {
	Pos   ty.Position
	Main  Stmt
	Decls []Decl
}

func (k KallaxStmt) Start() ty.Position {
	return ty.Default()
}

// Sequence Statement
type SeqStmt struct {
	Pos ty.Position
	X   []Stmt
}

func (s SeqStmt) Start() ty.Position {
	return s.Pos
}

// Expression Statement
type ExprStmt struct {
	A Expr
	ty.T
}

func (e ExprStmt) Start() ty.Position {
	return e.A.Start()
}

func (e ExprStmt) String() string {
	return fmt.Sprintf("%v", e.A)

}

// Assignment Statement
type AssignStmt struct {
	I       *Ident
	DeclTyp *ty.Token
	A       Expr
}

func (a AssignStmt) Start() ty.Position {
	return a.I.Start()
}

// Stmt implementations
func (*KallaxStmt) stmtNode() {}
func (*SeqStmt) stmtNode()    {}
func (*ExprStmt) stmtNode()   {}
func (*AssignStmt) stmtNode() {}

func BadStmt(pos ty.Position, value string) *ExprStmt {
	return &ExprStmt{
		A: &BadExpr{pos, value},
	}
}
