grammar PLAN;

WS: [ \t\r\n] -> channel(HIDDEN);

// STARTING RULES
progFile: (stmts | fn)* EOF;

// FN

fn: 'fn' name = Identifier fnBody;

fnBody: '(' fnParams? ')' '{' stmt* '}';

fnParams: Identifier (',' Identifier)*;

// GENERAL STATEMENTS

stmts: stmt+;

stmt: simpleStmt ';' | compountStmt;

simpleStmt:
	assignment
	| includeStmt
	| methodInvoke
	| csInvoke
	| fnInvoke
	| breakStmt
	| continueStmt
	| returnStmt;

compountStmt: whileStmt | forStmt | ifStmt;

// COMPOUND STATEMENTS

whileStmt: 'while' exp '{' stmt* '}';

forStmt: 'for' assignment ';' exp ';' assignment '{' stmt* '}';

ifStmt: ifBlock elifBlock* elseBlock?;

ifBlock: 'if' exp '{' stmt* '}' # ifBlockStmt;

elifBlock: 'elif' exp '{' stmt* '}' # elifBlockStmt;

elseBlock: 'else' '{' stmt* '}' # elseBlockStmt;

// SIMPLE STATEMENTS

includeStmt: 'include' '(' exp ')';

returnStmt: 'return' exp?;

continueStmt: 'continue';

breakStmt: 'break';

assignment:
	name = Identifier '=' exp			# assignRegular
	| name = Identifier AssSum exp		# assignSum
	| name = Identifier AssSub exp		# assignSub
	| name = Identifier AssMul exp		# assignMul
	| name = Identifier AssDiv exp		# assignDiv
	| name = Identifier AssMod exp		# assignMod
	| name = Identifier AssPow exp		# assignPow
	| name = Identifier idxs '=' exp	# assignIdxsRegular
	| name = Identifier idxs AssSum exp	# assignIdxsSum
	| name = Identifier idxs AssSub exp	# assignIdxsSub
	| name = Identifier idxs AssMul exp	# assignIdxsMul
	| name = Identifier idxs AssDiv exp	# assignIdxsDiv
	| name = Identifier idxs AssMod exp	# assignIdxsMod
	| name = Identifier idxs AssPow exp	# assignIdxsPow;

list: '[' (exp (',' exp)*)? ']';

dictUnit: exp ':' exp;

dict: '{' (dictUnit (',' dictUnit)*)? '}';

idx: '[' exp ']';

idxs: idx idx*;

methodInvoke:
	var = Identifier '.' name = Identifier '(' (exp (',' exp)*)? ')' # identifierMethodInvoke;

fnInvoke:
	name = Identifier '(' (exp (',' exp)*)? ')' # identifierFnInvoke;

csInvoke:
	Closure name = Identifier '(' (exp (',' exp)*)? ')' # identifierCsInvoke;

exp:
	Integer											# expInteger
	| IntegerHex									# expIntegerHex
	| Null											# expNull
	| Bool											# expBool
	| Identifier									# expIdentifier
	| Float											# expFloat
	| String										# expString
	| closure										# expCs
	| methodInvoke									# expMethodInvoke
	| fnInvoke										# expFnInvoke
	| csInvoke										# expCsInvoke
	| exp idx										# expIdx
	| Subtract exp									# expNeg
	| Not exp										# expLogicalNot
	| <assoc = right> exp Pow exp					# expPow
	| exp op = (Multiply | Division | Modulus) exp	# expMulDivMod
	| exp op = (Add | Subtract) exp					# expSumSub
	| exp op = (GtEq | LtEq | Gt | Lt) exp			# expComparison
	| exp op = (Eq | Neq) exp						# expEqual
	| exp Xor exp									# expXor
	| exp And exp									# expLogicalAnd
	| exp Or exp									# expLogicalOr
	| '(' exp ')'									# expParentheses
	| list											# expList
	| dict											# expDict;

closure: 'fn' fnBody;

Eq: '==';
Neq: '!=';
Or: '||';
And: '&&';
Pow: '**';
GtEq: '>=';
LtEq: '<=';
AssSum: '+=';
AssSub: '-=';
AssMul: '*=';
AssDiv: '/=';
AssMod: '%=';
AssPow: '**=';
Gt: '>';
Lt: '<';
Multiply: '*';
Division: '/';
Modulus: '%';
Add: '+';
Subtract: '-';
Xor: '^';
Not: '!';
Closure: '@';

Bool: 'true' | 'false';

Null: 'null';

Identifier: [a-zA-Z_][a-zA-Z0-9_]*;

Integer: [0-9]+;

IntegerHex: [0][xX][0-9a-fA-F]+;

Float: [0-9]* '.' [0-9]+;

String:
	["] (~["\r\n\\] | '\\' ~[\r\n])* ["]
	| ['] ( ~['\r\n\\] | '\\' ~[\r\n])* ['];

Comment: ( '//' ~[\r\n]* | '/*' .*? '*/') -> channel(HIDDEN);