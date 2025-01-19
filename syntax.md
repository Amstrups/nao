# NAO Syntax (Thoughts)
This serves as my thoughts on how the syntax for `nao` is defined.

```sh
Program := 
    | Statement
    | nil

Statement :=
    | Sequence
    | nil

Sequence :=
    | SimpleStmt SeqTail
    | SimpleStmt
    | nil

SeqTail :=
    | ';' SimpleStmt SeqTail
    | nil

Struct := -- TODO
    | "struct" CType Properties(?) -- TODO

SimpleStmt :=
    | Assignment
    | ExprStmt

Assignment :=
    | "let" Ident ':' '=' Expr
    | "let" Ident ':' Type '=' Expr
    | Ident '=' Expr

Type :=
    | 
    | NumberType
    | "string"
    | "bool"

NumberType :=
    | "int" -- Drop this or "int64"?
    | "int8"
    | "int64"
    | "float" -- Drop this or "float64"?
    | "float8"
    | "float64" 

Expr := 
    | Unary
    | BinaryOp
    | Paren
    | Basic
    | Ident
    | String

UnaryOp :=
    | '-' Basic
    | '-' Ident

BinaryOp :=
    | Expr Op Expr

Paren :=
    | '(' Expr ')'

BasicLiteral :=
    | Number
    | Float
    | String
    | Binary
    | VectorValue

Op :=
    | '+'
    | '-'
    | '*'
    | '/'
    | '^' -- TODO

Binary := 
    | '0b' {'0'|'1'}+

VectorType :=
    | '[' T ',' Number ']' 

VectorValue :=
    | '[' Expr {',', Expr}*  ']'
    | '[' ']'
```


# Notation
The notation above should be similar to *context-free grammars*, where each production rule have the following structure

### Basis
```sh
RULE :=
    | Non-terminal 
    | 'terminal character'
    | "terminal string"
```

### Branching productions

Each line beginnig with an `'|'`, defines an `or`, such that
```sh
RULE1 := 
    | Non-terminal1

RULE1 :=
    | Non-terminal2
```
is equivalent to
```sh
RULE1 :=
    | Non-terminal1
    | Non-terminal2
```

### Production inference (TODO)
Lines starting with `>`, defines *restriction* on the line above

```sh
NUMBER := 
    | '1'
    | '2'
    | '3'

OddIndex :=
    | '[' A = number ']'
    > A \in ['1', '3']
```

Due to limitations of markdown (and the fact I don't want to keep rendering some document with proper "inference rule" notation), this is a decent middle ground

Still working on how to define inference in a more compact way
