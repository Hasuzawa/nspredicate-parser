package scanner

import (
	"bytes"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

func (s *Scanner) scanString() (token.Token, string, error) {
	var buf bytes.Buffer
	ch, err := s.read()
	if err != nil {
		return token.ILLEGAL, "", err
	}

	lq := ch
	escape := false

	for {
		ch, err := s.read()
		if err != nil {
			break
		}
		if ch == '\\' {
			escape = !escape
			buf.WriteRune(ch)
			continue
		}
		// match the start quotes
		if ch == lq && !escape {
			escape = false
			break
		}
		buf.WriteRune(ch)
		escape = false
	}

	v := buf.String()

	return token.STRING, v, nil
}

func (s *Scanner) isQuote(ch rune) bool {
	switch ch {
	case '\'', '"', '`':
		return true
	default:
		break
	}
	return false
}
