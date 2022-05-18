grammar SimpleBoolean;

parse: expression EOF;

expression:
	LPAREN expression RPAREN
	| NOT expression
	| expression (	GT
		| GE
		| LT
		| LE
		| EQ
		| NEQ) expression
	| expression (AND | OR) expression
	| expression THEN expression (ELSE expression)?
	| expression IF expression (ELSE expression)?
	| SQRT LPAREN expression RPAREN
	| expression (MULT|DIV) expression
        | expression (ADD|SUB) expression
	| boolean
	| DECIMAL
;

comparator:
	GT
	| GE
	| LT
	| LE
	| EQ
	| NEQ;

binary: AND | OR;

boolean: TRUE | FALSE;

AND: 'AND' | '&';
OR: 'OR' | '|';
NOT: 'NOT' | '!';
TRUE: 'TRUE' | 'True' | 'true';
FALSE: 'FALSE' | 'False' | 'false';
IF: 'IF';
THEN: 'THEN' | 'Then' | 'then' | '->';
ELSE: 'ELSE' | 'Else' | 'else';
GT: '>' | 'GT';
GE: '>=' | 'GE' | '≥';
LT: '<' | 'LT';
LE: '<=' | 'LE' | '≤';
EQ: '==' | 'EQ';
NEQ: '!=' | '<>' | 'NEQ';
LPAREN: '(';
RPAREN: ')';
DECIMAL: '-'? [0-9]+ ( '.' [0-9]+)?;
SQRT: 'SQRT';
MULT: '*';
DIV: '/';
ADD: '+';
SUB: '-';
WS: [ \r\t\u000C\n]+ -> skip;
