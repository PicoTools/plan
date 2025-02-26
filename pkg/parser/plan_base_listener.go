// Code generated from ./grammar/PLAN.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PLAN
import "github.com/antlr4-go/antlr/v4"

// BasePLANListener is a complete listener for a parse tree produced by PLANParser.
type BasePLANListener struct{}

var _ PLANListener = &BasePLANListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePLANListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePLANListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePLANListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePLANListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgFile is called when production progFile is entered.
func (s *BasePLANListener) EnterProgFile(ctx *ProgFileContext) {}

// ExitProgFile is called when production progFile is exited.
func (s *BasePLANListener) ExitProgFile(ctx *ProgFileContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BasePLANListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BasePLANListener) ExitStmts(ctx *StmtsContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePLANListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePLANListener) ExitStmt(ctx *StmtContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BasePLANListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BasePLANListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterCompoundStmt is called when production compoundStmt is entered.
func (s *BasePLANListener) EnterCompoundStmt(ctx *CompoundStmtContext) {}

// ExitCompoundStmt is called when production compoundStmt is exited.
func (s *BasePLANListener) ExitCompoundStmt(ctx *CompoundStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BasePLANListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BasePLANListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BasePLANListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BasePLANListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BasePLANListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BasePLANListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterIfBlockStmt is called when production ifBlockStmt is entered.
func (s *BasePLANListener) EnterIfBlockStmt(ctx *IfBlockStmtContext) {}

// ExitIfBlockStmt is called when production ifBlockStmt is exited.
func (s *BasePLANListener) ExitIfBlockStmt(ctx *IfBlockStmtContext) {}

// EnterElifBlockStmt is called when production elifBlockStmt is entered.
func (s *BasePLANListener) EnterElifBlockStmt(ctx *ElifBlockStmtContext) {}

// ExitElifBlockStmt is called when production elifBlockStmt is exited.
func (s *BasePLANListener) ExitElifBlockStmt(ctx *ElifBlockStmtContext) {}

// EnterElseBlockStmt is called when production elseBlockStmt is entered.
func (s *BasePLANListener) EnterElseBlockStmt(ctx *ElseBlockStmtContext) {}

// ExitElseBlockStmt is called when production elseBlockStmt is exited.
func (s *BasePLANListener) ExitElseBlockStmt(ctx *ElseBlockStmtContext) {}

// EnterFnStmt is called when production fnStmt is entered.
func (s *BasePLANListener) EnterFnStmt(ctx *FnStmtContext) {}

// ExitFnStmt is called when production fnStmt is exited.
func (s *BasePLANListener) ExitFnStmt(ctx *FnStmtContext) {}

// EnterFnBody is called when production fnBody is entered.
func (s *BasePLANListener) EnterFnBody(ctx *FnBodyContext) {}

// ExitFnBody is called when production fnBody is exited.
func (s *BasePLANListener) ExitFnBody(ctx *FnBodyContext) {}

// EnterFnParams is called when production fnParams is entered.
func (s *BasePLANListener) EnterFnParams(ctx *FnParamsContext) {}

// ExitFnParams is called when production fnParams is exited.
func (s *BasePLANListener) ExitFnParams(ctx *FnParamsContext) {}

// EnterIncludeStmt is called when production includeStmt is entered.
func (s *BasePLANListener) EnterIncludeStmt(ctx *IncludeStmtContext) {}

// ExitIncludeStmt is called when production includeStmt is exited.
func (s *BasePLANListener) ExitIncludeStmt(ctx *IncludeStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BasePLANListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BasePLANListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BasePLANListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BasePLANListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BasePLANListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BasePLANListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterAssignRegular is called when production assignRegular is entered.
func (s *BasePLANListener) EnterAssignRegular(ctx *AssignRegularContext) {}

// ExitAssignRegular is called when production assignRegular is exited.
func (s *BasePLANListener) ExitAssignRegular(ctx *AssignRegularContext) {}

// EnterAssignSum is called when production assignSum is entered.
func (s *BasePLANListener) EnterAssignSum(ctx *AssignSumContext) {}

// ExitAssignSum is called when production assignSum is exited.
func (s *BasePLANListener) ExitAssignSum(ctx *AssignSumContext) {}

// EnterAssignSub is called when production assignSub is entered.
func (s *BasePLANListener) EnterAssignSub(ctx *AssignSubContext) {}

// ExitAssignSub is called when production assignSub is exited.
func (s *BasePLANListener) ExitAssignSub(ctx *AssignSubContext) {}

// EnterAssignMul is called when production assignMul is entered.
func (s *BasePLANListener) EnterAssignMul(ctx *AssignMulContext) {}

// ExitAssignMul is called when production assignMul is exited.
func (s *BasePLANListener) ExitAssignMul(ctx *AssignMulContext) {}

// EnterAssignDiv is called when production assignDiv is entered.
func (s *BasePLANListener) EnterAssignDiv(ctx *AssignDivContext) {}

// ExitAssignDiv is called when production assignDiv is exited.
func (s *BasePLANListener) ExitAssignDiv(ctx *AssignDivContext) {}

// EnterAssignMod is called when production assignMod is entered.
func (s *BasePLANListener) EnterAssignMod(ctx *AssignModContext) {}

// ExitAssignMod is called when production assignMod is exited.
func (s *BasePLANListener) ExitAssignMod(ctx *AssignModContext) {}

// EnterAssignPow is called when production assignPow is entered.
func (s *BasePLANListener) EnterAssignPow(ctx *AssignPowContext) {}

// ExitAssignPow is called when production assignPow is exited.
func (s *BasePLANListener) ExitAssignPow(ctx *AssignPowContext) {}

// EnterAssignIdxsRegular is called when production assignIdxsRegular is entered.
func (s *BasePLANListener) EnterAssignIdxsRegular(ctx *AssignIdxsRegularContext) {}

// ExitAssignIdxsRegular is called when production assignIdxsRegular is exited.
func (s *BasePLANListener) ExitAssignIdxsRegular(ctx *AssignIdxsRegularContext) {}

// EnterAssignIdxsSum is called when production assignIdxsSum is entered.
func (s *BasePLANListener) EnterAssignIdxsSum(ctx *AssignIdxsSumContext) {}

// ExitAssignIdxsSum is called when production assignIdxsSum is exited.
func (s *BasePLANListener) ExitAssignIdxsSum(ctx *AssignIdxsSumContext) {}

// EnterAssignIdxsSub is called when production assignIdxsSub is entered.
func (s *BasePLANListener) EnterAssignIdxsSub(ctx *AssignIdxsSubContext) {}

// ExitAssignIdxsSub is called when production assignIdxsSub is exited.
func (s *BasePLANListener) ExitAssignIdxsSub(ctx *AssignIdxsSubContext) {}

// EnterAssignIdxsMul is called when production assignIdxsMul is entered.
func (s *BasePLANListener) EnterAssignIdxsMul(ctx *AssignIdxsMulContext) {}

// ExitAssignIdxsMul is called when production assignIdxsMul is exited.
func (s *BasePLANListener) ExitAssignIdxsMul(ctx *AssignIdxsMulContext) {}

// EnterAssignIdxsDiv is called when production assignIdxsDiv is entered.
func (s *BasePLANListener) EnterAssignIdxsDiv(ctx *AssignIdxsDivContext) {}

// ExitAssignIdxsDiv is called when production assignIdxsDiv is exited.
func (s *BasePLANListener) ExitAssignIdxsDiv(ctx *AssignIdxsDivContext) {}

// EnterAssignIdxsMod is called when production assignIdxsMod is entered.
func (s *BasePLANListener) EnterAssignIdxsMod(ctx *AssignIdxsModContext) {}

// ExitAssignIdxsMod is called when production assignIdxsMod is exited.
func (s *BasePLANListener) ExitAssignIdxsMod(ctx *AssignIdxsModContext) {}

// EnterAssignIdxsPow is called when production assignIdxsPow is entered.
func (s *BasePLANListener) EnterAssignIdxsPow(ctx *AssignIdxsPowContext) {}

// ExitAssignIdxsPow is called when production assignIdxsPow is exited.
func (s *BasePLANListener) ExitAssignIdxsPow(ctx *AssignIdxsPowContext) {}

// EnterList is called when production list is entered.
func (s *BasePLANListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasePLANListener) ExitList(ctx *ListContext) {}

// EnterDictUnit is called when production dictUnit is entered.
func (s *BasePLANListener) EnterDictUnit(ctx *DictUnitContext) {}

// ExitDictUnit is called when production dictUnit is exited.
func (s *BasePLANListener) ExitDictUnit(ctx *DictUnitContext) {}

// EnterDict is called when production dict is entered.
func (s *BasePLANListener) EnterDict(ctx *DictContext) {}

// ExitDict is called when production dict is exited.
func (s *BasePLANListener) ExitDict(ctx *DictContext) {}

// EnterIdx is called when production idx is entered.
func (s *BasePLANListener) EnterIdx(ctx *IdxContext) {}

// ExitIdx is called when production idx is exited.
func (s *BasePLANListener) ExitIdx(ctx *IdxContext) {}

// EnterIdxs is called when production idxs is entered.
func (s *BasePLANListener) EnterIdxs(ctx *IdxsContext) {}

// ExitIdxs is called when production idxs is exited.
func (s *BasePLANListener) ExitIdxs(ctx *IdxsContext) {}

// EnterIdentifierMethodInvoke is called when production identifierMethodInvoke is entered.
func (s *BasePLANListener) EnterIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) {}

// ExitIdentifierMethodInvoke is called when production identifierMethodInvoke is exited.
func (s *BasePLANListener) ExitIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) {}

// EnterIdentifierFnInvoke is called when production identifierFnInvoke is entered.
func (s *BasePLANListener) EnterIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) {}

