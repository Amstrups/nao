# NAO Syntax (Thoughts)
```s
Program := 
    | Statement
    | nil

Statement :=
    | Sequence
    | Single -- 
    | nil

Sequence :=
    | SeqHead SeqTail
    | SeqHead
    | nil

SeqHead :=
    | Expr
    | Assignment

SeqTail :=
    | SeqHead SeqTail
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

Basic :=
    | Number
    | Float
    | String

Op :=
    | '+'
    | '-'
    | '*'
    | '/'
    | '^' -- TODO
```
