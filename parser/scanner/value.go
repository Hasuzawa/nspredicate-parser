package scanner

import (
	"bytes"
	"strconv"
	"unicode"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

func (s *Scanner) scanValue() (token.Token, string, error) {
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
		if unicode.IsDigit(ch) || ch == '.' || ch == '_' {
			buf.WriteRune(ch)
			continue
		}
		s.unread()
		break
	}

	v := buf.String()
	if _, err := strconv.ParseInt(v, 10, 32); err == nil {
		return token.INT, v, nil
	} else if _, err := strconv.ParseFloat(v, 32); err == nil {
		return token.FLOAT, v, nil
	}
	return token.ILLEGAL, v, nil
}
