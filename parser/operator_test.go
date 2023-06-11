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
				Token: token.EQ,
			},
			expected: "==",
		},
		{
			name: "!= operator",
			operator: parser.Operator{
				Token: parser.NEQ,
			},
			expected: "!=",
		},
		{
			name: "contains[cd] operator",
			operator: parser.Operator{
				Token:                parser.CONTAINS,
				CaseInsensitive:      true,
				DiacriticInsensitive: true,
			},
			expected: "CONTAINS[cd]",
		},
		{
			name: "and operator",
			operator: parser.Operator{
				Token: parser.AND,
			},
			expected: "AND",
		},
		{
			name: "beginswith operator",
			operator: parser.Operator{
				Token:           parser.BEGINSWITH,
				CaseInsensitive: true,
			},
			expected: "BEGINSWITH[c]",
		},
		{
			name: "like[c] operator",
			operator: parser.Operator{
				Token:           parser.LIKE,
				CaseInsensitive: true,
			},
			expected: "LIKE[c]",
		},
		{
			name: "==[cd] operator",
			operator: parser.Operator{
				Token:                parser.LIKE,
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
