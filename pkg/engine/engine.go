package engine

import (
	"os"

	"github.com/PicoTools/plan/pkg/engine/builtin"
	"github.com/PicoTools/plan/pkg/engine/scope"
	"github.com/PicoTools/plan/pkg/engine/storage"
	"github.com/PicoTools/plan/pkg/engine/types"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/engine/visitor"
	"github.com/go-faster/errors"
)

func init() {
	Init()
}

// Init initializes PLAN runtime
func Init() {
	scope.NewGlobalScope()
	builtin.Register()
}

// Reset clears PLAN runtime and reinit it
func Reset() {
	storage.ResetStorage()
	visitor.ClearRet()
	Init()
}

// Evaulate execute script by path with PLAN runtime
func Evaluate(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	tree, err := utils.CreateAstProgFile(string(data))
	if err != nil {
		return err
	}

	v := visitor.NewVisitor()
	res := v.Visit(tree)
	if res != types.Success {
		return errors.Wrap(v.GetError(), "runtime failure")
	}

	return nil
}
