package ns

import (
	"github.com/Hasuzawa/nspredicate-parser/ns/array"
)

func NSArray(values ...any) *array.NSArray {
	return array.Init(values)
}
