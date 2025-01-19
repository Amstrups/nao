package ast

import (
	"fmt"

	"github.com/amstrups/nao/types"
)

type (
	Decl interface {
		Node
		declNode()
	}
)

// Type Declaration
type TypeDecl struct {
	Pos  types.Position
	Name string
	Type *Ident
	types.T
}

// Struct Declaration
type StructDecl struct {
	Pos        types.Position
	Name       string
	Parameters []*TypedVariables
}

func (s StructDecl) Start() types.Position {
	return s.Pos
}

type TypedVariables struct {
	Name *Ident
	Type *Ident
}

func (v TypedVariables) String() string {
	return fmt.Sprintf("[%v:%v]", v.Name, v.Type)
}

// Implementation Declaration
type ImplDecl struct {
	Pos        types.Position
	Name       string
	Alias      *Ident
	Parameters []*FuncDecl
	types.T
}

// Function Declaration
type FuncDecl struct {
	Pos        types.Position
	Name       string
	Parameters []*TypedVariables
	Return     *Ident
	Body       Stmt
}

func (f FuncDecl) Start() types.Position {
	return f.Pos
}

func (*TypeDecl) declNode()   {}
func (*StructDecl) declNode() {}
func (*ImplDecl) declNode()   {}
func (*FuncDecl) declNode()   {}
