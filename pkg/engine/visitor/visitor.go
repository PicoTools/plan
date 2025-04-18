package visitor

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/PicoTools/plan/pkg/constants"
	"github.com/PicoTools/plan/pkg/engine/builtin"
	planerrors "github.com/PicoTools/plan/pkg/engine/errors"
	"github.com/PicoTools/plan/pkg/engine/object"
	"github.com/PicoTools/plan/pkg/engine/scope"
	"github.com/PicoTools/plan/pkg/engine/storage"
	"github.com/PicoTools/plan/pkg/engine/types"
	"github.com/PicoTools/plan/pkg/engine/utils"
	"github.com/PicoTools/plan/pkg/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/go-faster/errors"
)

// NewVisitor returns new Visitor object
func NewVisitor() *Visitor {
	return &Visitor{
		file:   "<entrypoint>",
		depth:  0,
		parent: nil,
		err:    nil,
	}
}

// Visitor implements BasePLANVisitor interface
type Visitor struct {
	parser.BasePLANVisitor

	// parent visitor object
	// used for includes to process new context
	parent *Visitor
	// depth of visitors
	depth int
	// name of evaluated file (ommited if entrypoint)
	file string
	// line number of executed script
	line int
	// visitor error
	err error
}

// SetError saves error with current execution line (and file)
func (v *Visitor) SetError(err error) {
	// TODO: try to change model of nested visitors
	t := v
	for {
		if t.parent == nil {
			if v.file != "" {
				t.err = errors.Wrap(err, fmt.Sprintf("%s: line %d", v.file, v.line))
			} else {
				t.err = errors.Wrap(err, fmt.Sprintf("line %d", v.line))
			}
			return
		}
		t = t.parent
	}
}

// GetError returns error
func (v *Visitor) GetError() error {
	return v.err
}

// GetLine returns current line
func (v *Visitor) GetLine() int {
	return v.line
}

// GetFile returns current evaluated file
func (v *Visitor) GetFile() string {
	return v.file
}

// SetFile sets name of evaluated file
func (v *Visitor) SetFile(name string) {
	v.file = name
}

// UnsetFile unsets filename used for previous evaluation
func (v *Visitor) UnsetFile() {
	v.file = ""
}

