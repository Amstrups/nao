package semantics

import (
	"github.com/amstrups/nao/parser"
	ty "github.com/amstrups/nao/types"
)

type (
	T uint

	Expr interface {
		parser.Node
		exprNode()
	}

	Stmt interface {
		parser.Node
		stmtNode()
	}

	ProgramExpr struct {
		Root Stmt
	}
)

const (
	ILLEGAL = iota

	INT
	FLOAT

	STRING

	BOOL

	VOID

	IDENT
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
	T
}

func Basic(b *parser.BasicLit, t T) *BasicLit {
	return &BasicLit{
		Pos:   b.Pos,
		Tok:   b.Tok,
		Value: b.Value,
		T:     t,
	}

}

func (b BasicLit) Start() ty.Position {
	return b.Pos
}

// Identifier
type Ident struct {
	name string
	pos  ty.Position
	T
}

func (i Ident) Start() ty.Position {
	return i.pos
}

// Unary Expression
type UnaryExpr struct {
	OP ty.Token
	A  Expr
	T
}

func (u UnaryExpr) Start() ty.Position {
	return u.OP.Pos
}

// Binary Expression
type BinaryExpr struct {
	A  Expr
	OP ty.Token
	B  Expr
	T
}

func (b BinaryExpr) Start() ty.Position {
	return b.A.Start()
}

// Parentethised(?) Expression
type ParenExpr struct {
	A Expr
	T
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
	Pos ty.Position
	X   []Expr
	T
}

func (s SeqStmt) Start() ty.Position {
	return s.Pos
}

func (*SeqStmt) stmtNode() {}