package scanner

import (
	"bytes"
	"unicode"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

func (s *Scanner) scanSpace() (token.Token, string, error) {
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
		if unicode.IsSpace(ch) {
			buf.WriteRune(ch)
			continue
		}
		s.unread()
		break
	}

	return token.WS, "", nil
}
