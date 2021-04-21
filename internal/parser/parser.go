package parser

import (
	"errors"
	"unicode"
)

var (
	EOF = rune(0)
	ErrScanRune = errors.New("can't read rune")
)

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (s *Scanner) read() rune {
	ch, _ , err := s.r.ReadRune()
	if err != nil {
		return EOF
	}
	return ch
}
