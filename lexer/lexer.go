package lexer

import (
	"io"
	"strings"
	"time"
	"unicode"

	t "github.com/amstrups/nao/types"
	"github.com/davecgh/go-spew/spew"
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

func (l *Lexer) isEof() bool {
	if l.err != nil {
		if l.err == io.EOF {
			return true
		}
		panic(l.err)
	}
	return false
}

func (l *Lexer) unread() {
	if l.pos.IsAtStart() {
		println("at start")
		return
	}
	l.err = l.reader.UnreadRune()

	l.ch = l.old
	if l.pos.Column == 1 {
		l.pos.Line = max(1, l.pos.Line-1)
	}
	l.pos.Column = max(0, l.pos.Column-1)
}

func (l *Lexer) next(c rune) t.Position {
	pos := l.pos
	for l.ch != c || l.err != io.EOF {
		l.advance()
	}
	return pos

}

func (l *Lexer) newline() {
	l.pos.Line++
	l.pos.Column = 1
}

func (l *Lexer) skip(c rune) t.Position {
	pos := l.pos
	for l.ch == c || l.err != io.EOF {
		l.advance()
	}
	return pos

}

func (l *Lexer) advance() {
	l.old = l.ch
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
		l.advance()
		if l.isEof() {
			return t.Token{T: t.EOF, Pos: l.pos, S: ""}
		}
		switch l.ch {
		case '\n':
			return t.Token{T: t.NEWLINE, Pos: l.pos, S: "\n"}
		case ' ':
			continue
		case '.':
			return t.Token{T: t.DOT, Pos: l.pos, S: "."}
		case ',':
			return t.Token{T: t.COMMA, Pos: l.pos, S: ","}
		case ';':
			return t.Token{T: t.SEMICOLON, Pos: l.pos, S: ";"}
		case ':':
			return t.Token{T: t.COLON, Pos: l.pos, S: ":"}
		case '"':
			return l.string()
		case '\'':
			return t.Token{T: t.SINGLEQUOTE, Pos: l.pos, S: "'"}
		case '(':
			return t.Token{T: t.LPAREN, Pos: l.pos, S: "("}
		case ')':
			return t.Token{T: t.RPAREN, Pos: l.pos, S: ")"}
		case '[':
			return t.Token{T: t.LBRACKET, Pos: l.pos, S: "["}
		case ']':
			return t.Token{T: t.RBRACKET, Pos: l.pos, S: "]"}
		case '{':
			return t.Token{T: t.LBRACE, Pos: l.pos, S: "{"}
		case '}':
			return t.Token{T: t.RBRACE, Pos: l.pos, S: "}"}
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
		case '0':
			return l.binOrNumber()
		case '#':
			return l.comment()
		default:
			if unicode.IsDigit(l.ch) {
				return l.number()
			} else if unicode.IsLetter(l.ch) {
				return l.ident()
			}
		}
	}
}

func (l *Lexer) binOrNumber() t.Token {
	pos := l.pos

	l.advance()
	if l.ch != 'b' {
		l.unread()
		n := l.number()
		return n
	}

	l.advance()

	s := ""
	for {
		if l.isEof() {
			return t.Token{T: t.BINARY, Pos: pos, S: s}
		}

		switch l.ch {
		case '0', '1':
			s += string(l.ch)
			l.advance()
			continue
		case ' ', '\n':
			l.unread()
			return t.Token{T: t.BINARY, Pos: pos, S: s}
		default:
			return t.Token{T: t.ILLEGAL, Pos: l.pos, S: "Expected x in {0,1}, found " + string(l.ch)}
		}
	}
}

func (l *Lexer) number() t.Token {
	tok := t.Token{T: t.NUMBER, Pos: l.pos, S: string(l.ch)}

	for {
		l.advance()
		if l.ch == '\n' {
			l.unread()
			return tok
		}
		if l.ch == '.' {
			if tok.T == t.NUMBER {
				tok.T = t.FLOAT
				tok.S += string(l.ch)
				continue
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
		l.advance()
		if unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) || l.ch == '\'' {
			tok.S += string(l.ch)
			continue
		}
		l.unread()

		tok.T = t.CheckIfKeyword(tok.S)
		spew.Dump(tok)
		return tok
	}
}

func (l *Lexer) string() t.Token {
	tok := t.Token{T: t.STRING, Pos: l.pos, S: string(l.ch)}
	for {
		l.advance()
		tok.S += string(l.ch)
		if l.ch == '"' {
			return tok
		} else if l.err == io.EOF {
			tok.T = t.ILLEGAL
			return tok
		}
	}
}

func (l *Lexer) comment() t.Token {
	tok := t.Token{T: t.COMMENT, Pos: l.pos}
	l.advance()
	if l.ch == '[' {
		tok.T = t.COMMENT_BLOCK
		tok.S = l.commentblock()
		return tok
	}

	for {
		if l.isEof() {
			return tok
		}
		if l.ch == '\n' {
			l.unread()
			return tok
		}

		tok.S += string(l.ch)
		l.advance()
	}
}

func (l *Lexer) commentblock() string {
	var comment string
	for {
		if l.isEof() {
			return comment
		}

		if l.ch == ']' {
			l.advance()
			if l.ch == '#' {
				return comment
			}
			comment += string(']') + string(l.ch)
		}
		l.advance()
	}
}
