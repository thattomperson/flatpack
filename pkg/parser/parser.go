package parser // import "ttp.sh/flatpack/pkg/parser"
import (
	"ttp.sh/flatpack/pkg/token"

	"go/scanner"
)

// The parser structure holds the parser's internal state.
type parser struct {
	file    *token.File
	errors  scanner.ErrorList
	scanner scanner.Scanner

	// Next token
	pos token.Pos   // token position
	tok token.Token // one token look-ahead
	lit string      // token literal
}
