package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Hasuzawa/nspredicate-parser/ns/array"
)

func TestNSArray(t *testing.T) {
	for _, subtest := range []struct {
		name     string
		values   []any
		expected string
	}{
		{
			name:     "empty array",
			values:   []any{},
			expected: "()",
		},
		{
			name:     "int array",
			values:   []any{2, 3, 5, 7, 11, 13, 17, 19, 23, 27},
			expected: "( 2 , 3 , 5 , 7 , 11 , 13 , 17 , 19 , 23 , 27 )",
		},
		{
			name:     "string array",
			values:   []any{"lorem", "ipsum", "dolor", "sit", "amet"},
			expected: `( "lorem" , "ipsum" , "dolor" , "sit" , "amet" )`,
		},
	} {
		t.Run(subtest.name, func(t *testing.T) {
			arr := array.Init(subtest.values...)
			assert.IsType(t, new(array.NSArray), arr)
			assert.Equal(t, subtest.expected, arr.String())
		})
	}
}
