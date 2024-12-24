package utils

import (
	"github.com/PicoTools/plan/pkg/parser"
	"github.com/PicoTools/plan/pkg/perror"
	"github.com/antlr4-go/antlr/v4"
)

// CreateAST creates AST from string data
func CreateAST(data string) (antlr.ParseTree, error) {
	iStream := antlr.NewInputStream(data)
	lexer := parser.NewPLANLexer(iStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPLANParser(stream)
	errorParser := &perror.Error{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorParser)
	tree := p.Prog()
	return tree, errorParser.Err
}
