package utils

import (
	"github.com/PicoTools/plan/pkg/parser"
	"github.com/PicoTools/plan/pkg/perror"
	"github.com/antlr4-go/antlr/v4"
)

// CreateAstProgFile creates AST from string data based on progFile entrypoint
func CreateAstProgFile(data string) (antlr.ParseTree, error) {
	iStream := antlr.NewInputStream(data)
	lexer := parser.NewPLANLexer(iStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPLANParser(stream)
	errorParser := &perror.Error{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorParser)
	tree := p.ProgFile()
	return tree, errorParser.Err
}
