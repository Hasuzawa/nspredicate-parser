package parser

import (
	"fmt"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

type Operator struct {
	token.Token
	CaseInsensitive      bool
	DiacriticInsensitive bool
}

const (
	FlgCaseInsensitive      = "c"
	FlgDiacriticInsensitive = "d"
)

func (o Operator) String() string {
	var flg string
	if o.CaseInsensitive {
		flg += FlgCaseInsensitive
	}
	if o.DiacriticInsensitive {
		flg += FlgDiacriticInsensitive
	}
	if len(flg) == 0 {
		return o.Token.String()
	}
	return fmt.Sprintf("%s[%s]", o.Token.String(), flg)
}
