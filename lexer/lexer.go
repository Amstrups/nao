package lexer

import (
	"io"
	t "nao/types"
	"strings"
	"time"
	"unicode"
)

type Lexer struct {
	pos    t.Position
	reader *strings.Reader

	old rune
	ch  rune
	err error
}

func New(s string) *Lexer {
	return &Lexer{
		pos: t.Position{
			Line:   1,
			Column: 0,
		},
		reader: strings.NewReader(s),
	}
}

func (l *Lexer) unread() {
	if l.pos.IsAtStart() {
		return
	}
	l.err = l.reader.UnreadRune()

	l.pos.Line = max(1, l.pos.Line-1)
	l.pos.Column = max(0, l.pos.Column-1)
}

func (l *Lexer) next() {
	l.ch, _, l.err = l.reader.ReadRune()
	if l.ch == '\n' {
		l.pos.Line++
		l.pos.Column = 0
	}

	l.pos.Column++
}

func (l *Lexer) Lex() t.Token {
	ch := make(chan t.Token, 1)
	go func() {
		ch <- l.lex()
	}()
	select {
	case x := <-ch:
		return x
	case <-time.After(3 * time.Second):
		panic("lexer timed out")
	}
}

func (l *Lexer) lex() t.Token {

	for {
		l.next()
		if l.err != nil {
			if l.err == io.EOF {
				return t.Token{T: t.EOF, Pos: l.pos, S: ""}
			}
			return t.Token{T: t.ILLEGAL, Pos: l.pos, S: "ILLEGAL"}
		}

		switch l.ch {
		case ' ':
			continue
		case '.':
			return t.Token{T: t.DOT, Pos: l.pos, S: "."}
		case ';':
			return t.Token{T: t.SEMICOLON, Pos: l.pos, S: ";"}
		case '"':
			return t.Token{T: t.DOUBLEQUOTE, Pos: l.pos, S: "\""}
		case '\'':
			return t.Token{T: t.SINGLEQUOTE, Pos: l.pos, S: "'"}
		case '(':
			return t.Token{T: t.LPAREN, Pos: l.pos, S: "("}
		case ')':
			return t.Token{T: t.RPAREN, Pos: l.pos, S: ")"}
		case '=':
			return t.Token{T: t.EQ, Pos: l.pos, S: "="}
		case '+':
			return t.Token{T: t.PLUS, Pos: l.pos, S: "+"}
		case '-':
			return t.Token{T: t.MINUS, Pos: l.pos, S: "-"}
		case '*':
			return t.Token{T: t.MULTI, Pos: l.pos, S: "*"}
		case '\\':
			return t.Token{T: t.BACKSLASH, Pos: l.pos, S: "\\"}
		case '/':
			return t.Token{T: t.SLASH, Pos: l.pos, S: "/"}
		default:
			if unicode.IsDigit(l.ch) {
				return l.number()
			} else if unicode.IsLetter(l.ch) {
				return l.ident()
			}

		}
	}
}

func (l *Lexer) number() t.Token {
	tok := t.Token{T: t.NUMBER, Pos: l.pos, S: string(l.ch)}
	for {
		l.next()
		if l.ch == '\n' {
			return tok
		}
		if l.ch == '.' {
			if tok.T == t.NUMBER {
				tok.T = t.FLOAT
			} else {
				tok := t.Token{T: t.ILLEGAL, Pos: tok.Pos, S: "Illegal dot after float"}
				return tok

			}

		}
		if unicode.IsDigit(l.ch) {
			tok.S += string(l.ch)
			continue
		}
		l.unread()

		return tok
	}
}

func (l *Lexer) ident() t.Token {
	tok := t.Token{T: t.IDENT, Pos: l.pos, S: string(l.ch)}
	for {
		l.next()
		if l.ch == '\n' {
			return tok
		}
		if unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
			tok.S += string(l.ch)
			continue
		}
		l.unread()

		return tok
	}
}

func (l *Lexer) string() t.Token {
	tok := t.Token{T: t.IDENT, Pos: l.pos, S: string(l.ch)}
	for {
		l.next()
		if l.ch == '"' {
			return tok
		}
		if unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
			tok.S += string(l.ch)
			continue
		}

		l.unread()

		return tok
	}
}
