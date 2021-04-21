package parser

const (
	// Special tokens
	IDENT = iota //
	SELECT
	FROM
	WHERE
	AND
	OR
)
type SelectStatement struct {
	Fields []string
	fileName string
}
type symbol string

const (
	equal    symbol = "=="
	asterisk symbol = "*"
	comma    symbol = ","
	lParen   symbol = "("
	rParen   symbol = ")"
)

