// My comments goes here for my package
package parser

import (
	"errors"
	"fmt"

	"github.com/amstrups/nao/ast"
	l "github.com/amstrups/nao/lexer"
	t "github.com/amstrups/nao/types"
	"github.com/davecgh/go-spew/spew"
)

type Parser struct {
	*l.Lexer
	ast.Program
	head t.Token
}

func (p *Parser) read() t.Token {
	p.head = p.Lex()
	for p.head.T == t.COMMENT || p.head.T == t.COMMENT_BLOCK {
		p.head = p.Lex()
	}
	return p.head
}

func New(lexer *l.Lexer) *Parser {
	return &Parser{
		Lexer:   lexer,
		Program: ast.Program{},
		head:    lexer.Lex(),
	}
}

func (p *Parser) ParseExpr() error {
	p.Program.Root = p.parseStmt()

	return nil
}

func (p *Parser) ParseFile() error {
	p.Program.Root = p.parseMain()

	return nil
}

// Reads until given tokencode is not equal to head
func (p *Parser) consume(tc t.TokenCode) {
	for p.head.T == tc {
		p.read()
	}
}

func (p *Parser) expect(toks ...t.TokenCode) bool {
	for i := 0; i < len(toks); i++ {
		p.read()
		if p.head.T != toks[i] {
			return false
		}
	}
	return true
}

func (p *Parser) try(tc t.TokenCode) (tok t.Token, err error) {
	tok = p.read()
	if tc != p.head.T {
		msg := fmt.Sprintf("[%v:%v]: expected %s, found %s:%s", tok.Pos.Line, tok.Pos.Column, tc, tok.T, tok.S)
		err = errors.New(msg)
	}
	return
}

func (p *Parser) tryAndRead(tc t.TokenCode) (tok t.Token, err error) {
	tok = p.read()
	if tc != p.head.T {
		msg := fmt.Sprintf("[%v:%v]: expected %s, found %s:%s", tok.Pos.Line, tok.Pos.Column, tc, tok.T, tok.S)
		err = errors.New(msg)
	}
	p.read()
	return
}

func (p *Parser) tryPanic(tc t.TokenCode) t.Token {
	tok := p.read()
	if tc != p.head.T {
		msg := fmt.Sprintf("expected %s, found %s:%s", tc, tok.T, tok.S)
		panic(msg)
	}
	return tok
}

func (p *Parser) nextIs(tc t.TokenCode) (tok t.Token, done bool) {
	tok = p.read()
	if tc != p.head.T {
		return
	}
	done = true
	return
}

func (p *Parser) error(pos t.Position, found t.TokenCode, expected ...t.TokenCode) error {
	msg := fmt.Sprintf("[%8s] found %s, expected token in [%v]", pos, found, expected)
	return errors.New(msg)
}

func (p *Parser) error2(found t.Token, expected ...t.TokenCode) error {
	msg := fmt.Sprintf("[%8s] found %s, expected token in [%v]", found.Pos, found.T, expected)
	return errors.New(msg)
}

/*
Returns new Ident expression
Will panic if input token is not identifier
Only use when token MUST be identifer
*/
func (p *Parser) newIdent(tok t.Token) *ast.Ident {
	if tok.T != t.IDENT {
		p.error2(tok, t.IDENT)
	}

	return &ast.Ident{
		Name: tok.S,
		Pos:  tok.Pos,
	}
}

func (p *Parser) parseMain() ast.Stmt {
	switch p.head.T {
	case t.EOF:
		return nil
	case t.KEY_MAIN:
		pos := p.head.Pos
		if !p.expect(t.EQ, t.LBRACE) {
			return &ast.SeqStmt{
				Pos: pos,
				X:   []ast.Stmt{ast.BadStmt(pos, "expect '= {' after main/master")},
			}
		}
		p.read()
		p.consume(t.NEWLINE)
		main := &ast.KallaxStmt{
			Pos: pos,
		}
		main.Main = p.parseSeqStmt()
		p.consume(t.RBRACE)
		p.consume(t.NEWLINE)
		main.Decls = p.parseDecls()
		spew.Dump(main.Decls[0])

		return main
	default:
		spew.Dump(p.head)
		panic("hit default in parseMain")
	}
}

