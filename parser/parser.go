package parser

import (
	"fmt"
	l "nao/lexer"
	t "nao/types"
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
	case t.NUMBER, t.FLOAT, t.IDENT, t.STRING:
		return p.binop()
	case t.LPAREN:
		p.read()
		x := p.parseExpr()
		if p.head.T != t.RPAREN {
			return &BadExpr{From: x.Pos(), value: "Unclosed parenthesis"}
		}
		return &ParenExpr{A: x}
	case t.EOF:
		return nil
	default:
		fmt.Printf("Hit default: %v\n", p.head)
	}
	return nil

}

func (p *Parser) unary() *UnaryExpr {
	op := p.head
	p.read()
	switch p.head.T {
	case t.NUMBER, t.FLOAT, t.IDENT, t.LPAREN:
		return &UnaryExpr{OP: op, A: p.parseExpr()}
	default:
		return nil
	}
}

func (p *Parser) binop() Expr {
	left := p.parseLhs()
	p.read()

	switch p.head.T {
	case t.PLUS, t.MINUS, t.MULTI, t.SLASH, t.LPAREN:
		bin := &BinaryExpr{A: left, OP: p.head}
		p.read()
		bin.B = p.parseExpr()
		return bin
	case t.EOF, t.SEMICOLON, t.RPAREN:
		return left
	}

	return &BadExpr{From: left.Pos(), value: p.head.S}

}

func (p *Parser) parseLhs() Expr {
	switch p.head.T {
	case t.IDENT:
		return p.parseIdent()
	case t.STRING, t.NUMBER, t.FLOAT:
		return &BasicLit{value: p.head.S, pos: p.head.Pos, tok: p.head.T}
	}
	return nil
}

func (p *Parser) parseIdent() *Ident {
	x := &Ident{name: p.head.S, pos: p.head.Pos}
	p.read()
	return x
}