// Visit visits AST and evaluate statements
func (v *Visitor) Visit(tree antlr.ParseTree) any {
	if retValue != nil {
		// exit from AST in case of return value presents
		return types.Success
	}

	// set line number
	v.line = tree.(antlr.ParserRuleContext).GetStart().GetLine()

	// exit in case of max visitor depth
	if v.depth > constants.MaxVisitorDepth {
		v.SetError(planerrors.ErrMaxVisitorDepth)
		return types.Failure
	}

	// exit in case of max scope depth
	if scope.CurrentScope.Depth() > constants.MaxScopeDepth {
		v.SetError(planerrors.ErrMaxRecurseDepth)
		return types.Failure
	}

	switch val := tree.(type) {
	case *parser.ProgFileContext:
		return v.VisitProgFile(val)
	case *parser.StmtsContext:
		return v.VisitStmts(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.SimpleStmtContext:
		return v.VisitSimpleStmt(val)
	case *parser.CompoundStmtContext:
		return v.VisitCompoundStmt(val)
	case *parser.AssignRegularContext:
		return v.VisitAssignRegular(val)
	case *parser.ExpBoolContext:
		return v.VisitExpBool(val)
	case *parser.IdentifierFnInvokeContext:
		return v.VisitIdentifierFnInvoke(val)
	case *parser.ExpIdentifierContext:
		return v.VisitExpIdentifier(val)
	case *parser.ExpIntegerContext:
		return v.VisitExpInteger(val)
	case *parser.ExpIntegerHexContext:
		return v.VisitExpIntegerHex(val)
	case *parser.ExpNullContext:
		return v.VisitExpNull(val)
	case *parser.ExpFloatContext:
		return v.VisitExpFloat(val)
	case *parser.ExpStringContext:
		return v.VisitExpString(val)
	case *parser.ExpListContext:
		return v.VisitExpList(val)
	case *parser.ListContext:
		return v.VisitList(val)
	case *parser.ExpDictContext:
		return v.VisitExpDict(val)
	case *parser.ExpLogicalNotContext:
		return v.VisitExpLogicalNot(val)
	case *parser.ExpLogicalOrContext:
		return v.VisitExpLogicalOr(val)
	case *parser.ExpLogicalAndContext:
		return v.VisitExpLogicalAnd(val)
	case *parser.ExpParenthesesContext:
		return v.VisitExpParentheses(val)
	case *parser.ExpEqualContext:
		return v.VisitExpEqual(val)
	case *parser.ExpNegContext:
		return v.VisitExpNeg(val)
	case *parser.ExpComparisonContext:
		return v.VisitExpComparison(val)
	case *parser.ExpSumSubContext:
		return v.VisitExpSumSub(val)
	case *parser.ExpPowContext:
		return v.VisitExpPow(val)
	case *parser.ExpMulDivModContext:
		return v.VisitExpMulDivMod(val)
	case *parser.ExpXorContext:
		return v.VisitExpXor(val)
	case *parser.ForStmtContext:
		return v.VisitForStmt(val)
	case *parser.BreakStmtContext:
		return v.VisitBreakStmt(val)
	case *parser.ContinueStmtContext:
		return v.VisitContinueStmt(val)
	case *parser.WhileStmtContext:
		return v.VisitWhileStmt(val)
	case *parser.IfStmtContext:
		return v.VisitIfStmt(val)
	case *parser.IfBlockStmtContext:
		return v.VisitIfBlockStmt(val)
	case *parser.ElseBlockStmtContext:
		return v.VisitElseBlockStmt(val)
	case *parser.ElifBlockStmtContext:
		return v.VisitElifBlockStmt(val)
	case *parser.ExpFnInvokeContext:
		return v.VisitExpFnInvoke(val)
	case *parser.ReturnStmtContext:
		return v.VisitReturnStmt(val)
	case *parser.ExpIdxContext:
		return v.VisitExpIdx(val)
	case *parser.AssignIdxsRegularContext:
		return v.VisitAssignIdxsRegular(val)
	case *parser.AssignIdxsSumContext:
		return v.VisitAssignIdxsSum(val)
	case *parser.AssignIdxsSubContext:
		return v.VisitAssignIdxsSub(val)
	case *parser.AssignIdxsMulContext:
		return v.VisitAssignIdxsMul(val)
	case *parser.AssignIdxsDivContext:
		return v.VisitAssignIdxsDiv(val)
	case *parser.AssignIdxsModContext:
		return v.VisitAssignIdxsMod(val)
	case *parser.AssignIdxsPowContext:
		return v.VisitAssignIdxsPow(val)
	case *parser.ExpCsInvokeContext:
		return v.VisitExpCsInvoke(val)
	case *parser.IdentifierCsInvokeContext:
		return v.VisitIdentifierCsInvoke(val)
	case *parser.IncludeStmtContext:
		return v.VisitIncludeStmt(val)
	case *parser.FnStmtContext:
		return v.VisitFnStmt(val)
	case *parser.FnBodyContext:
		return v.VisitFnBody(val)
	case *parser.AssignSumContext:
		return v.VisitAssignSum(val)
	case *parser.AssignSubContext:
		return v.VisitAssignSub(val)
	case *parser.AssignMulContext:
		return v.VisitAssignMul(val)
	case *parser.AssignDivContext:
		return v.VisitAssignDiv(val)
	case *parser.AssignModContext:
		return v.VisitAssignMod(val)
	case *parser.AssignPowContext:
		return v.VisitAssignPow(val)
	case *parser.ExpMethodInvokeContext:
		return v.VisitExpMethodInvoke(val)
	case *parser.IdentifierMethodInvokeContext:
		return v.VisitIdentifierMethodInvoke(val)
	case *parser.ExpCsContext:
		return v.VisitExpCs(val)
	case *parser.ClosureContext:
		return v.VisitClosure(val)
	case *parser.IdxContext:
		return v.VisitIdx(val)
	case *parser.DictContext:
		return v.VisitDict(val)
	default:
		v.SetError(fmt.Errorf("unknown context '%s' on visit", reflect.TypeOf(val).String()))
		return types.Failure
	}
}

func (v *Visitor) VisitProgFile(ctx *parser.ProgFileContext) any {
	// execute statements
	if ctx.Stmts() != nil {
		if ok := v.Visit(ctx.Stmts()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitStmts(ctx *parser.StmtsContext) any {
	for _, item := range ctx.AllStmt() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) any {
	// process simple statements
	if ctx.SimpleStmt() != nil {
		if ok := v.Visit(ctx.SimpleStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// process compound statements
	if ctx.CompoundStmt() != nil {
		if ok := v.Visit(ctx.CompoundStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}
	return types.Success
}

func (v *Visitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext) any {
	// assignment
	if ctx.Assignment() != nil {
		if ok := v.Visit(ctx.Assignment()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// resolve includes
	if ctx.IncludeStmt() != nil {
		if res := v.Visit(ctx.IncludeStmt()).(types.VisitResultType); res == types.Failure {
			return types.Failure
		}
	}

	// invoke function
	if ctx.FnInvoke() != nil {
		if res, ok := v.Visit(ctx.FnInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// invoke closure
	if ctx.CsInvoke() != nil {
		if res, ok := v.Visit(ctx.CsInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// invoke object's method
	if ctx.MethodInvoke() != nil {
		if res, ok := v.Visit(ctx.MethodInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// break
	if ctx.BreakStmt() != nil {
		if ok := v.Visit(ctx.BreakStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// continue
	if ctx.ContinueStmt() != nil {
		if ok := v.Visit(ctx.ContinueStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// return
	if ctx.ReturnStmt() != nil {
		if ok := v.Visit(ctx.ReturnStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitCompoundStmt(ctx *parser.CompoundStmtContext) any {
	// if/elif/else
	if ctx.IfStmt() != nil {
		scope.NewCurrentScope(false, false)
		// statements inside if block
		if ok := v.Visit(ctx.IfStmt()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// while
	if ctx.WhileStmt() != nil {
		scope.NewCurrentScope(false, true)
		// statements inside while block
		if ok := v.Visit(ctx.WhileStmt()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// for
	if ctx.ForStmt() != nil {
		scope.NewCurrentScope(false, true)
		// statements inside for block
		if ok := v.Visit(ctx.ForStmt()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// register native function
	if ctx.FnStmt() != nil {
		if ok := v.Visit(ctx.FnStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitAssignRegular(ctx *parser.AssignRegularContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}

	scope.CurrentScope.Put(ctx.GetName().GetText(), val)

	return types.Success
}

func (v *Visitor) VisitExpBool(ctx *parser.ExpBoolContext) any {
	switch ctx.GetText() {
	case "true":
		return object.NewBool(true)
	case "false":
		return object.NewBool(false)
	}
	v.SetError(fmt.Errorf("unable get 'bool' from '%s'", ctx.GetText()))
	return types.Failure
}

func (v *Visitor) VisitIdentifierFnInvoke(ctx *parser.IdentifierFnInvokeContext) any {
	var params []object.Object

	for _, item := range ctx.AllExp() {
		// function's arguments
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}

	return v.invokeFunc(ctx.GetName().GetText(), params...)
}

func (v *Visitor) VisitExpIdentifier(ctx *parser.ExpIdentifierContext) any {
	val := scope.CurrentScope.Get(ctx.GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetText()))
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpInteger(ctx *parser.ExpIntegerContext) any {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		v.SetError(fmt.Errorf("unable get 'int' from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewInt(val)
}

func (v *Visitor) VisitExpIntegerHex(ctx *parser.ExpIntegerHexContext) any {
	val, err := strconv.ParseInt(utils.StripHexPrefix(ctx.GetText()), 16, 64)
	if err != nil {
		v.SetError(fmt.Errorf("unable get 'int' from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewInt(val)
}

func (v *Visitor) VisitExpNull(_ *parser.ExpNullContext) any {
	return object.NewNull()
}

func (v *Visitor) VisitExpFloat(ctx *parser.ExpFloatContext) any {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		v.SetError(fmt.Errorf("unable get 'float' from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewFloat(val)
}

func (v *Visitor) VisitExpString(ctx *parser.ExpStringContext) any {
	val, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		v.SetError(fmt.Errorf("unable get 'str' from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewStr(val)
}

func (v *Visitor) VisitExpList(ctx *parser.ExpListContext) any {
	return v.Visit(ctx.List())
}

func (v *Visitor) VisitList(ctx *parser.ListContext) any {
	var list []object.Object
	for _, item := range ctx.AllExp() {
		val, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		list = append(list, val)
	}
	return object.NewList(list)
}

func (v *Visitor) VisitExpDict(ctx *parser.ExpDictContext) any {
	dict, ok := v.Visit(ctx.Dict()).(map[string]object.Object)
	if !ok {
		return types.Failure
	}
	return object.NewDict(dict)
}

func (v *Visitor) VisitDict(ctx *parser.DictContext) any {
	dict := make(map[string]object.Object)
	for _, item := range ctx.AllDictUnit() {
		// get dict's key
		key, ok := v.Visit(item.AllExp()[0]).(object.Object)
		if !ok {
			return types.Failure
		}

		// check if key is str
		var keyStr *object.Str
		switch obj := key.(type) {
		case *object.Str:
			keyStr = obj
		default:
			v.SetError(fmt.Errorf("key of dict must be 'str'"))
			return types.Failure
		}

		// check if duplicates exist
		if _, ok = dict[keyStr.Value()]; ok {
			v.SetError(fmt.Errorf("duplicated key '%s' in dict", keyStr.Value()))
			return types.Failure
		}

		// get value
		val, ok := v.Visit(item.AllExp()[1]).(object.Object)
		if !ok {
			return types.Failure
		}
		dict[keyStr.Value()] = val
	}
	return dict
}

func (v *Visitor) VisitExpLogicalNot(ctx *parser.ExpLogicalNotContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := val.UnaryOp(parser.PLANLexerNot)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: !%s", val.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpLogicalOr(ctx *parser.ExpLogicalOrContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	// fast exit hack to return value in case of left expression is true
	// true || X == true, where X is undefined
	switch value := lhs.(type) {
	case *object.Bool:
		if value.Value() {
			return lhs
		}
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.PLANLexerOr, rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s || %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpLogicalAnd(ctx *parser.ExpLogicalAndContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	// fast exit hack to return value in case of left expression is false
	// false && X == false, where X is undefined
	switch value := lhs.(type) {
	case *object.Bool:
		if !value.Value() {
			return lhs
		}
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.PLANLexerAnd, rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s && %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpParentheses(ctx *parser.ExpParenthesesContext) any {
	return v.Visit(ctx.Exp())
}

func (v *Visitor) VisitExpEqual(ctx *parser.ExpEqualContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpNeg(ctx *parser.ExpNegContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := val.UnaryOp(parser.PLANLexerSubtract)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: -%s", val.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpComparison(ctx *parser.ExpComparisonContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpSumSub(ctx *parser.ExpSumSubContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpPow(ctx *parser.ExpPowContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.PLANLexerPow, rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s ** %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpMulDivMod(ctx *parser.ExpMulDivModContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpXor(ctx *parser.ExpXorContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.PLANLexerXor, rhs)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s ^ %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitForStmt(ctx *parser.ForStmtContext) any {
	// initial counter value
	if ok := v.Visit(ctx.Assignment(0)).(types.VisitResultType); !ok {
		return types.Failure
	}

	for {
		// loop exit statement value
		res, ok := v.Visit(ctx.Exp()).(object.Object)
		if !ok {
			if v.GetError() != nil {
				v.SetError(fmt.Errorf("invalid expression for loop: %s", v.GetError()))
			} else {
				v.SetError(fmt.Errorf("invalid expression for loop"))
			}
			return types.Failure
		}

		if _, ok = res.(*object.Bool); !ok {
			v.SetError(fmt.Errorf("non-bool (%s) expression in for loop", res.TypeName()))
			return types.Failure
		}

		// false -> exit from loop
		// TODO: handle type is Bool
		if !res.(*object.Bool).GetValue().(bool) {
			break
		}

		for _, item := range ctx.AllStmt() {
			// break
			if scope.CurrentScope.IsBreak() {
				scope.CurrentScope.UnsetLoopBreak()
				return types.Success
			}
			// continue
			if scope.CurrentScope.IsContinue() {
				scope.CurrentScope.UnsetLoopContinue()
				break
			}
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
			if retValue != nil {
				return types.Success
			}
		}

		// update counter
		if ok := v.Visit(ctx.Assignment(1)).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) any {
	if scope.CurrentScope.IsInLoop() {
		scope.CurrentScope.SetLoopBreak()
		return types.Success
	}
	v.SetError(fmt.Errorf("break outside of loop"))
	return types.Failure
}

func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) any {
	if scope.CurrentScope.IsInLoop() {
		scope.CurrentScope.SetLoopContinue()
		return types.Success
	}
	v.SetError(fmt.Errorf("continue outside of loop"))
	return types.Failure
}

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) any {
	for {
		// loop exit statement value
		res, ok := v.Visit(ctx.Exp()).(object.Object)
		if !ok {
			return types.Failure
		}

		if _, ok = res.(*object.Bool); !ok {
			v.SetError(fmt.Errorf("non-bool (%s) expression in for loop", res.TypeName()))
			return types.Failure
		}

		// false -> exit loop
		// TODO: handle type is Bool
		if !res.(*object.Bool).GetValue().(bool) {
			break
		}

		for _, item := range ctx.AllStmt() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
			// break
			if scope.CurrentScope.IsBreak() {
				scope.CurrentScope.UnsetLoopBreak()
				return types.Success
			}
			// continue
			if scope.CurrentScope.IsContinue() {
				scope.CurrentScope.UnsetLoopContinue()
				break
			}
			if retValue != nil {
				return types.Success
			}
		}
	}

	return types.Success
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) any {
	// if
	if ctx.IfBlock() != nil {
		// if block result
		res, ok := v.Visit(ctx.IfBlock()).(object.Object)
		if !ok {
			return types.Failure
		}
		// TODO: handle type is Bool
		if res.(*object.Bool).GetValue().(bool) {
			return types.Success
		}
	}

	// elif
	for _, item := range ctx.AllElifBlock() {
		// elif block result
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		// TODO: handle type is Bool
		if res.(*object.Bool).GetValue().(bool) {
			return types.Success
		}
	}

	// else
	if ctx.ElseBlock() != nil {
		if ok := v.Visit(ctx.ElseBlock()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitIfBlockStmt(ctx *parser.IfBlockStmtContext) any {
	res, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}

	res, err := builtin.Bool(res)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}

	// TODO: handle type is Bool
	if res.(*object.Bool).GetValue().(bool) {
		// true -> execute block's statements
		for _, item := range ctx.AllStmt() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
		}
	}

	return res
}

func (v *Visitor) VisitElseBlockStmt(ctx *parser.ElseBlockStmtContext) any {
	for _, item := range ctx.AllStmt() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}
	return types.Success
}

func (v *Visitor) VisitElifBlockStmt(ctx *parser.ElifBlockStmtContext) any {
	res, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}

	res, err := builtin.Bool(res)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}

	// TODO: handle type is Bool
	if res.(*object.Bool).GetValue().(bool) {
		// true -> execute block's statements
		for _, item := range ctx.AllStmt() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
		}
	}

	return res
}

func (v *Visitor) VisitFnStmt(ctx *parser.FnStmtContext) any {
	// check if function already exists in builtin
	if val := storage.GetFunction(ctx.GetName().GetText()); val != nil {
		v.SetError(fmt.Errorf("function '%s' already defined", ctx.GetName().GetText()))
		return types.Failure
	}

	// check if function already exists in scope
	if _, ok := scope.CurrentScope.Functions()[ctx.GetName().GetText()]; ok {
		v.SetError(fmt.Errorf("function '%s' already defined", ctx.GetName().GetText()))
		return types.Failure
	}

	// get function body
	fn, ok := v.Visit(ctx.FnBody()).(*object.RuntimeFunc)
	if !ok {
		return types.Failure
	}

	// set function name
	fn.SetName(ctx.GetName().GetText())

	// save native function
	scope.CurrentScope.Functions()[ctx.GetName().GetText()] = fn

	return types.Success
}

func (v *Visitor) VisitFnBody(ctx *parser.FnBodyContext) any {
	// function's arguments
	args := make([]string, 0)
	if ctx.FnParams() != nil {
		for _, item := range ctx.FnParams().AllIdentifier() {
			args = append(args, item.GetText())
		}
	}
	return object.NewRuntimeFunc(args, ctx.AllStmt())
}

func (v *Visitor) VisitExpFnInvoke(ctx *parser.ExpFnInvokeContext) any {
	res, ok := v.Visit(ctx.FnInvoke()).(object.Object)
	if !ok {
		if v.GetError() == nil {
			return object.NewNull()
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) any {
	if !scope.CurrentScope.IsInFunc() {
		v.SetError(fmt.Errorf("return outside of function"))
		return types.Failure
	}
	// return value
	if ctx.Exp() != nil {
		res, ok := v.Visit(ctx.Exp()).(object.Object)
		if !ok {
			return types.Failure
		}
		retValue = res
		return types.Success
	} else {
		retValue = object.NewNull()
		return types.Success
	}
}

func (v *Visitor) VisitExpIdx(ctx *parser.ExpIdxContext) any {
	lhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	idx, ok := v.Visit(ctx.Idx()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := lhs.IndexGet(idx)
	if err != nil {
		if errors.Is(err, planerrors.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s[%s]", lhs.TypeName(), idx.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitIdx(ctx *parser.IdxContext) any {
	return v.Visit(ctx.Exp())
}

func (v *Visitor) VisitAssignIdxsRegular(ctx *parser.AssignIdxsRegularContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, val); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignIdxsSum(ctx *parser.AssignIdxsSumContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value by index
	val, err := lhs.IndexGet(lastIdx)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// get result of rhs expression
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// get result of expression
	res, err := val.BinaryOp(parser.PLANLexerAssSum, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, res); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignIdxsSub(ctx *parser.AssignIdxsSubContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value by index
	val, err := lhs.IndexGet(lastIdx)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// get result of rhs expression
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// get result of expression
	res, err := val.BinaryOp(parser.PLANLexerAssSub, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, res); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignIdxsMul(ctx *parser.AssignIdxsMulContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value by index
	val, err := lhs.IndexGet(lastIdx)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// get result of rhs expression
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// get result of expression
	res, err := val.BinaryOp(parser.PLANLexerAssMul, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, res); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignIdxsDiv(ctx *parser.AssignIdxsDivContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value by index
	val, err := lhs.IndexGet(lastIdx)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// get result of rhs expression
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// get result of expression
	res, err := val.BinaryOp(parser.PLANLexerAssDiv, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, res); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignIdxsMod(ctx *parser.AssignIdxsModContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value by index
	val, err := lhs.IndexGet(lastIdx)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// get result of rhs expression
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// get result of expression
	res, err := val.BinaryOp(parser.PLANLexerAssMod, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, res); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignIdxsPow(ctx *parser.AssignIdxsPowContext) any {
	// get variable from scope
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	var lastIdx object.Object
	// get all nested indexes
	for i, idxStmt := range ctx.Idxs().AllIdx() {
		idx, ok := v.Visit(idxStmt).(object.Object)
		if !ok {
			return types.Failure
		}
		if i == len(ctx.Idxs().AllIdx())-1 {
			lastIdx = idx
			break
		}
		var err error
		lhs, err = lhs.IndexGet(idx)
		if err != nil {
			v.SetError(err)
			return types.Failure
		}
	}
	// get value by index
	val, err := lhs.IndexGet(lastIdx)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// get result of rhs expression
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// get result of expression
	res, err := val.BinaryOp(parser.PLANLexerAssPow, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update value by index in variable
	if err := lhs.IndexSet(lastIdx, res); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitExpCsInvoke(ctx *parser.ExpCsInvokeContext) any {
	res, ok := v.Visit(ctx.CsInvoke()).(object.Object)
	if !ok {
		if v.GetError() == nil {
			return object.NewNull()
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitIdentifierCsInvoke(ctx *parser.IdentifierCsInvokeContext) any {
	var params []object.Object

	for _, item := range ctx.AllExp() {
		// arguments for closure invoke
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}

	return v.invokeClosure(ctx.GetName().GetText(), params...)
}

func (v *Visitor) VisitIncludeStmt(ctx *parser.IncludeStmtContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	path, ok := val.(*object.Str)
	if !ok {
		v.SetError(fmt.Errorf("invalid argument of type '%s'", val.TypeName()))
		return types.Failure
	}
	data, err := os.ReadFile(path.Value())
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	tree, err := utils.CreateAstProgFile(string(data))
	if err != nil {
		v.SetError(err)
		return types.Failure
	}

	// update visitor context
	v = &Visitor{
		parent: v,
		err:    nil,
		file:   path.Value(),
		depth:  v.depth + 1,
	}
	// traverse AST
	if res := v.Visit(tree); res != types.Success {
		return types.Failure
	}
	// revert parent visitor
	v = v.parent
	return types.Success
}

func (v *Visitor) VisitExpCs(ctx *parser.ExpCsContext) any {
	return v.Visit(ctx.Closure())
}

func (v *Visitor) VisitClosure(ctx *parser.ClosureContext) interface{} {
	return v.Visit(ctx.FnBody())
}

func (v *Visitor) VisitAssignSum(ctx *parser.AssignSumContext) any {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := val.BinaryOp(parser.PLANLexerAssSum, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update variable in scope
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitAssignSub(ctx *parser.AssignSubContext) any {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := val.BinaryOp(parser.PLANLexerAssSub, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update variable in scope
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitAssignMul(ctx *parser.AssignMulContext) any {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := val.BinaryOp(parser.PLANLexerAssMul, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update variable in scope
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitAssignDiv(ctx *parser.AssignDivContext) any {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := val.BinaryOp(parser.PLANLexerAssDiv, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update variable in scope
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitAssignMod(ctx *parser.AssignModContext) any {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := val.BinaryOp(parser.PLANLexerAssMod, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update variable in scope
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitAssignPow(ctx *parser.AssignPowContext) any {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := val.BinaryOp(parser.PLANLexerAssPow, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// update variable in scope
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitExpMethodInvoke(ctx *parser.ExpMethodInvokeContext) interface{} {
	return v.Visit(ctx.MethodInvoke())
}

func (v *Visitor) VisitIdentifierMethodInvoke(ctx *parser.IdentifierMethodInvokeContext) interface{} {
	// variable from scope
	val := scope.CurrentScope.Get(ctx.GetVar_().GetText(), false)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetVar_().GetText()))
		return types.Failure
	}
	// method's arguments
	var params []object.Object
	for _, item := range ctx.AllExp() {
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}
	// invoke method
	obj, err := val.MethodCall(ctx.GetName().GetText(), params...)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	return obj
}