func (p *Parser) parseStmt() ast.Stmt {
	return p.parseSeqStmt()
}

func (p *Parser) parseDecls() []ast.Decl {
	decls := []ast.Decl{}

	for {
		p.consume(t.NEWLINE)
		switch p.head.T {
		case t.EOF:
			return decls
		case t.KEY_STRUCT:
			fmt.Println("parsing struct")
			strct := p.parseStruct()
			fmt.Println("Struct output")
			spew.Dump(strct.Parameters)
			decls = append(decls, strct)
		case t.KEY_FUNCTION:
			fmt.Println("parsing func")
			fn := p.parseFuncDecl()
			decls = append(decls, fn)

		default:
			msg := fmt.Sprintf("token %v is currently not supported as decl", p.head.T.String())
			panic(msg)
		}
	}
}

func (p *Parser) parseStruct() *ast.StructDecl {
	tok := p.tryPanic(t.IDENT)
	_, err := p.tryAndRead(t.LBRACE)
	if err != nil {
		panic(err)
	}

	p.consume(t.NEWLINE)

	decl := &ast.StructDecl{Name: tok.S, Pos: tok.Pos}
	/*
		for {
			p.consume(t.NEWLINE)
			if p.head.T == t.RBRACE {
				return decl
			}
			props := []*ast.Ident{}
			for (p.head.T != t.RBRACE) && (p.head.T != t.NEWLINE) {
				p.consume(t.COMMA)
				props = append(props, p.newIdent(p.head))
				p.read()
			}

			last := len(props) - 1
			for _, name := range props[:last] {
				decl.Parameters = append(decl.Parameters, &ast.TypedVariables{
					Name: name,
					Type: props[last],
				})
			}
			p.read()
		}
	*/
	params, err := p.parseParamsOrProps(t.RBRACE, false)
	if err != nil {
		panic(err)
	}
	decl.Parameters = params
	return decl
}

func (p *Parser) parseFuncDecl() *ast.FuncDecl {
	tok := p.tryPanic(t.IDENT)
	_, err := p.tryAndRead(t.LPAREN)
	if err != nil {
		panic(err)
	}

	fn := &ast.FuncDecl{Name: tok.S, Pos: tok.Pos}
	params, err := p.parseParamsOrProps(t.RPAREN, true)
	if err != nil {
		panic(err)
	}

	fn.Parameters = params

	p.consume(t.NEWLINE)
	if p.head.T != t.LBRACE {
		panic(p.error2(p.head, t.LBRACE))
	}

  p.consume(t.LBRACE)
	p.consume(t.NEWLINE)

	stmt := p.parseStmt()
	// cant have a 'main' stmt inside func decl
	// ...obviously
	if v, ok := stmt.(*ast.KallaxStmt); ok {
		spew.Dump(v)
		panic("it all went to shit")
	}

	fn.Body = stmt

	p.consume(t.NEWLINE)

	if p.head.T != t.RBRACE {
		panic(p.error2(p.head, t.RBRACE))
	}

  p.consume(t.RBRACE)
	p.consume(t.NEWLINE)

	return fn
}

type TokenCheck map[t.TokenCode]bool

var (
	paramsEnder = TokenCheck{
		t.RPAREN: true,
		t.EOF:    true,
	}
	propEnder = TokenCheck{
		t.RBRACE:  true,
		t.NEWLINE: true,
		t.EOF:     true,
	}
)

