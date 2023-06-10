package array

import (
	"fmt"
	"strings"
)

// NSArray is a Go representation of Objective-C NSArray.
//
// Data of different data type can be put into it.
type NSArray struct {
	elements []any
}

func Init(values ...any) *NSArray {
	return &NSArray{
		elements: values,
	}
}

func (a NSArray) String() string {
	var es []string
	for _, element := range a.elements {
		var v any
		switch element.(type) {
		case string:
			v = fmt.Sprintf("%q", element)
		default:
			v = element
		}
		es = append(es, fmt.Sprint(v))
	}
	if len(es) == 0 {
		return "()"
	}
	return fmt.Sprintf("( %s )", strings.Join(es, " , "))
}
