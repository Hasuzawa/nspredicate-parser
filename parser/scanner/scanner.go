package scanner

import (
	"bufio"
	"errors"
	"fmt"
	"unicode"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r *bufio.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

func (s *Scanner) read() (rune, error) {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return rune(0), err
	}
	return ch, nil
}

func (s *Scanner) unread() error {
	if err := s.r.UnreadRune(); err != nil {
		return err
	}
	return nil
}

func (s *Scanner) Scan() (token.Token, string, error) {
	// Read the next rune.
	ch, err := s.read()
	fmt.Println(string(ch))
	if err != nil {
		return token.EOF, "", err
	}

	// branch off to more specific tokens
	if unicode.IsSpace(ch) {
		s.unread()
		return s.scanSpace()
	} else if IsSymbol(ch) {
		s.unread()
		return s.scanSymbol()
	}

	return token.ILLEGAL, "", errors.New("invalid token")
}