func (p *Parser) parseParamsOrProps(
	closing t.TokenCode,
	commaSeparated bool,
) ([]*ast.TypedVariables, error) {

	defer func() {
		p.consume(closing)
	}()

	list := []*ast.TypedVariables{}

	fmt.Println("entering loop")

	for { // for each chain in list
		p.consume(t.NEWLINE)

		if p.head.T == closing {
			return list, nil
		}

		if p.head.T != t.IDENT {
			panic(p.error2(p.head, t.IDENT))
		}

		list = append(list, p.parseChain()...)

		p.read()

		if commaSeparated {
			switch p.head.T {
			case closing:
				return list, nil
			case t.COMMA:
				p.consume(t.COMMA)
				break
			default:
				panic(p.error2(p.head, t.IDENT, t.COMMA, closing))
			}
		} else {
			switch p.head.T {
			case closing:
				return list, nil
			case t.NEWLINE:
				break
			default:
				panic(p.error2(p.head, t.NEWLINE, closing))

			}
		}
	}
}

func (p *Parser) parseChain() []*ast.TypedVariables {
	chain := []*ast.Ident{}

	for {
		chain = append(chain, p.newIdent(p.head))

		p.read()

		switch p.head.T {
		case t.IDENT: // end of chain
			list := []*ast.TypedVariables{}
			a2 := p.newIdent(p.head)

			for _, name := range chain {
				list = append(list, &ast.TypedVariables{
					Name: name,
					Type: a2,
				})
			}
			return list

		case t.COMMA:
			p.read()
			continue
		default:
			panic(p.error2(p.head, t.IDENT, t.COMMA))
		}
	}
}

// Requires left brace to be consume before call
func (p *Parser) parseSeqStmt() ast.Stmt {
	p.consume(t.NEWLINE)
	seq := &ast.SeqStmt{Pos: p.head.Pos, X: []ast.Stmt{}}

	fmt.Println("starting seq stmt")
	for {
		p.consume(t.NEWLINE)
		fmt.Println(p.head)
		switch p.head.T {
		case t.EOF, t.RBRACE:
			if len(seq.X) == 1 {
				return seq.X[0]
			}
			return seq
		case t.SEMICOLON, t.KEY_LET, t.IDENT:
			p.consume(t.SEMICOLON)
			fmt.Println("here")
			seq.X = append(seq.X, p.parseSimpleStmt())
			continue
		case t.COMMENT, t.COMMENT_BLOCK:
			p.read()
			continue
		default:
			msg := fmt.Sprintf("Hit default in parseSeqStmt: %v\n", p.head)
			panic(msg)
		}
	}

}

func (p *Parser) parseSimpleStmt() ast.Stmt {
	switch p.head.T {

	case t.KEY_LET:
		ident, err := p.try(t.IDENT)
		if err != nil {
			return ast.BadStmt(ident.Pos, err.Error())
		}
		eq, err := p.try(t.EQ)
		if err != nil {
			return ast.BadStmt(eq.Pos, err.Error())
		}
		p.read()

		ae := &ast.AssignStmt{I: &ast.Ident{Name: ident.S, Pos: ident.Pos}}
		ae.A = p.parseExpr()
		return ae
	case t.IDENT:
		ident := p.newIdent(p.head)
		p.read()
		switch p.head.T {
		case t.LPAREN:
			return &ast.ExprStmt{A: p.parseArguments(ident)}
		case t.PLUS, t.MINUS:
			crement, err := p.try(p.head.T)
			if err != nil {
				return &ast.ExprStmt{A: p.parseExpr()}
			}

			bin := &ast.BinaryExpr{
				A:  ident,
				OP: crement,
				B: &ast.BasicLit{
					Pos:   crement.Pos,
					Tok:   t.NUMBER,
					Value: "1",
				},
			}

			p.read()

			return &ast.AssignStmt{I: ident, A: bin}
		default:
			return &ast.ExprStmt{A: p.parseExpr()}
		}
	default:
		return &ast.ExprStmt{A: p.parseExpr()}
	}
}

func (p *Parser) parseExpr() ast.Expr {
	fmt.Println("parseExpr")
	spew.Dump(p.head)
	p.consume(t.NEWLINE)

	switch p.head.T {
	case t.MINUS:
		return p.unary()
	case t.NUMBER, t.FLOAT, t.BINARY, t.IDENT, t.STRING:
		return p.binop()
	case t.LPAREN:
		p.read()
		x := p.parseExpr()
		if p.head.T != t.RPAREN {
			return &ast.BadExpr{From: x.Start(), Value: "Unclosed parenthesis"}
		}
		return &ast.ParenExpr{A: x}
	case t.EOF, t.RBRACE:
		return nil
	default:
		msg := fmt.Sprintf("Hit default in parseExpr: %v\n", p.head)
		panic(msg)
	}

}

