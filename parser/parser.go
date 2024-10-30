package parser

import (
	"fmt"

	l "github.com/amstrups/nao/lexer"
	t "github.com/amstrups/nao/types"
)

type Parser struct {
	*l.Lexer
	Context
	head t.Token
}

func (p *Parser) read() {
	p.head = p.Lex()
}

func New(lexer *l.Lexer) *Parser {
	return &Parser{
		Lexer:   lexer,
		Context: Context{},
		head:    lexer.Lex(),
	}
}

func (p *Parser) ParseExpr() error {
	top := ProgramExpr{}
	p.Program = top

	p.Program.Root = p.parseStmt()

	return nil
}

func (p *Parser) parseStmt() Stmt {
	seq := &SeqStmt{pos: p.head.Pos, X: []Expr{}}
	for {
		switch p.head.T {
		case t.EOF:
			return seq
		case t.SEMICOLON:
			p.read()
			xs := p.parseExpr()
			seq.X = append(seq.X, xs)
		default:
			seq.X = append(seq.X, p.parseExpr())
		}
		p.read()
	}
}

func (p *Parser) parseExpr() Expr {
	switch p.head.T {
	case t.MINUS:
		return p.unary()
	case t.NUMBER, t.FLOAT, t.BINARY, t.IDENT, t.STRING:

		return p.binop()
	case t.LPAREN:
		p.read()
		x := p.parseExpr()
		if p.head.T != t.RPAREN {
			return &BadExpr{From: x.Start(), value: "Unclosed parenthesis"}
		}
		return &ParenExpr{A: x}
	case t.EOF:
		return nil
	default:
		fmt.Printf("Hit default: %v\n", p.head)
	}
	return nil

}

func (p *Parser) unary() Expr {
	op := p.head
	p.read()
	switch p.head.T {
	case t.NUMBER, t.FLOAT, t.IDENT, t.LPAREN:
		a := p.parseExpr()
		switch at := a.(type) {
		case *BinaryExpr:
			at.A = &UnaryExpr{OP: op, A: at.A}
			return at
		default:
			return &UnaryExpr{OP: op, A: a}
		}
	default:
		return &BadExpr{From: p.head.Pos, value: fmt.Sprintf("Bad unary: %s", p.head.S)}
	}
}

func (p *Parser) binop() Expr {
	left := p.parseLhs()
	p.read()

	switch p.head.T {
	case t.MULTI, t.SLASH, t.STRING:
		bin := &BinaryExpr{A: left, OP: p.head}
		p.read()

		B := p.parseExpr()
		switch ty := B.(type) {
		case *BinaryExpr:
			bin.B = ty.A
			bi := *bin
			ty.A = &bi
			return ty
		default:
			bin.B = B
		}
		return bin
	case t.PLUS, t.MINUS, t.LPAREN:
		bin := &BinaryExpr{A: left, OP: p.head}
		p.read()
		bin.B = p.parseExpr()
		return bin
	case t.EOF, t.SEMICOLON, t.RPAREN:
		return left
	}

	return &BadExpr{From: left.Start(), value: fmt.Sprintf("Bad binop: %s", p.head.S)}

}

func (p *Parser) parseLhs() Expr {
	switch p.head.T {
	case t.IDENT:
		return p.parseIdent()
	case t.STRING, t.NUMBER, t.FLOAT, t.BINARY:
		return &BasicLit{Value: p.head.S, Pos: p.head.Pos, Tok: p.head.T}
	}
	return nil
}

func (p *Parser) parseIdent() *Ident {
	x := &Ident{Name: p.head.S, Pos: p.head.Pos}
	p.read()
	return x
}
