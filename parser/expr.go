package parser

import (
	ty "nao/types"
)

type (
	Expr interface {
		Pos() ty.Position
		exprNode()
	}

	Stmt interface {
		Pos() ty.Position
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

func (b BadExpr) Pos() ty.Position {
	return b.From
}

// Basic Literal
type BasicLit struct {
	pos   ty.Position
	tok   ty.TokenCode
	value string
}

func (b BasicLit) Pos() ty.Position {
	return b.pos
}

// Identifier
type Ident struct {
	name string
	pos  ty.Position
}

func (i Ident) Pos() ty.Position {
	return i.pos
}

// Unary Expression
type UnaryExpr struct {
	OP ty.Token
	A  Expr
}

func (u UnaryExpr) Pos() ty.Position {
	return u.OP.Pos
}

// Binary Expression
type BinaryExpr struct {
	A  Expr
	OP ty.Token
	B  Expr
}

func (b BinaryExpr) Pos() ty.Position {
	return b.A.Pos()
}

// Parentethised(?) Expression
type ParenExpr struct {
	A Expr
}

func (p ParenExpr) Pos() ty.Position {
	return p.A.Pos()
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

func (s SeqStmt) Pos() ty.Position {
	return s.pos
}

func (*SeqStmt) stmtNode() {}
