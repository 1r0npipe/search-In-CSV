package parser

import (
	"bufio"
	"io"
)

const (
	// Special tokens
	IDENT = iota //
	SELECT
	FROM
	WHERE
	AND
	OR
)
type Statement struct {
	Fields []string
	fileName string
}
type symbol string

const (
	equal    symbol = "="
	asterisk symbol = "*"
	comma    symbol = ","
	lParen   symbol = "("
	rParen   symbol = ")"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

