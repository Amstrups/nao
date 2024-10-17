package types

import "fmt"

type Position struct {
	Line, Column int
}

func (p Position) IsAtStart() bool {
	return p.Line == 1 && p.Column == 0
}

func Default() Position {
	return Position{
		Line:   1,
		Column: 0,
	}
}

func (pos Position) String() string {
	return fmt.Sprintf("%-2d:%d]", pos.Line, pos.Column)
}
