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

	// string comparison operators

	// contains
	CONTAINS
	// beginswith
	BEGINSWITH
	// endswith
	ENDSWITH
	// like
	LIKE

	// logical operators

	// and
	AND
	// or
	OR
	// all
	ALL
	// any
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

// some token has one and only one valid (or preferred) representation. This facilitates formatting token.
func (t Token) Literal() string {
	switch t {
	case EQL:
		return "=="
	case NEQ:
		return "!="
	case LT:
		return "<"
	case LE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	case CONTAINS:
		return "contains"
	case BEGINSWITH:
		return "beginswith"
	case ENDSWITH:
		return "endswith"
	case LIKE:
		return "like"
	case AND:
		return "and"
	case OR:
		return "or"
	case ALL:
		return "all"
	case ANY:
		return "any"
	case YES:
		return "YES"
	case NO:
		return "NO"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case NULL:
		return "NULL"
	case NIL:
		return "nil"
	case LPAREN:
		return "("
	case LBRACK:
		return "["
	case LBRACE:
		return "{"
	case RPAREN:
		return ")"
	case RBRACK:
		return "]"
	case RBRACE:
		return "}"
	case COMMA:
		return ","
	default:
		break
	}
	return ""
}