// ExitIdentifierFnInvoke is called when production identifierFnInvoke is exited.
func (s *BasePLANListener) ExitIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) {}

// EnterIdentifierCsInvoke is called when production identifierCsInvoke is entered.
func (s *BasePLANListener) EnterIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) {}

// ExitIdentifierCsInvoke is called when production identifierCsInvoke is exited.
func (s *BasePLANListener) ExitIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) {}

// EnterExpBool is called when production expBool is entered.
func (s *BasePLANListener) EnterExpBool(ctx *ExpBoolContext) {}

// ExitExpBool is called when production expBool is exited.
func (s *BasePLANListener) ExitExpBool(ctx *ExpBoolContext) {}

// EnterExpComparison is called when production expComparison is entered.
func (s *BasePLANListener) EnterExpComparison(ctx *ExpComparisonContext) {}

// ExitExpComparison is called when production expComparison is exited.
func (s *BasePLANListener) ExitExpComparison(ctx *ExpComparisonContext) {}

// EnterExpIdx is called when production expIdx is entered.
func (s *BasePLANListener) EnterExpIdx(ctx *ExpIdxContext) {}

// ExitExpIdx is called when production expIdx is exited.
func (s *BasePLANListener) ExitExpIdx(ctx *ExpIdxContext) {}

