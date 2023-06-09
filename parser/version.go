package parser

import "github.com/Hasuzawa/nspredicate-parser/internal/version"

const (
	NAME   = "nspredicate-go"
	AUTHOR = "Hasuzawa"
)

var (
	VERSION = version.Versioning{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}.String()
)
