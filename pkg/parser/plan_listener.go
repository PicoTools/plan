// Code generated from ./grammar/PLAN.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PLAN
import "github.com/antlr4-go/antlr/v4"

// PLANListener is a complete listener for a parse tree produced by PLANParser.
type PLANListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterAssignRegular is called when entering the assignRegular production.
	EnterAssignRegular(c *AssignRegularContext)

	// EnterAssignSum is called when entering the assignSum production.
	EnterAssignSum(c *AssignSumContext)

	// EnterAssignSub is called when entering the assignSub production.
	EnterAssignSub(c *AssignSubContext)

	// EnterAssignMul is called when entering the assignMul production.
	EnterAssignMul(c *AssignMulContext)

	// EnterAssignDiv is called when entering the assignDiv production.
	EnterAssignDiv(c *AssignDivContext)

	// EnterAssignMod is called when entering the assignMod production.
	EnterAssignMod(c *AssignModContext)

	// EnterAssignPow is called when entering the assignPow production.
	EnterAssignPow(c *AssignPowContext)

	// EnterAssignIdxRegular is called when entering the assignIdxRegular production.
	EnterAssignIdxRegular(c *AssignIdxRegularContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterDictUnit is called when entering the dictUnit production.
	EnterDictUnit(c *DictUnitContext)

	// EnterDict is called when entering the dict production.
	EnterDict(c *DictContext)

	// EnterIdx is called when entering the idx production.
	EnterIdx(c *IdxContext)

	// EnterIdentifierMethodInvoke is called when entering the identifierMethodInvoke production.
	EnterIdentifierMethodInvoke(c *IdentifierMethodInvokeContext)

	// EnterIdentifierFnInvoke is called when entering the identifierFnInvoke production.
	EnterIdentifierFnInvoke(c *IdentifierFnInvokeContext)

	// EnterIdentifierCsInvoke is called when entering the identifierCsInvoke production.
	EnterIdentifierCsInvoke(c *IdentifierCsInvokeContext)

	// EnterExpBool is called when entering the expBool production.
	EnterExpBool(c *ExpBoolContext)

	// EnterExpComparison is called when entering the expComparison production.
	EnterExpComparison(c *ExpComparisonContext)

	// EnterExpIdx is called when entering the expIdx production.
	EnterExpIdx(c *ExpIdxContext)

	// EnterExpString is called when entering the expString production.
	EnterExpString(c *ExpStringContext)

	// EnterExpCsInvoke is called when entering the expCsInvoke production.
	EnterExpCsInvoke(c *ExpCsInvokeContext)

	// EnterExpFloat is called when entering the expFloat production.
	EnterExpFloat(c *ExpFloatContext)

	// EnterExpPow is called when entering the expPow production.
	EnterExpPow(c *ExpPowContext)

	// EnterExpDict is called when entering the expDict production.
	EnterExpDict(c *ExpDictContext)

	// EnterExpXor is called when entering the expXor production.
	EnterExpXor(c *ExpXorContext)

	// EnterExpNeg is called when entering the expNeg production.
	EnterExpNeg(c *ExpNegContext)

	// EnterExpInteger is called when entering the expInteger production.
	EnterExpInteger(c *ExpIntegerContext)

	// EnterExpLogicalOr is called when entering the expLogicalOr production.
	EnterExpLogicalOr(c *ExpLogicalOrContext)

	// EnterExpCs is called when entering the expCs production.
	EnterExpCs(c *ExpCsContext)

	// EnterExpMulDivMod is called when entering the expMulDivMod production.
	EnterExpMulDivMod(c *ExpMulDivModContext)

	// EnterExpNull is called when entering the expNull production.
	EnterExpNull(c *ExpNullContext)

	// EnterExpFnInvoke is called when entering the expFnInvoke production.
	EnterExpFnInvoke(c *ExpFnInvokeContext)

	// EnterExpList is called when entering the expList production.
	EnterExpList(c *ExpListContext)

	// EnterExpLogicalAnd is called when entering the expLogicalAnd production.
	EnterExpLogicalAnd(c *ExpLogicalAndContext)

	// EnterExpParentheses is called when entering the expParentheses production.
	EnterExpParentheses(c *ExpParenthesesContext)

	// EnterExpEqual is called when entering the expEqual production.
	EnterExpEqual(c *ExpEqualContext)

	// EnterExpMethodInvoke is called when entering the expMethodInvoke production.
	EnterExpMethodInvoke(c *ExpMethodInvokeContext)

	// EnterExpLogicalNot is called when entering the expLogicalNot production.
	EnterExpLogicalNot(c *ExpLogicalNotContext)

	// EnterExpIntegerHex is called when entering the expIntegerHex production.
	EnterExpIntegerHex(c *ExpIntegerHexContext)

	// EnterExpIdentifier is called when entering the expIdentifier production.
	EnterExpIdentifier(c *ExpIdentifierContext)

	// EnterExpSumSub is called when entering the expSumSub production.
	EnterExpSumSub(c *ExpSumSubContext)

	// EnterIfBlockStmt is called when entering the ifBlockStmt production.
	EnterIfBlockStmt(c *IfBlockStmtContext)

	// EnterElifBlockStmt is called when entering the elifBlockStmt production.
	EnterElifBlockStmt(c *ElifBlockStmtContext)

	// EnterElseBlockStmt is called when entering the elseBlockStmt production.
	EnterElseBlockStmt(c *ElseBlockStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterFnParams is called when entering the fnParams production.
	EnterFnParams(c *FnParamsContext)

	// EnterFnBody is called when entering the fnBody production.
	EnterFnBody(c *FnBodyContext)

	// EnterFn is called when entering the fn production.
	EnterFn(c *FnContext)

	// EnterClosure is called when entering the closure production.
	EnterClosure(c *ClosureContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitAssignRegular is called when exiting the assignRegular production.
	ExitAssignRegular(c *AssignRegularContext)

	// ExitAssignSum is called when exiting the assignSum production.
	ExitAssignSum(c *AssignSumContext)

	// ExitAssignSub is called when exiting the assignSub production.
	ExitAssignSub(c *AssignSubContext)

	// ExitAssignMul is called when exiting the assignMul production.
	ExitAssignMul(c *AssignMulContext)

	// ExitAssignDiv is called when exiting the assignDiv production.
	ExitAssignDiv(c *AssignDivContext)

	// ExitAssignMod is called when exiting the assignMod production.
	ExitAssignMod(c *AssignModContext)

	// ExitAssignPow is called when exiting the assignPow production.
	ExitAssignPow(c *AssignPowContext)

	// ExitAssignIdxRegular is called when exiting the assignIdxRegular production.
	ExitAssignIdxRegular(c *AssignIdxRegularContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitDictUnit is called when exiting the dictUnit production.
	ExitDictUnit(c *DictUnitContext)

	// ExitDict is called when exiting the dict production.
	ExitDict(c *DictContext)

	// ExitIdx is called when exiting the idx production.
	ExitIdx(c *IdxContext)

	// ExitIdentifierMethodInvoke is called when exiting the identifierMethodInvoke production.
	ExitIdentifierMethodInvoke(c *IdentifierMethodInvokeContext)

	// ExitIdentifierFnInvoke is called when exiting the identifierFnInvoke production.
	ExitIdentifierFnInvoke(c *IdentifierFnInvokeContext)

	// ExitIdentifierCsInvoke is called when exiting the identifierCsInvoke production.
	ExitIdentifierCsInvoke(c *IdentifierCsInvokeContext)

	// ExitExpBool is called when exiting the expBool production.
	ExitExpBool(c *ExpBoolContext)

	// ExitExpComparison is called when exiting the expComparison production.
	ExitExpComparison(c *ExpComparisonContext)

	// ExitExpIdx is called when exiting the expIdx production.
	ExitExpIdx(c *ExpIdxContext)

	// ExitExpString is called when exiting the expString production.
	ExitExpString(c *ExpStringContext)

	// ExitExpCsInvoke is called when exiting the expCsInvoke production.
	ExitExpCsInvoke(c *ExpCsInvokeContext)

	// ExitExpFloat is called when exiting the expFloat production.
	ExitExpFloat(c *ExpFloatContext)

	// ExitExpPow is called when exiting the expPow production.
	ExitExpPow(c *ExpPowContext)

	// ExitExpDict is called when exiting the expDict production.
	ExitExpDict(c *ExpDictContext)

	// ExitExpXor is called when exiting the expXor production.
	ExitExpXor(c *ExpXorContext)

	// ExitExpNeg is called when exiting the expNeg production.
	ExitExpNeg(c *ExpNegContext)

	// ExitExpInteger is called when exiting the expInteger production.
	ExitExpInteger(c *ExpIntegerContext)

	// ExitExpLogicalOr is called when exiting the expLogicalOr production.
	ExitExpLogicalOr(c *ExpLogicalOrContext)

	// ExitExpCs is called when exiting the expCs production.
	ExitExpCs(c *ExpCsContext)

	// ExitExpMulDivMod is called when exiting the expMulDivMod production.
	ExitExpMulDivMod(c *ExpMulDivModContext)

	// ExitExpNull is called when exiting the expNull production.
	ExitExpNull(c *ExpNullContext)

	// ExitExpFnInvoke is called when exiting the expFnInvoke production.
	ExitExpFnInvoke(c *ExpFnInvokeContext)

	// ExitExpList is called when exiting the expList production.
	ExitExpList(c *ExpListContext)

	// ExitExpLogicalAnd is called when exiting the expLogicalAnd production.
	ExitExpLogicalAnd(c *ExpLogicalAndContext)

	// ExitExpParentheses is called when exiting the expParentheses production.
	ExitExpParentheses(c *ExpParenthesesContext)

	// ExitExpEqual is called when exiting the expEqual production.
	ExitExpEqual(c *ExpEqualContext)

	// ExitExpMethodInvoke is called when exiting the expMethodInvoke production.
	ExitExpMethodInvoke(c *ExpMethodInvokeContext)

	// ExitExpLogicalNot is called when exiting the expLogicalNot production.
	ExitExpLogicalNot(c *ExpLogicalNotContext)

	// ExitExpIntegerHex is called when exiting the expIntegerHex production.
	ExitExpIntegerHex(c *ExpIntegerHexContext)

	// ExitExpIdentifier is called when exiting the expIdentifier production.
	ExitExpIdentifier(c *ExpIdentifierContext)

	// ExitExpSumSub is called when exiting the expSumSub production.
	ExitExpSumSub(c *ExpSumSubContext)

	// ExitIfBlockStmt is called when exiting the ifBlockStmt production.
	ExitIfBlockStmt(c *IfBlockStmtContext)

	// ExitElifBlockStmt is called when exiting the elifBlockStmt production.
	ExitElifBlockStmt(c *ElifBlockStmtContext)

	// ExitElseBlockStmt is called when exiting the elseBlockStmt production.
	ExitElseBlockStmt(c *ElseBlockStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitFnParams is called when exiting the fnParams production.
	ExitFnParams(c *FnParamsContext)

	// ExitFnBody is called when exiting the fnBody production.
	ExitFnBody(c *FnBodyContext)

	// ExitFn is called when exiting the fn production.
	ExitFn(c *FnContext)

	// ExitClosure is called when exiting the closure production.
	ExitClosure(c *ClosureContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)
}
