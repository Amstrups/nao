package ast

import (
	"fmt"

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
)

// Bad Expression
type BadExpr struct {
	From  ty.Position
	Value string
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
	ty.T
}

func (b BasicLit) Start() ty.Position {
	return b.Pos
}

func (b BasicLit) String() string {
	return b.Tok.String() + ": " + b.Value

}

// Identifier
type Ident struct {
	Name string
	Pos  ty.Position
	ty.T
}

func (i Ident) Start() ty.Position {
	return i.Pos
}

func (i Ident) String() string {
	return "Ident: '" + i.Name + "'"
}

// Unary Expression
type UnaryExpr struct {
	OP ty.Token
	A  Expr
	ty.T
}

func (u UnaryExpr) Start() ty.Position {
	return u.OP.Pos
}

// Binary Expression
type BinaryExpr struct {
	A  Expr
	OP ty.Token
	B  Expr
	ty.T
}

func (b BinaryExpr) Start() ty.Position {
	return b.A.Start()
}

func (b BinaryExpr) String() string {
	return fmt.Sprintf("BinOp: [%v;%s;%v]", b.A, b.OP.T.String(), b.B)
}

// Not-A-Vector Expression
type ArrayExpr struct {
	Pos  ty.Position
	Xt   *Ident
	XDim []Expr
	Xs   []Expr
	ty.T
}

func (a ArrayExpr) Start() ty.Position {
	return a.Pos
}

func (a ArrayExpr) String() string {
	return fmt.Sprintf("Dim(%v): %v", a.XDim, a.Xs)
}

// Parentethised(?) Expression
type ParenExpr struct {
	A Expr
	ty.T
}

func (p ParenExpr) Start() ty.Position {
	return p.A.Start()
}

// Call Expression
type CallExpr struct {
	Name string
	Pos  ty.Position
	Args []Expr
	ty.T
}

func (c CallExpr) Start() ty.Position {
	return c.Pos
}

func (c CallExpr) String() string {
	return "Call of: " + c.Name +
		fmt.Sprintf("%v", c.Args)
}

// Constructor Expression
type ConstructExpr struct {
	Name string
	Pos  ty.Position
	Args []NamedArgument
	ty.T
}

func (c ConstructExpr) Start() ty.Position {
	return c.Pos
}

func (c ConstructExpr) String() string {
	return "Constructor of: " + c.Name +
		fmt.Sprintf("%v", c.Args)
}

func (c *ConstructExpr) AddArgument(variable *Ident, value Expr) {
	c.Args = append(c.Args,
		NamedArgument{
			Var: variable,
			Val: value,
		})
}

// Named argument
type NamedArgument struct {
	Var *Ident
	Val Expr
}

func (n NamedArgument) String() string {
	return fmt.Sprintf("[%v,%v]", n.Var, n.Val)

}

// Expr implementations
func (*BadExpr) exprNode()       {}
func (*Ident) exprNode()         {}
func (*BasicLit) exprNode()      {}
func (*ParenExpr) exprNode()     {}
func (*CallExpr) exprNode()      {}
func (*ConstructExpr) exprNode() {}
func (*UnaryExpr) exprNode()     {}
func (*BinaryExpr) exprNode()    {}
func (*ArrayExpr) exprNode()     {}
