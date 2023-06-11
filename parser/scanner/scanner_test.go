package scanner_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner"
	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

func TestScanner(t *testing.T) {
	for _, subtest := range []struct {
		name   string
		line   string
		tokens []token.Token
	}{
		{
			name:   "blanks",
			line:   "       \t      \n     \t\n    \n    \t  \n\t",
			tokens: []token.Token{token.WS},
		},
		{
			name: "space and symbol",
			line: "== != > >= < <= ( ) [ ] { } ,",
			tokens: []token.Token{
				token.EQL, token.WS, token.NEQ, token.WS, token.GT, token.WS, token.GE, token.WS, token.LT, token.WS,
				token.LE, token.WS, token.LPAREN, token.WS, token.RPAREN, token.WS, token.LBRACK, token.WS, token.RBRACK,
				token.WS, token.LBRACE, token.WS, token.RBRACE, token.WS, token.COMMA},
		},
	} {
		t.Run(subtest.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(subtest.line))
			scanner := scanner.NewScanner(reader)

			var tokens []token.Token
			for {
				tok, _, err := scanner.Scan()
				if err != nil || tok == token.EOF {
					break
				}
				tokens = append(tokens, tok)
			}
			assert.Equal(t, subtest.tokens, tokens)
		})
	}
}
