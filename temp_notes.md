
("fn"|"struct") IDENT 
OPENING 
( {\n}*
{ (IDENT COMMA)* IDENT IDENT ('\n' | COMMA ['\n']*) }*
{\n}* )*
CLOSING

