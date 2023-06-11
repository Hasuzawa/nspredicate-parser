package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Hasuzawa/nspredicate-parser/parser"
	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

func TestOperator(t *testing.T) {
	for _, subtest := range []struct {
		name     string
		operator parser.Operator
		expected string
	}{
		{
			name: "== operator",
			operator: parser.Operator{
				Token: token.EQL,
			},
			expected: "==",
		},
		{
			name: "!= operator",
			operator: parser.Operator{
				Token: token.NEQ,
			},
			expected: "!=",
		},
		{
			name: "contains[cd] operator",
			operator: parser.Operator{
				Token:                token.CONTAINS,
				CaseInsensitive:      true,
				DiacriticInsensitive: true,
			},
			expected: "contains[cd]",
		},
		{
			name: "and operator",
			operator: parser.Operator{
				Token: token.AND,
			},
			expected: "and",
		},
		{
			name: "beginswith operator",
			operator: parser.Operator{
				Token:           token.BEGINSWITH,
				CaseInsensitive: true,
			},
			expected: "beginswith[c]",
		},
		{
			name: "like[c] operator",
			operator: parser.Operator{
				Token:           token.LIKE,
				CaseInsensitive: true,
			},
			expected: "like[c]",
		},
		{
			name: "==[cd] operator",
			operator: parser.Operator{
				Token:                token.EQL,
				CaseInsensitive:      true,
				DiacriticInsensitive: true,
			},
			expected: "==[cd]",
		},
	} {
		t.Run(subtest.name, func(t *testing.T) {
			assert.Equal(t, subtest.expected, subtest.operator.String())
		})
	}
}
