# NAO Syntax (Thoughts)
```ML
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
    | Ident ':' '=' Expr
    | Ident ':' Type '=' Expr

Type :=
    | NumberType
    | "string"
    | "bool"

NumberType :=
    | "int"
    | "int8"
    | "int64"
    | "float"
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

Op :=
    | '+'
    | '-'
    | '*'
    | '/'
    | '^' -- TODO

Binary :=
    | '0b' ZOs+ P2 
    | '0b' ZOs+

ZOs+ :=
    | '0'ZOs*
    | '1'ZOs*

ZOs* :=
    | '0'ZOs*
    | '1'ZOs*
    | nil

P2 := 
    | 'x8' 
    | 'x64' 
```
