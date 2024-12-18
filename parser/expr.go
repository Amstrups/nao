package parser

import (
	ty "github.com/amstrups/nao/types"
)

type (
	Node interface {
		Start() ty.Position
	}

	Expr interface {
		Node
		exprNode()
	}

	Stmt interface {
		Node
		stmtNode()
	}

	ProgramExpr struct {
		Root Stmt
	}
)

// Context
type Context struct {
	Program ProgramExpr
}

// Expressions

// Bad Expression
type BadExpr struct {
	From  ty.Position
	value string
	//From, To ty.Position // Not using "To" yet
}

func (b BadExpr) Start() ty.Position {
	return b.From
}

// Basic Literal
type BasicLit struct {
	Pos   ty.Position
	Tok   ty.TokenCode
	Value string
}

func (b BasicLit) Start() ty.Position {
	return b.Pos
}

// Identifier
type Ident struct {
	Name string
	Pos  ty.Position
}

func (i Ident) Start() ty.Position {
	return i.Pos
}

// Unary Expression
type UnaryExpr struct {
	OP ty.Token
	A  Expr
}

func (u UnaryExpr) Start() ty.Position {
	return u.OP.Pos
}

// Binary Expression
type BinaryExpr struct {
	A  Expr
	OP ty.Token
	B  Expr
}

func (b BinaryExpr) Start() ty.Position {
	return b.A.Start()
}

// Parentethised(?) Expression
type ParenExpr struct {
	A Expr
}

func (p ParenExpr) Start() ty.Position {
	return p.A.Start()
}

func (*BadExpr) exprNode()    {}
func (*Ident) exprNode()      {}
func (*BasicLit) exprNode()   {}
func (*ParenExpr) exprNode()  {}
func (*UnaryExpr) exprNode()  {}
func (*BinaryExpr) exprNode() {}

// Statements

// Sequence Expression
type SeqStmt struct {
	pos ty.Position
	X   []Expr
}

func (s SeqStmt) Start() ty.Position {
	return s.pos
}

func (*SeqStmt) stmtNode() {}
