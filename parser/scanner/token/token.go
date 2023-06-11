package token

type Token uint

const (
	ILLEGAL Token = iota
	EOF
	WS

	// identifiers, e.g. variable name
	IDENT

	// data type

	INT
	FLOAT
	STRING

	// operators

	// == , =
	EQL
	// !=
	NEQ
	// <
	LT
	// <= , =<
	LE
	// >
	GT
	// >= , =>
	GE

	// logical operators

	// AND
	AND
	// OR
	OR
	// ALL
	ALL
	// ANY
	ANY

	// reserved words

	// true , TRUE
	TRUE
	// false , FALSE
	FALSE
	// YES
	YES
	// NO
	NO

	// NULL
	NULL
	// nil
	NIL

	// symbols

	// (
	LPAREN
	// [
	LBRACK
	// {
	LBRACE

	// )
	RPAREN
	// ]
	RBRACK
	// }
	RBRACE

	// ,
	COMMA
)
