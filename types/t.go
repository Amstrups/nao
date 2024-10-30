package types

//go:generate stringer -type=T
type T uint

const (
	T_ILLEGAL T = iota

	T_INT
	T_BINARY
	T_FLOAT

	T_STRING

	T_BOOL

	T_VOID

	T_IDENT
)

var TtT map[TokenCode]T = map[TokenCode]T{
	NUMBER: T_INT,
	FLOAT:  T_FLOAT,
	BINARY: T_BINARY,
	STRING: T_STRING,
}