// EnterExpString is called when production expString is entered.
func (s *BasePLANListener) EnterExpString(ctx *ExpStringContext) {}

// ExitExpString is called when production expString is exited.
func (s *BasePLANListener) ExitExpString(ctx *ExpStringContext) {}

// EnterExpCsInvoke is called when production expCsInvoke is entered.
func (s *BasePLANListener) EnterExpCsInvoke(ctx *ExpCsInvokeContext) {}

// ExitExpCsInvoke is called when production expCsInvoke is exited.
func (s *BasePLANListener) ExitExpCsInvoke(ctx *ExpCsInvokeContext) {}

// EnterExpFloat is called when production expFloat is entered.
func (s *BasePLANListener) EnterExpFloat(ctx *ExpFloatContext) {}

// ExitExpFloat is called when production expFloat is exited.
func (s *BasePLANListener) ExitExpFloat(ctx *ExpFloatContext) {}

// EnterExpPow is called when production expPow is entered.
func (s *BasePLANListener) EnterExpPow(ctx *ExpPowContext) {}

// ExitExpPow is called when production expPow is exited.
func (s *BasePLANListener) ExitExpPow(ctx *ExpPowContext) {}

// EnterExpDict is called when production expDict is entered.
func (s *BasePLANListener) EnterExpDict(ctx *ExpDictContext) {}

// ExitExpDict is called when production expDict is exited.
func (s *BasePLANListener) ExitExpDict(ctx *ExpDictContext) {}

// EnterExpXor is called when production expXor is entered.
func (s *BasePLANListener) EnterExpXor(ctx *ExpXorContext) {}

// ExitExpXor is called when production expXor is exited.
func (s *BasePLANListener) ExitExpXor(ctx *ExpXorContext) {}

// EnterExpNeg is called when production expNeg is entered.
func (s *BasePLANListener) EnterExpNeg(ctx *ExpNegContext) {}

// ExitExpNeg is called when production expNeg is exited.
func (s *BasePLANListener) ExitExpNeg(ctx *ExpNegContext) {}

// EnterExpInteger is called when production expInteger is entered.
func (s *BasePLANListener) EnterExpInteger(ctx *ExpIntegerContext) {}

// ExitExpInteger is called when production expInteger is exited.
func (s *BasePLANListener) ExitExpInteger(ctx *ExpIntegerContext) {}

// EnterExpLogicalOr is called when production expLogicalOr is entered.
func (s *BasePLANListener) EnterExpLogicalOr(ctx *ExpLogicalOrContext) {}

