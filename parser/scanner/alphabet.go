package scanner

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

var (
	identifierRegex = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
)

func (s *Scanner) scanAlphabet() (token.Token, string, error) {
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
		if unicode.IsLetter(ch) || ch == '_' || unicode.IsDigit(ch) {
			buf.WriteRune(ch)
			continue
		}

		s.unread()
		break
	}

	v := buf.String()

	switch strings.ToUpper(v) {
	case "AND":
		return token.AND, v, nil
	case "OR":
		return token.OR, v, nil
	case "ALL":
		return token.ALL, v, nil
	case "ANY":
		return token.ANY, v, nil
	case "CONTAINS":
		return token.CONTAINS, v, nil
	case "BEGINSWITH":
		return token.BEGINSWITH, v, nil
	case "ENDSWITH":
		return token.ENDSWITH, v, nil
	case "LIKE":
		return token.LIKE, v, nil
	case "TRUE":
		return token.TRUE, v, nil
	case "FALSE":
		return token.FALSE, v, nil
	case "NULL":
		return token.NULL, v, nil
	case "NIL":
		return token.NIL, v, nil
	}

	if match := identifierRegex.FindString(v); match != "" {
		return token.IDENT, v, nil
	}

	return token.ILLEGAL, "", nil
}
