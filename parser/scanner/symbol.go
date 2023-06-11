package scanner

import (
	"bytes"
	"strings"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

func (s *Scanner) scanSymbol() (tok token.Token, lit string, err error) {
	var buf bytes.Buffer
	ch, err := s.read()
	if err != nil {
		return token.ILLEGAL, "", err
	}

	buf.WriteRune(ch)

	for {
		ch, err := s.read()
		if err != nil {
			break
		}
		if IsSymbol(ch) {
			buf.WriteRune(ch)
			continue
		}
		s.unread()
		break
	}

	v := buf.String()

	switch strings.ToUpper(v) {
	case "==", "=":
		return token.EQL, v, nil
	case "!=":
		return token.NEQ, v, nil
	case ">=", "=>":
		return token.GE, v, nil
	case "<=", "=<":
		return token.LE, v, nil
	case ">":
		return token.GT, v, nil
	case "<":
		return token.LT, v, nil
	case "(":
		return token.LPAREN, v, nil
	case ")":
		return token.RPAREN, v, nil
	case "[":
		return token.LBRACK, v, nil
	case "]":
		return token.RBRACK, v, nil
	case "{":
		return token.LBRACE, v, nil
	case "}":
		return token.RBRACE, v, nil
	case ",":
		return token.COMMA, v, nil
	default:
		break
	}
	return token.ILLEGAL, v, err
}

func IsSymbol(ch rune) bool {
	switch ch {
	case '=', '!', '>', '<', '(', ')', '[', ']', '{', '}', ',':
		return true
	default:
		break
	}
	return false
}
