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
		{
			name: "identifiers",
			line: "page __name__ CoUnT _ iron9 i64 DEBUG_MODE _123_",
			tokens: []token.Token{
				token.IDENT, token.WS, token.IDENT, token.WS, token.IDENT, token.WS, token.IDENT, token.WS, token.IDENT,
				token.WS, token.IDENT, token.WS, token.IDENT, token.WS, token.IDENT,
			},
		},
		{
			name: "reserevd keywords",
			line: "and or all any true false null nil",
			tokens: []token.Token{
				token.AND, token.WS, token.OR, token.WS, token.ALL, token.WS, token.ANY, token.WS, token.TRUE,
				token.WS, token.FALSE, token.WS, token.NULL, token.WS, token.NIL,
			},
		},
		{
			name: "string comparison",
			line: "contains beginswith endswith like CONTAINS BEGINSWITH ENDSWITH LIKE",
			tokens: []token.Token{
				token.CONTAINS, token.WS, token.BEGINSWITH, token.WS, token.ENDSWITH, token.WS, token.LIKE, token.WS,
				token.CONTAINS, token.WS, token.BEGINSWITH, token.WS, token.ENDSWITH, token.WS, token.LIKE,
			},
		},
		{
			name: "int and float",
			line: "0 2 3.14 0.75 99. .57 00.0 +3 -7.1 1_234 567_8.9 0.1_234",
			tokens: []token.Token{
				token.INT, token.WS, token.INT, token.WS, token.FLOAT, token.WS, token.FLOAT, token.WS, token.FLOAT,
				token.WS, token.FLOAT, token.WS, token.FLOAT, token.WS, token.INT, token.WS, token.FLOAT, token.WS,
				token.FLOAT, token.WS, token.FLOAT, token.WS, token.FLOAT,
			},
		},
		{
			name: "string",
			line: `"abcdef" "12345" '__dir__' '~p()&#$*(!$)' "\n\t\x\1\2\3\"\'" '\"\'\"|\"\'|'`,
			tokens: []token.Token{
				token.STRING, token.WS, token.STRING, token.WS, token.STRING, token.WS, token.STRING, token.WS,
				token.STRING, token.WS, token.STRING,
			},
		},
		{
			name: "simple predicate 1",
			line: `name beginswith 'Mary' and page > 7 `,
			tokens: []token.Token{
				token.IDENT, token.WS, token.BEGINSWITH, token.WS, token.STRING, token.WS, token.AND, token.WS,
				token.IDENT, token.WS, token.GT, token.WS, token.INT, token.WS,
			},
		},
		{
			name: "simple predicate 2",
			line: `DEBUG == true OR __dir__ contains[cd] "staticasset"`,
			tokens: []token.Token{
				token.IDENT, token.WS, token.EQL, token.WS, token.TRUE, token.WS, token.OR, token.WS,
				token.IDENT, token.WS, token.CONTAINS, token.LBRACK, token.IDENT, token.RBRACK, token.WS, token.STRING,
			},
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
