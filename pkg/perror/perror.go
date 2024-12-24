package perror

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Error implements ErrorListener interface
type Error struct {
	antlr.ErrorListener
	Line   int
	Column int
	Err    error
}

// SyntaxError indicates error in syntax
func (c *Error) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	c.Column = column
	c.Line = line
	c.Err = fmt.Errorf("%d:%d %s", line, column, msg)
}

func (c *Error) ReportAttemptingFullContext(antlr.Parser, *antlr.DFA, int, int, *antlr.BitSet, *antlr.ATNConfigSet) {
}

func (c *Error) ReportContextSensitivity(antlr.Parser, *antlr.DFA, int, int, int, *antlr.ATNConfigSet) {
}

func (c *Error) ReportAmbiguity(antlr.Parser, *antlr.DFA, int, int, bool, *antlr.BitSet, *antlr.ATNConfigSet) {
}
