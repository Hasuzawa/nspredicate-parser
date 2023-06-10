## Introduction

Objective-C NSPredicate, NSExpression parser in Golang. Created for own usage and work.

## Compiling Objective-C and Swift file

Objective-C
```
gcc -framework Foundation ./objective-c/example.m
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

## Documentation

[NSPredicate | Apple Documentation](https://developer.apple.com/documentation/foundation/nspredicate)

[NSExpression | Apple Documentation](https://developer.apple.com/documentation/foundation/nsexpression)

[Predicate Programming Guide | Apple Documentation Archive](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html)

[token module | Go Standard Library](https://pkg.go.dev/go/token)

[scanner module | Go Standard Library](https://pkg.go.dev/go/scanner)


## Reference & Article

[Article on NSPredicate by Fluffy.es (Axel)](https://nspredicate.xyz/)

[Handwritten Parsers & Lexers in Go by Ben Jonhnson](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)

[SQL parser in Go by Mariano Gappa](https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/)
