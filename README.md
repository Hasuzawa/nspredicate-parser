## Introduction

Objective-C NSPredicate, NSExpression parser in Golang. Created for own usage and work.

## Compiling & Running Objective-C and Swift file

Objective-C
```
gcc -framework Foundation ./objective-c/example.m
```

Swift
```
swift ./swift/example.swift
```

## The [cd] flags

[c] flag will enable case-insensitive filtering.

[d] flag will enable diacritic-insensitive filtering.

Examples of diacritics include French accent marks, German umlauts and 'sharp s', Spanish 'tilde n'.

e.g.

`René Descartes` vs `Rene Descartes`

`Français` vs `Francais`

`reiß` vs `reiss`

`español` vs `espanol`

## EBNF (Extended Backus-Naur form)

`alphabet = "a" | "b" | ... | "z" | "A" | "B" | ... | "Z" ;`  (omitted for brevity)

`digit = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;`

`operator = "==" | "=" | "!=" | "<" | "<=" | ">=" | ">" | "beginswith" | "endswith" | "contains" | "matches" ;`

`logical operator = "and" | "or" | "all" | "any" ;`

`special value = "true" | "false" | "YES" | "NO" | "NULL" | "nil" ;`

`blank = " " | "\n" \ "t" | "\r" | "\f" | "\b" ;`

`symbol = "(" | ")" | "[" | "]" | "{" | "}" | "," ;`

`identifier = alphabet , { alphabet | digit | _ } ;`

## Documentation

[NSPredicate | Apple Documentation](https://developer.apple.com/documentation/foundation/nspredicate)

[NSExpression | Apple Documentation](https://developer.apple.com/documentation/foundation/nsexpression)

[Predicate Programming Guide | Apple Documentation Archive](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html)

[token module | Go Standard Library](https://pkg.go.dev/go/token)

[scanner module | Go Standard Library](https://pkg.go.dev/go/scanner)

[ast module | Go Standard Library](https://pkg.go.dev/go/ast)


## Reference & Article

[Article on NSPredicate by Fluffy.es (Axel)](https://nspredicate.xyz/)

[Handwritten Parsers & Lexers in Go by Ben Jonhnson](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)

[SQL parser in Go by Mariano Gappa](https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/)

[NSPredicate by Mattt](https://nshipster.com/nspredicate/)

[Introduction to Abstract Syntax by Arkadiusz Ziobrowski](https://tech.ingrid.com/introduction-ast-golang/)
