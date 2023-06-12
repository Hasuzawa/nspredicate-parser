package parser

import "github.com/Hasuzawa/nspredicate-parser/internal/version"

const (
	Name   = "nspredicate_go"
	Author = "Hasuzawa"
)

var (
	Version = version.Versioning{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}.String()
)
