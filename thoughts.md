# NAO Design
## Proposals/TODOs
- Functions
- Arrays
- Lists
- Scopes
- Declarations
- Multifile-support
- Imports

## Missing Tests
- BinaryExpr
- UnaryExpr
- BasicLit
- Ident
- ParenExpr
- SeqStmt
 
## Pure functions
Idea is to limit side effects from a function and provide guarantees to a user, about the behaviour of a function. 
Similar to async/await, functions 
### Initial syntax thoughts
Normal function:
```
func Foo(arg a) b { ... }
```
"Pure" function
```
pure Foo(arg a) b { ... }
```