func (p *Parser) unary() ast.Expr {
	op := p.head
	p.read()
	switch p.head.T {
	case t.NUMBER, t.FLOAT, t.IDENT, t.LPAREN:
		a := p.parseExpr()
		switch at := a.(type) {
		case *ast.BinaryExpr:
			at.A = &ast.UnaryExpr{OP: op, A: at.A}
			return at
		default:
			return &ast.UnaryExpr{OP: op, A: a}
		}
	default:
		return &ast.BadExpr{
			From: p.head.Pos, Value: fmt.Sprintf("Bad unary: %s", p.head.S),
		}
	}
}

func (p *Parser) binop() ast.Expr {
	left := p.parseLhs()
	p.read()

	switch p.head.T {
	case t.MULTI, t.SLASH, t.DOT:
		bin := &ast.BinaryExpr{A: left, OP: p.head}
		p.read()

		B := p.parseExpr()
		switch ty := B.(type) {
		case *ast.BinaryExpr:
			bin.B = ty.A
			bi := *bin
			ty.A = &bi
			return ty
		default:
			bin.B = B
		}
		return bin
	case t.LPAREN:
		switch ty := left.(type) {
		case *ast.Ident:
			return p.parseArguments(ty)
		default:
			spew.Dump(left)
			panic("wtf are you trying to do with that left paren?")
		}
	case t.LBRACE:
		switch ty := left.(type) {
		case *ast.Ident:
			return p.parseStructArguments(ty)
		default:
			spew.Dump(left)
			panic("wtf are you trying to do with that left brace?")
		}

	case t.PLUS, t.MINUS:
		bin := &ast.BinaryExpr{A: left, OP: p.head}
		p.read()
		if p.head.T == bin.OP.T {

		}
		bin.B = p.parseExpr()
		return bin
	default:
		return left
	}
}

func (p *Parser) parseLhs() ast.Expr {
	switch p.head.T {
	case t.IDENT:
		return p.parseIdent()
	case t.STRING, t.NUMBER, t.FLOAT, t.BINARY:
		return &ast.BasicLit{Value: p.head.S, Pos: p.head.Pos, Tok: p.head.T}
	}
	return nil
}

func (p *Parser) parseIdent() *ast.Ident {
	x := &ast.Ident{Name: p.head.S, Pos: p.head.Pos}
	return x
}

func (p *Parser) parseArguments(ident *ast.Ident) *ast.CallExpr {
	ce := &ast.CallExpr{Name: ident.Name, Pos: ident.Pos}

	for {
		p.read()
		next := p.parseExpr()
		ce.Args = append(ce.Args, next)
		if p.head.T == t.RPAREN {
			p.read()
			return ce
		}

		if p.head.T != t.COMMA {
			tok := p.head
			panic("in args parsing: " + p.error(tok.Pos, tok.T, t.COMMA, t.LPAREN).Error())
		}
	}
}

func (p *Parser) parseStructArguments(ident *ast.Ident) *ast.ConstructExpr {
	ce := &ast.ConstructExpr{Name: ident.Name, Pos: ident.Pos}

	for {
		p.read()
		p.consume(t.NEWLINE)
		arg := p.parseIdent()
		if !p.expect(t.COLON) {
			panic("want colon after identifier in constructor argument")
		}
		p.read()

		val := p.parseExpr()

		ce.AddArgument(arg, val)
		if p.head.T == t.RBRACE {
			p.read()
			return ce
		}

		if p.head.T != t.COMMA {
			tok := p.head
			panic("in construct args parsing: " + p.error(tok.Pos, tok.T, t.COMMA, t.LPAREN).Error())
		}
	}
}
