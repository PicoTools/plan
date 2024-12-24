// Code generated from ./grammar/PLAN.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PLAN
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PLANParser.
type PLANVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PLANParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by PLANParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by PLANParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#assignRegular.
	VisitAssignRegular(ctx *AssignRegularContext) interface{}

	// Visit a parse tree produced by PLANParser#assignSum.
	VisitAssignSum(ctx *AssignSumContext) interface{}

	// Visit a parse tree produced by PLANParser#assignSub.
	VisitAssignSub(ctx *AssignSubContext) interface{}

	// Visit a parse tree produced by PLANParser#assignMul.
	VisitAssignMul(ctx *AssignMulContext) interface{}

	// Visit a parse tree produced by PLANParser#assignDiv.
	VisitAssignDiv(ctx *AssignDivContext) interface{}

	// Visit a parse tree produced by PLANParser#assignMod.
	VisitAssignMod(ctx *AssignModContext) interface{}

	// Visit a parse tree produced by PLANParser#assignPow.
	VisitAssignPow(ctx *AssignPowContext) interface{}

	// Visit a parse tree produced by PLANParser#assignIdxRegular.
	VisitAssignIdxRegular(ctx *AssignIdxRegularContext) interface{}

	// Visit a parse tree produced by PLANParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by PLANParser#dictUnit.
	VisitDictUnit(ctx *DictUnitContext) interface{}

	// Visit a parse tree produced by PLANParser#dict.
	VisitDict(ctx *DictContext) interface{}

	// Visit a parse tree produced by PLANParser#idx.
	VisitIdx(ctx *IdxContext) interface{}

	// Visit a parse tree produced by PLANParser#identifierMethodInvoke.
	VisitIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) interface{}

	// Visit a parse tree produced by PLANParser#identifierFnInvoke.
	VisitIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) interface{}

	// Visit a parse tree produced by PLANParser#identifierCsInvoke.
	VisitIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) interface{}

	// Visit a parse tree produced by PLANParser#expBool.
	VisitExpBool(ctx *ExpBoolContext) interface{}

	// Visit a parse tree produced by PLANParser#expComparison.
	VisitExpComparison(ctx *ExpComparisonContext) interface{}

	// Visit a parse tree produced by PLANParser#expIdx.
	VisitExpIdx(ctx *ExpIdxContext) interface{}

	// Visit a parse tree produced by PLANParser#expString.
	VisitExpString(ctx *ExpStringContext) interface{}

	// Visit a parse tree produced by PLANParser#expCsInvoke.
	VisitExpCsInvoke(ctx *ExpCsInvokeContext) interface{}

	// Visit a parse tree produced by PLANParser#expFloat.
	VisitExpFloat(ctx *ExpFloatContext) interface{}

	// Visit a parse tree produced by PLANParser#expPow.
	VisitExpPow(ctx *ExpPowContext) interface{}

	// Visit a parse tree produced by PLANParser#expDict.
	VisitExpDict(ctx *ExpDictContext) interface{}

	// Visit a parse tree produced by PLANParser#expXor.
	VisitExpXor(ctx *ExpXorContext) interface{}

	// Visit a parse tree produced by PLANParser#expNeg.
	VisitExpNeg(ctx *ExpNegContext) interface{}

	// Visit a parse tree produced by PLANParser#expInteger.
	VisitExpInteger(ctx *ExpIntegerContext) interface{}

	// Visit a parse tree produced by PLANParser#expLogicalOr.
	VisitExpLogicalOr(ctx *ExpLogicalOrContext) interface{}

	// Visit a parse tree produced by PLANParser#expCs.
	VisitExpCs(ctx *ExpCsContext) interface{}

	// Visit a parse tree produced by PLANParser#expMulDivMod.
	VisitExpMulDivMod(ctx *ExpMulDivModContext) interface{}

	// Visit a parse tree produced by PLANParser#expNull.
	VisitExpNull(ctx *ExpNullContext) interface{}

	// Visit a parse tree produced by PLANParser#expFnInvoke.
	VisitExpFnInvoke(ctx *ExpFnInvokeContext) interface{}

	// Visit a parse tree produced by PLANParser#expList.
	VisitExpList(ctx *ExpListContext) interface{}

	// Visit a parse tree produced by PLANParser#expLogicalAnd.
	VisitExpLogicalAnd(ctx *ExpLogicalAndContext) interface{}

	// Visit a parse tree produced by PLANParser#expParentheses.
	VisitExpParentheses(ctx *ExpParenthesesContext) interface{}

	// Visit a parse tree produced by PLANParser#expEqual.
	VisitExpEqual(ctx *ExpEqualContext) interface{}

	// Visit a parse tree produced by PLANParser#expMethodInvoke.
	VisitExpMethodInvoke(ctx *ExpMethodInvokeContext) interface{}

	// Visit a parse tree produced by PLANParser#expLogicalNot.
	VisitExpLogicalNot(ctx *ExpLogicalNotContext) interface{}

	// Visit a parse tree produced by PLANParser#expIntegerHex.
	VisitExpIntegerHex(ctx *ExpIntegerHexContext) interface{}

	// Visit a parse tree produced by PLANParser#expIdentifier.
	VisitExpIdentifier(ctx *ExpIdentifierContext) interface{}

	// Visit a parse tree produced by PLANParser#expSumSub.
	VisitExpSumSub(ctx *ExpSumSubContext) interface{}

	// Visit a parse tree produced by PLANParser#ifBlockStmt.
	VisitIfBlockStmt(ctx *IfBlockStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#elifBlockStmt.
	VisitElifBlockStmt(ctx *ElifBlockStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#elseBlockStmt.
	VisitElseBlockStmt(ctx *ElseBlockStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by PLANParser#fnParams.
	VisitFnParams(ctx *FnParamsContext) interface{}

	// Visit a parse tree produced by PLANParser#fnBody.
	VisitFnBody(ctx *FnBodyContext) interface{}

	// Visit a parse tree produced by PLANParser#fn.
	VisitFn(ctx *FnContext) interface{}

	// Visit a parse tree produced by PLANParser#closure.
	VisitClosure(ctx *ClosureContext) interface{}

	// Visit a parse tree produced by PLANParser#include.
	VisitInclude(ctx *IncludeContext) interface{}
}