// ExitExpLogicalOr is called when production expLogicalOr is exited.
func (s *BasePLANListener) ExitExpLogicalOr(ctx *ExpLogicalOrContext) {}

// EnterExpCs is called when production expCs is entered.
func (s *BasePLANListener) EnterExpCs(ctx *ExpCsContext) {}

// ExitExpCs is called when production expCs is exited.
func (s *BasePLANListener) ExitExpCs(ctx *ExpCsContext) {}

// EnterExpMulDivMod is called when production expMulDivMod is entered.
func (s *BasePLANListener) EnterExpMulDivMod(ctx *ExpMulDivModContext) {}

// ExitExpMulDivMod is called when production expMulDivMod is exited.
func (s *BasePLANListener) ExitExpMulDivMod(ctx *ExpMulDivModContext) {}

// EnterExpNull is called when production expNull is entered.
func (s *BasePLANListener) EnterExpNull(ctx *ExpNullContext) {}

// ExitExpNull is called when production expNull is exited.
func (s *BasePLANListener) ExitExpNull(ctx *ExpNullContext) {}

// EnterExpFnInvoke is called when production expFnInvoke is entered.
func (s *BasePLANListener) EnterExpFnInvoke(ctx *ExpFnInvokeContext) {}

// ExitExpFnInvoke is called when production expFnInvoke is exited.
func (s *BasePLANListener) ExitExpFnInvoke(ctx *ExpFnInvokeContext) {}

// EnterExpList is called when production expList is entered.
func (s *BasePLANListener) EnterExpList(ctx *ExpListContext) {}

// ExitExpList is called when production expList is exited.
func (s *BasePLANListener) ExitExpList(ctx *ExpListContext) {}

// EnterExpLogicalAnd is called when production expLogicalAnd is entered.
func (s *BasePLANListener) EnterExpLogicalAnd(ctx *ExpLogicalAndContext) {}

// ExitExpLogicalAnd is called when production expLogicalAnd is exited.
func (s *BasePLANListener) ExitExpLogicalAnd(ctx *ExpLogicalAndContext) {}

// EnterExpParentheses is called when production expParentheses is entered.
func (s *BasePLANListener) EnterExpParentheses(ctx *ExpParenthesesContext) {}

// ExitExpParentheses is called when production expParentheses is exited.
func (s *BasePLANListener) ExitExpParentheses(ctx *ExpParenthesesContext) {}

// EnterExpEqual is called when production expEqual is entered.
func (s *BasePLANListener) EnterExpEqual(ctx *ExpEqualContext) {}

// ExitExpEqual is called when production expEqual is exited.
func (s *BasePLANListener) ExitExpEqual(ctx *ExpEqualContext) {}

// EnterExpMethodInvoke is called when production expMethodInvoke is entered.
func (s *BasePLANListener) EnterExpMethodInvoke(ctx *ExpMethodInvokeContext) {}

// ExitExpMethodInvoke is called when production expMethodInvoke is exited.
func (s *BasePLANListener) ExitExpMethodInvoke(ctx *ExpMethodInvokeContext) {}

// EnterExpLogicalNot is called when production expLogicalNot is entered.
func (s *BasePLANListener) EnterExpLogicalNot(ctx *ExpLogicalNotContext) {}

// ExitExpLogicalNot is called when production expLogicalNot is exited.
func (s *BasePLANListener) ExitExpLogicalNot(ctx *ExpLogicalNotContext) {}

// EnterExpIntegerHex is called when production expIntegerHex is entered.
func (s *BasePLANListener) EnterExpIntegerHex(ctx *ExpIntegerHexContext) {}

// ExitExpIntegerHex is called when production expIntegerHex is exited.
func (s *BasePLANListener) ExitExpIntegerHex(ctx *ExpIntegerHexContext) {}

// EnterExpIdentifier is called when production expIdentifier is entered.
func (s *BasePLANListener) EnterExpIdentifier(ctx *ExpIdentifierContext) {}

// ExitExpIdentifier is called when production expIdentifier is exited.
func (s *BasePLANListener) ExitExpIdentifier(ctx *ExpIdentifierContext) {}

// EnterExpSumSub is called when production expSumSub is entered.
func (s *BasePLANListener) EnterExpSumSub(ctx *ExpSumSubContext) {}

// ExitExpSumSub is called when production expSumSub is exited.
func (s *BasePLANListener) ExitExpSumSub(ctx *ExpSumSubContext) {}

// EnterClosure is called when production closure is entered.
func (s *BasePLANListener) EnterClosure(ctx *ClosureContext) {}

// ExitClosure is called when production closure is exited.
func (s *BasePLANListener) ExitClosure(ctx *ClosureContext) {}
