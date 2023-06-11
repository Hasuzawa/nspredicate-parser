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

	switch strings.ToUpper(buf.String()) {
	case "==", "=":
		return token.EQL, "==", nil
	case "!=":
		return token.NEQ, "!=", nil
	case ">=", "=>":
		return token.GE, ">=", nil
	case "<=", "=<":
		return token.LE, "<=", nil
	case ">":
		return token.GT, ">", nil
	case "<":
		return token.LT, "<", nil
	case "(":
		return token.LPAREN, "(", nil
	case ")":
		return token.RPAREN, ")", nil
	case "[":
		return token.LBRACK, "[", nil
	case "]":
		return token.RBRACK, "]", nil
	case "{":
		return token.LBRACE, "{", nil
	case "}":
		return token.RBRACE, "}", nil
	case ",":
		return token.COMMA, ",", nil
	case ".":
		return token.PERIOD, ".", nil
	default:
		break
	}
	return token.ILLEGAL, "", err
}

func IsSymbol(ch rune) bool {
	switch ch {
	case '=', '!', '>', '<', '(', ')', '[', ']', '{', '}', ',', '.':
		return true
	default:
		break
	}
	return false
}
