# NAO Syntax (Thoughts)
```ML
Program := 
    | Statement
    | nil

Statement :=
    | Sequence
    | nil

Sequence :=
    | SeqHead SeqTail
    | SeqHead
    | nil

SeqHead :=
    | Expr
    | Assignment

SeqTail :=
    | ';' SeqHead SeqTail
    | nil

Assignment :=
    | Ident '=' Expr

Expr := 
    | Unary
    | BinaryOp
    | Paren
    | Basic
    | Ident

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
    | 'x4'
    | 'x8'
    | 'x16'
    | 'x32'
    | 'x64'

```
