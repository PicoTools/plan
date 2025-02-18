// Code generated from ./grammar/PLAN.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PLAN
import "github.com/antlr4-go/antlr/v4"

type BasePLANVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePLANVisitor) VisitProgFile(ctx *ProgFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitStmts(ctx *StmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitCompoundStmt(ctx *CompoundStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIfBlockStmt(ctx *IfBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitElifBlockStmt(ctx *ElifBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitElseBlockStmt(ctx *ElseBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitFnStmt(ctx *FnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitFnBody(ctx *FnBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitFnParams(ctx *FnParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIncludeStmt(ctx *IncludeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignRegular(ctx *AssignRegularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignSum(ctx *AssignSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignSub(ctx *AssignSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignMul(ctx *AssignMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignDiv(ctx *AssignDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignMod(ctx *AssignModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignPow(ctx *AssignPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsRegular(ctx *AssignIdxsRegularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsSum(ctx *AssignIdxsSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsSub(ctx *AssignIdxsSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsMul(ctx *AssignIdxsMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsDiv(ctx *AssignIdxsDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsMod(ctx *AssignIdxsModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitAssignIdxsPow(ctx *AssignIdxsPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitDictUnit(ctx *DictUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitDict(ctx *DictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIdx(ctx *IdxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIdxs(ctx *IdxsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpBool(ctx *ExpBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpComparison(ctx *ExpComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpIdx(ctx *ExpIdxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpString(ctx *ExpStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpCsInvoke(ctx *ExpCsInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpFloat(ctx *ExpFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpPow(ctx *ExpPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpDict(ctx *ExpDictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpXor(ctx *ExpXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpNeg(ctx *ExpNegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpInteger(ctx *ExpIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpLogicalOr(ctx *ExpLogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpCs(ctx *ExpCsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpMulDivMod(ctx *ExpMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpNull(ctx *ExpNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpFnInvoke(ctx *ExpFnInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpList(ctx *ExpListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpLogicalAnd(ctx *ExpLogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpParentheses(ctx *ExpParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpEqual(ctx *ExpEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpMethodInvoke(ctx *ExpMethodInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpLogicalNot(ctx *ExpLogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpIntegerHex(ctx *ExpIntegerHexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpIdentifier(ctx *ExpIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitExpSumSub(ctx *ExpSumSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePLANVisitor) VisitClosure(ctx *ClosureContext) interface{} {
	return v.VisitChildren(ctx)
}
