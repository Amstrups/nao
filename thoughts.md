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
Similar to async/await, `pure` can only call `pure`. Pure functions used in a cryptographic setting are should use keyword `puré`
### Initial syntax thoughts
Normal function:
```
func Foo(arg a) b { ... }
```
"Pure" function
```
pure Foo(arg a) b { ... }
```
