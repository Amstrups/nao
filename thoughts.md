# NAO Design
## Proposals/TODOs
- Functions
- Arrays
- Lists
- Scopes
- Declarations
- Multifile-support
- Imports
- Pure functions
- "nao what" (yes this is obviously a serious language)
    - Command for dictionary of concepts?
    - Maybe should be part of `arit` so "arit what {CONCEPT}"
    - Ex. "{arit|nao} what hoeffding inequality"
        - Requires search functionality
- Result types ( $a \times err?$ )

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
Normal function, with no guarentees of purity:
```
func Foo(arg a) b { ... }
```
"Pure" function
```
pure Foo(arg a) b { ... }
```
### Remark
After discussing this idea with an acquaintance, they pointed out that this field is studied at Aarhus University for the language [Flix](https://flix.dev). 
Further expansion of the *purity* concept may draw large amount of inspiration from the Flix language.  


