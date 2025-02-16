// Code generated from ./grammar/PLAN.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PLAN
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PLANParser struct {
	*antlr.BaseParser
}

var PLANParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func planParserInit() {
	staticData := &PLANParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'while'", "'{'", "'}'", "'for'", "'return'", "'continue'",
		"'break'", "'='", "'['", "','", "']'", "':'", "'.'", "'('", "')'", "'if'",
		"'elif'", "'else'", "'fn'", "'include'", "", "'=='", "'!='", "'||'",
		"'&&'", "'**'", "'>='", "'<='", "'+='", "'-='", "'*='", "'/='", "'%='",
		"'**='", "'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'^'", "'!'",
		"'@'", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "WS", "Eq", "Neq", "Or", "And", "Pow", "GtEq", "LtEq",
		"AssSum", "AssSub", "AssMul", "AssDiv", "AssMod", "AssPow", "Gt", "Lt",
		"Multiply", "Division", "Modulus", "Add", "Subtract", "Xor", "Not",
		"Closure", "Bool", "Null", "Identifier", "Integer", "IntegerHex", "Float",
		"String", "Comment",
	}
	staticData.RuleNames = []string{
		"prog", "stmt", "whileStmt", "forStmt", "returnStmt", "continueStmt",
		"breakStmt", "assignment", "list", "dictUnit", "dict", "idx", "methodInvoke",
		"fnInvoke", "csInvoke", "exp", "ifBlock", "elifBlock", "elseBlock",
		"ifStmt", "fnParams", "fnBody", "fn", "closure", "include",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 53, 390, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 85, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 91, 8, 2, 10, 2, 12, 2, 94, 9, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 106, 8, 3, 10, 3, 12,
		3, 109, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 115, 8, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 177, 8, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 183, 8, 8, 10, 8, 12, 8, 186, 9, 8, 3, 8,
		188, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 5, 10, 200, 8, 10, 10, 10, 12, 10, 203, 9, 10, 3, 10, 205, 8, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 220, 8, 12, 10, 12, 12, 12, 223, 9, 12, 3, 12, 225,
		8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 234, 8,
		13, 10, 13, 12, 13, 237, 9, 13, 3, 13, 239, 8, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 249, 8, 14, 10, 14, 12, 14, 252,
		9, 14, 3, 14, 254, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 280, 8, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 308, 8, 15, 10, 15, 12, 15, 311,
		9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 317, 8, 16, 10, 16, 12, 16, 320,
		9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 328, 8, 17, 10,
		17, 12, 17, 331, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 5, 18, 338,
		8, 18, 10, 18, 12, 18, 341, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 347,
		8, 19, 10, 19, 12, 19, 350, 9, 19, 1, 19, 3, 19, 353, 8, 19, 1, 20, 1,
		20, 1, 20, 5, 20, 358, 8, 20, 10, 20, 12, 20, 361, 9, 20, 1, 21, 1, 21,
		3, 21, 365, 8, 21, 1, 21, 1, 21, 1, 21, 5, 21, 370, 8, 21, 10, 21, 12,
		21, 373, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 0, 1, 30, 25, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 0, 4, 1, 0, 38, 40, 1, 0, 41, 42, 2, 0, 28, 29, 36,
		37, 1, 0, 23, 24, 434, 0, 55, 1, 0, 0, 0, 2, 84, 1, 0, 0, 0, 4, 86, 1,
		0, 0, 0, 6, 97, 1, 0, 0, 0, 8, 112, 1, 0, 0, 0, 10, 116, 1, 0, 0, 0, 12,
		118, 1, 0, 0, 0, 14, 176, 1, 0, 0, 0, 16, 178, 1, 0, 0, 0, 18, 191, 1,
		0, 0, 0, 20, 195, 1, 0, 0, 0, 22, 208, 1, 0, 0, 0, 24, 212, 1, 0, 0, 0,
		26, 228, 1, 0, 0, 0, 28, 242, 1, 0, 0, 0, 30, 279, 1, 0, 0, 0, 32, 312,
		1, 0, 0, 0, 34, 323, 1, 0, 0, 0, 36, 334, 1, 0, 0, 0, 38, 344, 1, 0, 0,
		0, 40, 354, 1, 0, 0, 0, 42, 362, 1, 0, 0, 0, 44, 376, 1, 0, 0, 0, 46, 380,
		1, 0, 0, 0, 48, 383, 1, 0, 0, 0, 50, 54, 3, 2, 1, 0, 51, 54, 3, 44, 22,
		0, 52, 54, 3, 48, 24, 0, 53, 50, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 52,
		1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 5, 0, 0, 1, 59, 1, 1, 0,
		0, 0, 60, 61, 3, 14, 7, 0, 61, 62, 5, 1, 0, 0, 62, 85, 1, 0, 0, 0, 63,
		64, 3, 24, 12, 0, 64, 65, 5, 1, 0, 0, 65, 85, 1, 0, 0, 0, 66, 67, 3, 28,
		14, 0, 67, 68, 5, 1, 0, 0, 68, 85, 1, 0, 0, 0, 69, 70, 3, 26, 13, 0, 70,
		71, 5, 1, 0, 0, 71, 85, 1, 0, 0, 0, 72, 73, 3, 12, 6, 0, 73, 74, 5, 1,
		0, 0, 74, 85, 1, 0, 0, 0, 75, 76, 3, 10, 5, 0, 76, 77, 5, 1, 0, 0, 77,
		85, 1, 0, 0, 0, 78, 79, 3, 8, 4, 0, 79, 80, 5, 1, 0, 0, 80, 85, 1, 0, 0,
		0, 81, 85, 3, 4, 2, 0, 82, 85, 3, 6, 3, 0, 83, 85, 3, 38, 19, 0, 84, 60,
		1, 0, 0, 0, 84, 63, 1, 0, 0, 0, 84, 66, 1, 0, 0, 0, 84, 69, 1, 0, 0, 0,
		84, 72, 1, 0, 0, 0, 84, 75, 1, 0, 0, 0, 84, 78, 1, 0, 0, 0, 84, 81, 1,
		0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0, 86,
		87, 5, 2, 0, 0, 87, 88, 3, 30, 15, 0, 88, 92, 5, 3, 0, 0, 89, 91, 3, 2,
		1, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 4, 0, 0,
		96, 5, 1, 0, 0, 0, 97, 98, 5, 5, 0, 0, 98, 99, 3, 14, 7, 0, 99, 100, 5,
		1, 0, 0, 100, 101, 3, 30, 15, 0, 101, 102, 5, 1, 0, 0, 102, 103, 3, 14,
		7, 0, 103, 107, 5, 3, 0, 0, 104, 106, 3, 2, 1, 0, 105, 104, 1, 0, 0, 0,
		106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 5, 4, 0, 0, 111, 7, 1,
		0, 0, 0, 112, 114, 5, 6, 0, 0, 113, 115, 3, 30, 15, 0, 114, 113, 1, 0,
		0, 0, 114, 115, 1, 0, 0, 0, 115, 9, 1, 0, 0, 0, 116, 117, 5, 7, 0, 0, 117,
		11, 1, 0, 0, 0, 118, 119, 5, 8, 0, 0, 119, 13, 1, 0, 0, 0, 120, 121, 5,
		48, 0, 0, 121, 122, 5, 9, 0, 0, 122, 177, 3, 30, 15, 0, 123, 124, 5, 48,
		0, 0, 124, 125, 5, 30, 0, 0, 125, 177, 3, 30, 15, 0, 126, 127, 5, 48, 0,
		0, 127, 128, 5, 31, 0, 0, 128, 177, 3, 30, 15, 0, 129, 130, 5, 48, 0, 0,
		130, 131, 5, 32, 0, 0, 131, 177, 3, 30, 15, 0, 132, 133, 5, 48, 0, 0, 133,
		134, 5, 33, 0, 0, 134, 177, 3, 30, 15, 0, 135, 136, 5, 48, 0, 0, 136, 137,
		5, 34, 0, 0, 137, 177, 3, 30, 15, 0, 138, 139, 5, 48, 0, 0, 139, 140, 5,
		35, 0, 0, 140, 177, 3, 30, 15, 0, 141, 142, 5, 48, 0, 0, 142, 143, 3, 22,
		11, 0, 143, 144, 5, 9, 0, 0, 144, 145, 3, 30, 15, 0, 145, 177, 1, 0, 0,
		0, 146, 147, 5, 48, 0, 0, 147, 148, 3, 22, 11, 0, 148, 149, 5, 30, 0, 0,
		149, 150, 3, 30, 15, 0, 150, 177, 1, 0, 0, 0, 151, 152, 5, 48, 0, 0, 152,
		153, 3, 22, 11, 0, 153, 154, 5, 31, 0, 0, 154, 155, 3, 30, 15, 0, 155,
		177, 1, 0, 0, 0, 156, 157, 5, 48, 0, 0, 157, 158, 3, 22, 11, 0, 158, 159,
		5, 32, 0, 0, 159, 160, 3, 30, 15, 0, 160, 177, 1, 0, 0, 0, 161, 162, 5,
		48, 0, 0, 162, 163, 3, 22, 11, 0, 163, 164, 5, 33, 0, 0, 164, 165, 3, 30,
		15, 0, 165, 177, 1, 0, 0, 0, 166, 167, 5, 48, 0, 0, 167, 168, 3, 22, 11,
		0, 168, 169, 5, 34, 0, 0, 169, 170, 3, 30, 15, 0, 170, 177, 1, 0, 0, 0,
		171, 172, 5, 48, 0, 0, 172, 173, 3, 22, 11, 0, 173, 174, 5, 35, 0, 0, 174,
		175, 3, 30, 15, 0, 175, 177, 1, 0, 0, 0, 176, 120, 1, 0, 0, 0, 176, 123,
		1, 0, 0, 0, 176, 126, 1, 0, 0, 0, 176, 129, 1, 0, 0, 0, 176, 132, 1, 0,
		0, 0, 176, 135, 1, 0, 0, 0, 176, 138, 1, 0, 0, 0, 176, 141, 1, 0, 0, 0,
		176, 146, 1, 0, 0, 0, 176, 151, 1, 0, 0, 0, 176, 156, 1, 0, 0, 0, 176,
		161, 1, 0, 0, 0, 176, 166, 1, 0, 0, 0, 176, 171, 1, 0, 0, 0, 177, 15, 1,
		0, 0, 0, 178, 187, 5, 10, 0, 0, 179, 184, 3, 30, 15, 0, 180, 181, 5, 11,
		0, 0, 181, 183, 3, 30, 15, 0, 182, 180, 1, 0, 0, 0, 183, 186, 1, 0, 0,
		0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 187, 179, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 190, 5, 12, 0, 0, 190, 17, 1, 0, 0, 0, 191, 192, 3, 30,
		15, 0, 192, 193, 5, 13, 0, 0, 193, 194, 3, 30, 15, 0, 194, 19, 1, 0, 0,
		0, 195, 204, 5, 3, 0, 0, 196, 201, 3, 18, 9, 0, 197, 198, 5, 11, 0, 0,
		198, 200, 3, 18, 9, 0, 199, 197, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201,
		199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201,
		1, 0, 0, 0, 204, 196, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0,
		0, 0, 206, 207, 5, 4, 0, 0, 207, 21, 1, 0, 0, 0, 208, 209, 5, 10, 0, 0,
		209, 210, 3, 30, 15, 0, 210, 211, 5, 12, 0, 0, 211, 23, 1, 0, 0, 0, 212,
		213, 5, 48, 0, 0, 213, 214, 5, 14, 0, 0, 214, 215, 5, 48, 0, 0, 215, 224,
		5, 15, 0, 0, 216, 221, 3, 30, 15, 0, 217, 218, 5, 11, 0, 0, 218, 220, 3,
		30, 15, 0, 219, 217, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0,
		0, 0, 221, 222, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0,
		224, 216, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226,
		227, 5, 16, 0, 0, 227, 25, 1, 0, 0, 0, 228, 229, 5, 48, 0, 0, 229, 238,
		5, 15, 0, 0, 230, 235, 3, 30, 15, 0, 231, 232, 5, 11, 0, 0, 232, 234, 3,
		30, 15, 0, 233, 231, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0,
		0, 0, 235, 236, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0,
		238, 230, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240,
		241, 5, 16, 0, 0, 241, 27, 1, 0, 0, 0, 242, 243, 5, 45, 0, 0, 243, 244,
		5, 48, 0, 0, 244, 253, 5, 15, 0, 0, 245, 250, 3, 30, 15, 0, 246, 247, 5,
		11, 0, 0, 247, 249, 3, 30, 15, 0, 248, 246, 1, 0, 0, 0, 249, 252, 1, 0,
		0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0,
		252, 250, 1, 0, 0, 0, 253, 245, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		255, 1, 0, 0, 0, 255, 256, 5, 16, 0, 0, 256, 29, 1, 0, 0, 0, 257, 258,
		6, 15, -1, 0, 258, 280, 5, 49, 0, 0, 259, 280, 5, 50, 0, 0, 260, 280, 5,
		47, 0, 0, 261, 280, 5, 46, 0, 0, 262, 280, 5, 48, 0, 0, 263, 280, 5, 51,
		0, 0, 264, 280, 5, 52, 0, 0, 265, 280, 3, 46, 23, 0, 266, 280, 3, 24, 12,
		0, 267, 280, 3, 26, 13, 0, 268, 280, 3, 28, 14, 0, 269, 270, 5, 42, 0,
		0, 270, 280, 3, 30, 15, 13, 271, 272, 5, 44, 0, 0, 272, 280, 3, 30, 15,
		12, 273, 274, 5, 15, 0, 0, 274, 275, 3, 30, 15, 0, 275, 276, 5, 16, 0,
		0, 276, 280, 1, 0, 0, 0, 277, 280, 3, 16, 8, 0, 278, 280, 3, 20, 10, 0,
		279, 257, 1, 0, 0, 0, 279, 259, 1, 0, 0, 0, 279, 260, 1, 0, 0, 0, 279,
		261, 1, 0, 0, 0, 279, 262, 1, 0, 0, 0, 279, 263, 1, 0, 0, 0, 279, 264,
		1, 0, 0, 0, 279, 265, 1, 0, 0, 0, 279, 266, 1, 0, 0, 0, 279, 267, 1, 0,
		0, 0, 279, 268, 1, 0, 0, 0, 279, 269, 1, 0, 0, 0, 279, 271, 1, 0, 0, 0,
		279, 273, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 278, 1, 0, 0, 0, 280,
		309, 1, 0, 0, 0, 281, 282, 10, 11, 0, 0, 282, 283, 5, 27, 0, 0, 283, 308,
		3, 30, 15, 11, 284, 285, 10, 10, 0, 0, 285, 286, 7, 0, 0, 0, 286, 308,
		3, 30, 15, 11, 287, 288, 10, 9, 0, 0, 288, 289, 7, 1, 0, 0, 289, 308, 3,
		30, 15, 10, 290, 291, 10, 8, 0, 0, 291, 292, 7, 2, 0, 0, 292, 308, 3, 30,
		15, 9, 293, 294, 10, 7, 0, 0, 294, 295, 7, 3, 0, 0, 295, 308, 3, 30, 15,
		8, 296, 297, 10, 6, 0, 0, 297, 298, 5, 43, 0, 0, 298, 308, 3, 30, 15, 7,
		299, 300, 10, 5, 0, 0, 300, 301, 5, 26, 0, 0, 301, 308, 3, 30, 15, 6, 302,
		303, 10, 4, 0, 0, 303, 304, 5, 25, 0, 0, 304, 308, 3, 30, 15, 5, 305, 306,
		10, 14, 0, 0, 306, 308, 3, 22, 11, 0, 307, 281, 1, 0, 0, 0, 307, 284, 1,
		0, 0, 0, 307, 287, 1, 0, 0, 0, 307, 290, 1, 0, 0, 0, 307, 293, 1, 0, 0,
		0, 307, 296, 1, 0, 0, 0, 307, 299, 1, 0, 0, 0, 307, 302, 1, 0, 0, 0, 307,
		305, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310,
		1, 0, 0, 0, 310, 31, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 5, 17,
		0, 0, 313, 314, 3, 30, 15, 0, 314, 318, 5, 3, 0, 0, 315, 317, 3, 2, 1,
		0, 316, 315, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318,
		319, 1, 0, 0, 0, 319, 321, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 322,
		5, 4, 0, 0, 322, 33, 1, 0, 0, 0, 323, 324, 5, 18, 0, 0, 324, 325, 3, 30,
		15, 0, 325, 329, 5, 3, 0, 0, 326, 328, 3, 2, 1, 0, 327, 326, 1, 0, 0, 0,
		328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330,
		332, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 332, 333, 5, 4, 0, 0, 333, 35, 1,
		0, 0, 0, 334, 335, 5, 19, 0, 0, 335, 339, 5, 3, 0, 0, 336, 338, 3, 2, 1,
		0, 337, 336, 1, 0, 0, 0, 338, 341, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339,
		340, 1, 0, 0, 0, 340, 342, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 342, 343,
		5, 4, 0, 0, 343, 37, 1, 0, 0, 0, 344, 348, 3, 32, 16, 0, 345, 347, 3, 34,
		17, 0, 346, 345, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0,
		348, 349, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351,
		353, 3, 36, 18, 0, 352, 351, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 39,
		1, 0, 0, 0, 354, 359, 5, 48, 0, 0, 355, 356, 5, 11, 0, 0, 356, 358, 5,
		48, 0, 0, 357, 355, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0,
		0, 359, 360, 1, 0, 0, 0, 360, 41, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362,
		364, 5, 15, 0, 0, 363, 365, 3, 40, 20, 0, 364, 363, 1, 0, 0, 0, 364, 365,
		1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 16, 0, 0, 367, 371, 5, 3,
		0, 0, 368, 370, 3, 2, 1, 0, 369, 368, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0,
		371, 369, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0, 373,
		371, 1, 0, 0, 0, 374, 375, 5, 4, 0, 0, 375, 43, 1, 0, 0, 0, 376, 377, 5,
		20, 0, 0, 377, 378, 5, 48, 0, 0, 378, 379, 3, 42, 21, 0, 379, 45, 1, 0,
		0, 0, 380, 381, 5, 20, 0, 0, 381, 382, 3, 42, 21, 0, 382, 47, 1, 0, 0,
		0, 383, 384, 5, 21, 0, 0, 384, 385, 5, 15, 0, 0, 385, 386, 3, 30, 15, 0,
		386, 387, 5, 16, 0, 0, 387, 388, 5, 1, 0, 0, 388, 49, 1, 0, 0, 0, 28, 53,
		55, 84, 92, 107, 114, 176, 184, 187, 201, 204, 221, 224, 235, 238, 250,
		253, 279, 307, 309, 318, 329, 339, 348, 352, 359, 364, 371,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PLANParserInit initializes any static state used to implement PLANParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPLANParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PLANParserInit() {
	staticData := &PLANParserStaticData
	staticData.once.Do(planParserInit)
}

// NewPLANParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPLANParser(input antlr.TokenStream) *PLANParser {
	PLANParserInit()
	this := new(PLANParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PLANParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PLAN.g4"

	return this
}

// PLANParser tokens.
const (
	PLANParserEOF        = antlr.TokenEOF
	PLANParserT__0       = 1
	PLANParserT__1       = 2
	PLANParserT__2       = 3
	PLANParserT__3       = 4
	PLANParserT__4       = 5
	PLANParserT__5       = 6
	PLANParserT__6       = 7
	PLANParserT__7       = 8
	PLANParserT__8       = 9
	PLANParserT__9       = 10
	PLANParserT__10      = 11
	PLANParserT__11      = 12
	PLANParserT__12      = 13
	PLANParserT__13      = 14
	PLANParserT__14      = 15
	PLANParserT__15      = 16
	PLANParserT__16      = 17
	PLANParserT__17      = 18
	PLANParserT__18      = 19
	PLANParserT__19      = 20
	PLANParserT__20      = 21
	PLANParserWS         = 22
	PLANParserEq         = 23
	PLANParserNeq        = 24
	PLANParserOr         = 25
	PLANParserAnd        = 26
	PLANParserPow        = 27
	PLANParserGtEq       = 28
	PLANParserLtEq       = 29
	PLANParserAssSum     = 30
	PLANParserAssSub     = 31
	PLANParserAssMul     = 32
	PLANParserAssDiv     = 33
	PLANParserAssMod     = 34
	PLANParserAssPow     = 35
	PLANParserGt         = 36
	PLANParserLt         = 37
	PLANParserMultiply   = 38
	PLANParserDivision   = 39
	PLANParserModulus    = 40
	PLANParserAdd        = 41
	PLANParserSubtract   = 42
	PLANParserXor        = 43
	PLANParserNot        = 44
	PLANParserClosure    = 45
	PLANParserBool       = 46
	PLANParserNull       = 47
	PLANParserIdentifier = 48
	PLANParserInteger    = 49
	PLANParserIntegerHex = 50
	PLANParserFloat      = 51
	PLANParserString_    = 52
	PLANParserComment    = 53
)

// PLANParser rules.
const (
	PLANParserRULE_prog         = 0
	PLANParserRULE_stmt         = 1
	PLANParserRULE_whileStmt    = 2
	PLANParserRULE_forStmt      = 3
	PLANParserRULE_returnStmt   = 4
	PLANParserRULE_continueStmt = 5
	PLANParserRULE_breakStmt    = 6
	PLANParserRULE_assignment   = 7
	PLANParserRULE_list         = 8
	PLANParserRULE_dictUnit     = 9
	PLANParserRULE_dict         = 10
	PLANParserRULE_idx          = 11
	PLANParserRULE_methodInvoke = 12
	PLANParserRULE_fnInvoke     = 13
	PLANParserRULE_csInvoke     = 14
	PLANParserRULE_exp          = 15
	PLANParserRULE_ifBlock      = 16
	PLANParserRULE_elifBlock    = 17
	PLANParserRULE_elseBlock    = 18
	PLANParserRULE_ifStmt       = 19
	PLANParserRULE_fnParams     = 20
	PLANParserRULE_fnBody       = 21
	PLANParserRULE_fn           = 22
	PLANParserRULE_closure      = 23
	PLANParserRULE_include      = 24
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllFn() []IFnContext
	Fn(i int) IFnContext
	AllInclude() []IIncludeContext
	Include(i int) IIncludeContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(PLANParserEOF, 0)
}

func (s *ProgContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgContext) AllFn() []IFnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFnContext); ok {
			len++
		}
	}

	tst := make([]IFnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFnContext); ok {
			tst[i] = t.(IFnContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Fn(i int) IFnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnContext)
}

func (s *ProgContext) AllInclude() []IIncludeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIncludeContext); ok {
			len++
		}
	}

	tst := make([]IIncludeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIncludeContext); ok {
			tst[i] = t.(IIncludeContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Include(i int) IIncludeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PLANParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659352076772) != 0 {
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PLANParserT__1, PLANParserT__4, PLANParserT__5, PLANParserT__6, PLANParserT__7, PLANParserT__16, PLANParserClosure, PLANParserIdentifier:
			{
				p.SetState(50)
				p.Stmt()
			}

		case PLANParserT__19:
			{
				p.SetState(51)
				p.Fn()
			}

		case PLANParserT__20:
			{
				p.SetState(52)
				p.Include()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(PLANParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	MethodInvoke() IMethodInvokeContext
	CsInvoke() ICsInvokeContext
	FnInvoke() IFnInvokeContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	ReturnStmt() IReturnStmtContext
	WhileStmt() IWhileStmtContext
	ForStmt() IForStmtContext
	IfStmt() IIfStmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StmtContext) MethodInvoke() IMethodInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodInvokeContext)
}

func (s *StmtContext) CsInvoke() ICsInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICsInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICsInvokeContext)
}

func (s *StmtContext) FnInvoke() IFnInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnInvokeContext)
}

func (s *StmtContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StmtContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StmtContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StmtContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PLANParserRULE_stmt)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Assignment()
		}
		{
			p.SetState(61)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.MethodInvoke()
		}
		{
			p.SetState(64)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.CsInvoke()
		}
		{
			p.SetState(67)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.FnInvoke()
		}
		{
			p.SetState(70)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.BreakStmt()
		}
		{
			p.SetState(73)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.ContinueStmt()
		}
		{
			p.SetState(76)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(78)
			p.ReturnStmt()
		}
		{
			p.SetState(79)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(81)
			p.WhileStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(82)
			p.ForStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(83)
			p.IfStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PLANParserRULE_whileStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(PLANParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.exp(0)
	}
	{
		p.SetState(88)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(89)
			p.Stmt()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	Exp() IExpContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForStmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PLANParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(PLANParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Assignment()
	}
	{
		p.SetState(99)
		p.Match(PLANParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.exp(0)
	}
	{
		p.SetState(101)
		p.Match(PLANParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Assignment()
	}
	{
		p.SetState(103)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(104)
			p.Stmt()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PLANParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(PLANParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(113)
			p.exp(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PLANParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(PLANParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PLANParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(PLANParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) CopyAll(ctx *AssignmentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignPowContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignPowContext {
	var p = new(AssignPowContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignPowContext) GetName() antlr.Token { return s.name }

func (s *AssignPowContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignPowContext) AssPow() antlr.TerminalNode {
	return s.GetToken(PLANParserAssPow, 0)
}

func (s *AssignPowContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignPowContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignPowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignPow(s)
	}
}

func (s *AssignPowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignPow(s)
	}
}

func (s *AssignPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxPowContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxPowContext {
	var p = new(AssignIdxPowContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxPowContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxPowContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxPowContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxPowContext) AssPow() antlr.TerminalNode {
	return s.GetToken(PLANParserAssPow, 0)
}

func (s *AssignIdxPowContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxPowContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxPowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxPow(s)
	}
}

func (s *AssignIdxPowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxPow(s)
	}
}

func (s *AssignIdxPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxRegularContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxRegularContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxRegularContext {
	var p = new(AssignIdxRegularContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxRegularContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxRegularContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxRegularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxRegularContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxRegularContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxRegularContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxRegularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxRegular(s)
	}
}

func (s *AssignIdxRegularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxRegular(s)
	}
}

func (s *AssignIdxRegularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxRegular(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignDivContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignDivContext {
	var p = new(AssignDivContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignDivContext) GetName() antlr.Token { return s.name }

func (s *AssignDivContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignDivContext) AssDiv() antlr.TerminalNode {
	return s.GetToken(PLANParserAssDiv, 0)
}

func (s *AssignDivContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignDivContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignDiv(s)
	}
}

func (s *AssignDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignDiv(s)
	}
}

func (s *AssignDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxDivContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxDivContext {
	var p = new(AssignIdxDivContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxDivContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxDivContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxDivContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxDivContext) AssDiv() antlr.TerminalNode {
	return s.GetToken(PLANParserAssDiv, 0)
}

func (s *AssignIdxDivContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxDivContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxDiv(s)
	}
}

func (s *AssignIdxDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxDiv(s)
	}
}

func (s *AssignIdxDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignSumContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSumContext {
	var p = new(AssignSumContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignSumContext) GetName() antlr.Token { return s.name }

func (s *AssignSumContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSumContext) AssSum() antlr.TerminalNode {
	return s.GetToken(PLANParserAssSum, 0)
}

func (s *AssignSumContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignSumContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignSum(s)
	}
}

func (s *AssignSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignSum(s)
	}
}

func (s *AssignSumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignSum(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxSumContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxSumContext {
	var p = new(AssignIdxSumContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxSumContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxSumContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxSumContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxSumContext) AssSum() antlr.TerminalNode {
	return s.GetToken(PLANParserAssSum, 0)
}

func (s *AssignIdxSumContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxSumContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxSum(s)
	}
}

func (s *AssignIdxSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxSum(s)
	}
}

func (s *AssignIdxSumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxSum(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignMulContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignMulContext {
	var p = new(AssignMulContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignMulContext) GetName() antlr.Token { return s.name }

func (s *AssignMulContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignMulContext) AssMul() antlr.TerminalNode {
	return s.GetToken(PLANParserAssMul, 0)
}

func (s *AssignMulContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignMulContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignMul(s)
	}
}

func (s *AssignMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignMul(s)
	}
}

func (s *AssignMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxModContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxModContext {
	var p = new(AssignIdxModContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxModContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxModContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxModContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxModContext) AssMod() antlr.TerminalNode {
	return s.GetToken(PLANParserAssMod, 0)
}

func (s *AssignIdxModContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxModContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxMod(s)
	}
}

func (s *AssignIdxModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxMod(s)
	}
}

func (s *AssignIdxModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxMulContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxMulContext {
	var p = new(AssignIdxMulContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxMulContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxMulContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxMulContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxMulContext) AssMul() antlr.TerminalNode {
	return s.GetToken(PLANParserAssMul, 0)
}

func (s *AssignIdxMulContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxMulContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxMul(s)
	}
}

func (s *AssignIdxMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxMul(s)
	}
}

func (s *AssignIdxMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignSubContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSubContext {
	var p = new(AssignSubContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignSubContext) GetName() antlr.Token { return s.name }

func (s *AssignSubContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSubContext) AssSub() antlr.TerminalNode {
	return s.GetToken(PLANParserAssSub, 0)
}

func (s *AssignSubContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignSubContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignSub(s)
	}
}

func (s *AssignSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignSub(s)
	}
}

func (s *AssignSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignRegularContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignRegularContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignRegularContext {
	var p = new(AssignRegularContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignRegularContext) GetName() antlr.Token { return s.name }

func (s *AssignRegularContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignRegularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignRegularContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignRegularContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignRegularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignRegular(s)
	}
}

func (s *AssignRegularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignRegular(s)
	}
}

func (s *AssignRegularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignRegular(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignModContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignModContext {
	var p = new(AssignModContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignModContext) GetName() antlr.Token { return s.name }

func (s *AssignModContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignModContext) AssMod() antlr.TerminalNode {
	return s.GetToken(PLANParserAssMod, 0)
}

func (s *AssignModContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignModContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignMod(s)
	}
}

func (s *AssignModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignMod(s)
	}
}

func (s *AssignModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxSubContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxSubContext {
	var p = new(AssignIdxSubContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxSubContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxSubContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxSubContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *AssignIdxSubContext) AssSub() antlr.TerminalNode {
	return s.GetToken(PLANParserAssSub, 0)
}

func (s *AssignIdxSubContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignIdxSubContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxSub(s)
	}
}

func (s *AssignIdxSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxSub(s)
	}
}

func (s *AssignIdxSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PLANParserRULE_assignment)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignRegularContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignRegularContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(PLANParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.exp(0)
		}

	case 2:
		localctx = NewAssignSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignSumContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(PLANParserAssSum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.exp(0)
		}

	case 3:
		localctx = NewAssignSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignSubContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(PLANParserAssSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.exp(0)
		}

	case 4:
		localctx = NewAssignMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignMulContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(PLANParserAssMul)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.exp(0)
		}

	case 5:
		localctx = NewAssignDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignDivContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(PLANParserAssDiv)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.exp(0)
		}

	case 6:
		localctx = NewAssignModContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(135)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignModContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(PLANParserAssMod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.exp(0)
		}

	case 7:
		localctx = NewAssignPowContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(138)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignPowContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(PLANParserAssPow)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.exp(0)
		}

	case 8:
		localctx = NewAssignIdxRegularContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(141)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxRegularContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Idx()
		}
		{
			p.SetState(143)
			p.Match(PLANParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.exp(0)
		}

	case 9:
		localctx = NewAssignIdxSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(146)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxSumContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Idx()
		}
		{
			p.SetState(148)
			p.Match(PLANParserAssSum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.exp(0)
		}

	case 10:
		localctx = NewAssignIdxSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(151)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxSubContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Idx()
		}
		{
			p.SetState(153)
			p.Match(PLANParserAssSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.exp(0)
		}

	case 11:
		localctx = NewAssignIdxMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(156)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxMulContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Idx()
		}
		{
			p.SetState(158)
			p.Match(PLANParserAssMul)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.exp(0)
		}

	case 12:
		localctx = NewAssignIdxDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(161)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxDivContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Idx()
		}
		{
			p.SetState(163)
			p.Match(PLANParserAssDiv)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.exp(0)
		}

	case 13:
		localctx = NewAssignIdxModContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(166)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxModContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Idx()
		}
		{
			p.SetState(168)
			p.Match(PLANParserAssMod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.exp(0)
		}

	case 14:
		localctx = NewAssignIdxPowContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(171)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxPowContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Idx()
		}
		{
			p.SetState(173)
			p.Match(PLANParserAssPow)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.exp(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PLANParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(PLANParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(179)
			p.exp(0)
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(180)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(181)
				p.exp(0)
			}

			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(189)
		p.Match(PLANParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictUnitContext is an interface to support dynamic dispatch.
type IDictUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext

	// IsDictUnitContext differentiates from other interfaces.
	IsDictUnitContext()
}

type DictUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictUnitContext() *DictUnitContext {
	var p = new(DictUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_dictUnit
	return p
}

func InitEmptyDictUnitContext(p *DictUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_dictUnit
}

func (*DictUnitContext) IsDictUnitContext() {}

func NewDictUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictUnitContext {
	var p = new(DictUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_dictUnit

	return p
}

func (s *DictUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *DictUnitContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *DictUnitContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *DictUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterDictUnit(s)
	}
}

func (s *DictUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitDictUnit(s)
	}
}

func (s *DictUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitDictUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) DictUnit() (localctx IDictUnitContext) {
	localctx = NewDictUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PLANParserRULE_dictUnit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.exp(0)
	}
	{
		p.SetState(192)
		p.Match(PLANParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.exp(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDictContext is an interface to support dynamic dispatch.
type IDictContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDictUnit() []IDictUnitContext
	DictUnit(i int) IDictUnitContext

	// IsDictContext differentiates from other interfaces.
	IsDictContext()
}

type DictContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictContext() *DictContext {
	var p = new(DictContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_dict
	return p
}

func InitEmptyDictContext(p *DictContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_dict
}

func (*DictContext) IsDictContext() {}

func NewDictContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictContext {
	var p = new(DictContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_dict

	return p
}

func (s *DictContext) GetParser() antlr.Parser { return s.parser }

func (s *DictContext) AllDictUnit() []IDictUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictUnitContext); ok {
			len++
		}
	}

	tst := make([]IDictUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictUnitContext); ok {
			tst[i] = t.(IDictUnitContext)
			i++
		}
	}

	return tst
}

func (s *DictContext) DictUnit(i int) IDictUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictUnitContext)
}

func (s *DictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterDict(s)
	}
}

func (s *DictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitDict(s)
	}
}

func (s *DictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitDict(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Dict() (localctx IDictContext) {
	localctx = NewDictContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PLANParserRULE_dict)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(196)
			p.DictUnit()
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(197)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(198)
				p.DictUnit()
			}

			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(206)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdxContext is an interface to support dynamic dispatch.
type IIdxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext

	// IsIdxContext differentiates from other interfaces.
	IsIdxContext()
}

type IdxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdxContext() *IdxContext {
	var p = new(IdxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_idx
	return p
}

func InitEmptyIdxContext(p *IdxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_idx
}

func (*IdxContext) IsIdxContext() {}

func NewIdxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdxContext {
	var p = new(IdxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_idx

	return p
}

func (s *IdxContext) GetParser() antlr.Parser { return s.parser }

func (s *IdxContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IdxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIdx(s)
	}
}

func (s *IdxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIdx(s)
	}
}

func (s *IdxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIdx(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Idx() (localctx IIdxContext) {
	localctx = NewIdxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PLANParserRULE_idx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(PLANParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.exp(0)
	}
	{
		p.SetState(210)
		p.Match(PLANParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodInvokeContext is an interface to support dynamic dispatch.
type IMethodInvokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMethodInvokeContext differentiates from other interfaces.
	IsMethodInvokeContext()
}

type MethodInvokeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodInvokeContext() *MethodInvokeContext {
	var p = new(MethodInvokeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_methodInvoke
	return p
}

func InitEmptyMethodInvokeContext(p *MethodInvokeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_methodInvoke
}

func (*MethodInvokeContext) IsMethodInvokeContext() {}

func NewMethodInvokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodInvokeContext {
	var p = new(MethodInvokeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_methodInvoke

	return p
}

func (s *MethodInvokeContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodInvokeContext) CopyAll(ctx *MethodInvokeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MethodInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodInvokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierMethodInvokeContext struct {
	MethodInvokeContext
	var_ antlr.Token
	name antlr.Token
}

func NewIdentifierMethodInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierMethodInvokeContext {
	var p = new(IdentifierMethodInvokeContext)

	InitEmptyMethodInvokeContext(&p.MethodInvokeContext)
	p.parser = parser
	p.CopyAll(ctx.(*MethodInvokeContext))

	return p
}

func (s *IdentifierMethodInvokeContext) GetVar_() antlr.Token { return s.var_ }

func (s *IdentifierMethodInvokeContext) GetName() antlr.Token { return s.name }

func (s *IdentifierMethodInvokeContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *IdentifierMethodInvokeContext) SetName(v antlr.Token) { s.name = v }

func (s *IdentifierMethodInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierMethodInvokeContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(PLANParserIdentifier)
}

func (s *IdentifierMethodInvokeContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, i)
}

func (s *IdentifierMethodInvokeContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierMethodInvokeContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IdentifierMethodInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIdentifierMethodInvoke(s)
	}
}

func (s *IdentifierMethodInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIdentifierMethodInvoke(s)
	}
}

func (s *IdentifierMethodInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIdentifierMethodInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) MethodInvoke() (localctx IMethodInvokeContext) {
	localctx = NewMethodInvokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PLANParserRULE_methodInvoke)
	var _la int

	localctx = NewIdentifierMethodInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierMethodInvokeContext).var_ = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(PLANParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierMethodInvokeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(216)
			p.exp(0)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(217)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(218)
				p.exp(0)
			}

			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(226)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnInvokeContext is an interface to support dynamic dispatch.
type IFnInvokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFnInvokeContext differentiates from other interfaces.
	IsFnInvokeContext()
}

type FnInvokeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnInvokeContext() *FnInvokeContext {
	var p = new(FnInvokeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fnInvoke
	return p
}

func InitEmptyFnInvokeContext(p *FnInvokeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fnInvoke
}

func (*FnInvokeContext) IsFnInvokeContext() {}

func NewFnInvokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnInvokeContext {
	var p = new(FnInvokeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_fnInvoke

	return p
}

func (s *FnInvokeContext) GetParser() antlr.Parser { return s.parser }

func (s *FnInvokeContext) CopyAll(ctx *FnInvokeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FnInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnInvokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierFnInvokeContext struct {
	FnInvokeContext
	name antlr.Token
}

func NewIdentifierFnInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierFnInvokeContext {
	var p = new(IdentifierFnInvokeContext)

	InitEmptyFnInvokeContext(&p.FnInvokeContext)
	p.parser = parser
	p.CopyAll(ctx.(*FnInvokeContext))

	return p
}

func (s *IdentifierFnInvokeContext) GetName() antlr.Token { return s.name }

func (s *IdentifierFnInvokeContext) SetName(v antlr.Token) { s.name = v }

func (s *IdentifierFnInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierFnInvokeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *IdentifierFnInvokeContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierFnInvokeContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IdentifierFnInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIdentifierFnInvoke(s)
	}
}

func (s *IdentifierFnInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIdentifierFnInvoke(s)
	}
}

func (s *IdentifierFnInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIdentifierFnInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) FnInvoke() (localctx IFnInvokeContext) {
	localctx = NewFnInvokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PLANParserRULE_fnInvoke)
	var _la int

	localctx = NewIdentifierFnInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierFnInvokeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(230)
			p.exp(0)
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(231)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(232)
				p.exp(0)
			}

			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(240)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICsInvokeContext is an interface to support dynamic dispatch.
type ICsInvokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCsInvokeContext differentiates from other interfaces.
	IsCsInvokeContext()
}

type CsInvokeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCsInvokeContext() *CsInvokeContext {
	var p = new(CsInvokeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_csInvoke
	return p
}

func InitEmptyCsInvokeContext(p *CsInvokeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_csInvoke
}

func (*CsInvokeContext) IsCsInvokeContext() {}

func NewCsInvokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CsInvokeContext {
	var p = new(CsInvokeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_csInvoke

	return p
}

func (s *CsInvokeContext) GetParser() antlr.Parser { return s.parser }

func (s *CsInvokeContext) CopyAll(ctx *CsInvokeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CsInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CsInvokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierCsInvokeContext struct {
	CsInvokeContext
	name antlr.Token
}

func NewIdentifierCsInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierCsInvokeContext {
	var p = new(IdentifierCsInvokeContext)

	InitEmptyCsInvokeContext(&p.CsInvokeContext)
	p.parser = parser
	p.CopyAll(ctx.(*CsInvokeContext))

	return p
}

func (s *IdentifierCsInvokeContext) GetName() antlr.Token { return s.name }

func (s *IdentifierCsInvokeContext) SetName(v antlr.Token) { s.name = v }

func (s *IdentifierCsInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierCsInvokeContext) Closure() antlr.TerminalNode {
	return s.GetToken(PLANParserClosure, 0)
}

func (s *IdentifierCsInvokeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *IdentifierCsInvokeContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierCsInvokeContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IdentifierCsInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIdentifierCsInvoke(s)
	}
}

func (s *IdentifierCsInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIdentifierCsInvoke(s)
	}
}

func (s *IdentifierCsInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIdentifierCsInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) CsInvoke() (localctx ICsInvokeContext) {
	localctx = NewCsInvokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PLANParserRULE_csInvoke)
	var _la int

	localctx = NewIdentifierCsInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(PLANParserClosure)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierCsInvokeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(245)
			p.exp(0)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(246)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(247)
				p.exp(0)
			}

			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(255)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyAll(ctx *ExpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpBoolContext struct {
	ExpContext
}

func NewExpBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpBoolContext {
	var p = new(ExpBoolContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpBoolContext) Bool() antlr.TerminalNode {
	return s.GetToken(PLANParserBool, 0)
}

func (s *ExpBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpBool(s)
	}
}

func (s *ExpBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpBool(s)
	}
}

func (s *ExpBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpComparisonContext struct {
	ExpContext
	op antlr.Token
}

func NewExpComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpComparisonContext {
	var p = new(ExpComparisonContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ExpComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpComparisonContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpComparisonContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpComparisonContext) GtEq() antlr.TerminalNode {
	return s.GetToken(PLANParserGtEq, 0)
}

func (s *ExpComparisonContext) LtEq() antlr.TerminalNode {
	return s.GetToken(PLANParserLtEq, 0)
}

func (s *ExpComparisonContext) Gt() antlr.TerminalNode {
	return s.GetToken(PLANParserGt, 0)
}

func (s *ExpComparisonContext) Lt() antlr.TerminalNode {
	return s.GetToken(PLANParserLt, 0)
}

func (s *ExpComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpComparison(s)
	}
}

func (s *ExpComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpComparison(s)
	}
}

func (s *ExpComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpIdxContext struct {
	ExpContext
}

func NewExpIdxContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpIdxContext {
	var p = new(ExpIdxContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpIdxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpIdxContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpIdxContext) Idx() IIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxContext)
}

func (s *ExpIdxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpIdx(s)
	}
}

func (s *ExpIdxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpIdx(s)
	}
}

func (s *ExpIdxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpIdx(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpStringContext struct {
	ExpContext
}

func NewExpStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpStringContext {
	var p = new(ExpStringContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpStringContext) String_() antlr.TerminalNode {
	return s.GetToken(PLANParserString_, 0)
}

func (s *ExpStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpString(s)
	}
}

func (s *ExpStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpString(s)
	}
}

func (s *ExpStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpCsInvokeContext struct {
	ExpContext
}

func NewExpCsInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpCsInvokeContext {
	var p = new(ExpCsInvokeContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpCsInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpCsInvokeContext) CsInvoke() ICsInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICsInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICsInvokeContext)
}

func (s *ExpCsInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpCsInvoke(s)
	}
}

func (s *ExpCsInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpCsInvoke(s)
	}
}

func (s *ExpCsInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpCsInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpFloatContext struct {
	ExpContext
}

func NewExpFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFloatContext {
	var p = new(ExpFloatContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFloatContext) Float() antlr.TerminalNode {
	return s.GetToken(PLANParserFloat, 0)
}

func (s *ExpFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpFloat(s)
	}
}

func (s *ExpFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpFloat(s)
	}
}

func (s *ExpFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpPowContext struct {
	ExpContext
}

func NewExpPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpPowContext {
	var p = new(ExpPowContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpPowContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpPowContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpPowContext) Pow() antlr.TerminalNode {
	return s.GetToken(PLANParserPow, 0)
}

func (s *ExpPowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpPow(s)
	}
}

func (s *ExpPowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpPow(s)
	}
}

func (s *ExpPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpDictContext struct {
	ExpContext
}

func NewExpDictContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpDictContext {
	var p = new(ExpDictContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpDictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpDictContext) Dict() IDictContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictContext)
}

func (s *ExpDictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpDict(s)
	}
}

func (s *ExpDictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpDict(s)
	}
}

func (s *ExpDictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpDict(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpXorContext struct {
	ExpContext
}

func NewExpXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpXorContext {
	var p = new(ExpXorContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpXorContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpXorContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpXorContext) Xor() antlr.TerminalNode {
	return s.GetToken(PLANParserXor, 0)
}

func (s *ExpXorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpXor(s)
	}
}

func (s *ExpXorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpXor(s)
	}
}

func (s *ExpXorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpXor(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpNegContext struct {
	ExpContext
}

func NewExpNegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNegContext {
	var p = new(ExpNegContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpNegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNegContext) Subtract() antlr.TerminalNode {
	return s.GetToken(PLANParserSubtract, 0)
}

func (s *ExpNegContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpNegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpNeg(s)
	}
}

func (s *ExpNegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpNeg(s)
	}
}

func (s *ExpNegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpNeg(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpIntegerContext struct {
	ExpContext
}

func NewExpIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpIntegerContext {
	var p = new(ExpIntegerContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpIntegerContext) Integer() antlr.TerminalNode {
	return s.GetToken(PLANParserInteger, 0)
}

func (s *ExpIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpInteger(s)
	}
}

func (s *ExpIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpInteger(s)
	}
}

func (s *ExpIntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpLogicalOrContext struct {
	ExpContext
}

func NewExpLogicalOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpLogicalOrContext {
	var p = new(ExpLogicalOrContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpLogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpLogicalOrContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpLogicalOrContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpLogicalOrContext) Or() antlr.TerminalNode {
	return s.GetToken(PLANParserOr, 0)
}

func (s *ExpLogicalOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpLogicalOr(s)
	}
}

func (s *ExpLogicalOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpLogicalOr(s)
	}
}

func (s *ExpLogicalOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpLogicalOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpCsContext struct {
	ExpContext
}

func NewExpCsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpCsContext {
	var p = new(ExpCsContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpCsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpCsContext) Closure() IClosureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureContext)
}

func (s *ExpCsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpCs(s)
	}
}

func (s *ExpCsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpCs(s)
	}
}

func (s *ExpCsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpCs(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpMulDivModContext struct {
	ExpContext
	op antlr.Token
}

func NewExpMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpMulDivModContext {
	var p = new(ExpMulDivModContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpMulDivModContext) GetOp() antlr.Token { return s.op }

func (s *ExpMulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpMulDivModContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpMulDivModContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpMulDivModContext) Multiply() antlr.TerminalNode {
	return s.GetToken(PLANParserMultiply, 0)
}

func (s *ExpMulDivModContext) Division() antlr.TerminalNode {
	return s.GetToken(PLANParserDivision, 0)
}

func (s *ExpMulDivModContext) Modulus() antlr.TerminalNode {
	return s.GetToken(PLANParserModulus, 0)
}

func (s *ExpMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpMulDivMod(s)
	}
}

func (s *ExpMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpMulDivMod(s)
	}
}

func (s *ExpMulDivModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpMulDivMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpNullContext struct {
	ExpContext
}

func NewExpNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNullContext {
	var p = new(ExpNullContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNullContext) Null() antlr.TerminalNode {
	return s.GetToken(PLANParserNull, 0)
}

func (s *ExpNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpNull(s)
	}
}

func (s *ExpNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpNull(s)
	}
}

func (s *ExpNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpFnInvokeContext struct {
	ExpContext
}

func NewExpFnInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFnInvokeContext {
	var p = new(ExpFnInvokeContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpFnInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFnInvokeContext) FnInvoke() IFnInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnInvokeContext)
}

func (s *ExpFnInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpFnInvoke(s)
	}
}

func (s *ExpFnInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpFnInvoke(s)
	}
}

func (s *ExpFnInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpFnInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpListContext struct {
	ExpContext
}

func NewExpListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpListContext {
	var p = new(ExpListContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpListContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ExpListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpList(s)
	}
}

func (s *ExpListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpList(s)
	}
}

func (s *ExpListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpList(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpLogicalAndContext struct {
	ExpContext
}

func NewExpLogicalAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpLogicalAndContext {
	var p = new(ExpLogicalAndContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpLogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpLogicalAndContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpLogicalAndContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpLogicalAndContext) And() antlr.TerminalNode {
	return s.GetToken(PLANParserAnd, 0)
}

func (s *ExpLogicalAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpLogicalAnd(s)
	}
}

func (s *ExpLogicalAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpLogicalAnd(s)
	}
}

func (s *ExpLogicalAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpLogicalAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpParenthesesContext struct {
	ExpContext
}

func NewExpParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpParenthesesContext {
	var p = new(ExpParenthesesContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpParenthesesContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpParentheses(s)
	}
}

func (s *ExpParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpParentheses(s)
	}
}

func (s *ExpParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpEqualContext struct {
	ExpContext
	op antlr.Token
}

func NewExpEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpEqualContext {
	var p = new(ExpEqualContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpEqualContext) GetOp() antlr.Token { return s.op }

func (s *ExpEqualContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpEqualContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpEqualContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpEqualContext) Eq() antlr.TerminalNode {
	return s.GetToken(PLANParserEq, 0)
}

func (s *ExpEqualContext) Neq() antlr.TerminalNode {
	return s.GetToken(PLANParserNeq, 0)
}

func (s *ExpEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpEqual(s)
	}
}

func (s *ExpEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpEqual(s)
	}
}

func (s *ExpEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpMethodInvokeContext struct {
	ExpContext
}

func NewExpMethodInvokeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpMethodInvokeContext {
	var p = new(ExpMethodInvokeContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpMethodInvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpMethodInvokeContext) MethodInvoke() IMethodInvokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodInvokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodInvokeContext)
}

func (s *ExpMethodInvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpMethodInvoke(s)
	}
}

func (s *ExpMethodInvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpMethodInvoke(s)
	}
}

func (s *ExpMethodInvokeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpMethodInvoke(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpLogicalNotContext struct {
	ExpContext
}

func NewExpLogicalNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpLogicalNotContext {
	var p = new(ExpLogicalNotContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpLogicalNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpLogicalNotContext) Not() antlr.TerminalNode {
	return s.GetToken(PLANParserNot, 0)
}

func (s *ExpLogicalNotContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpLogicalNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpLogicalNot(s)
	}
}

func (s *ExpLogicalNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpLogicalNot(s)
	}
}

func (s *ExpLogicalNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpLogicalNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpIntegerHexContext struct {
	ExpContext
}

func NewExpIntegerHexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpIntegerHexContext {
	var p = new(ExpIntegerHexContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpIntegerHexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpIntegerHexContext) IntegerHex() antlr.TerminalNode {
	return s.GetToken(PLANParserIntegerHex, 0)
}

func (s *ExpIntegerHexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpIntegerHex(s)
	}
}

func (s *ExpIntegerHexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpIntegerHex(s)
	}
}

func (s *ExpIntegerHexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpIntegerHex(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpIdentifierContext struct {
	ExpContext
}

func NewExpIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpIdentifierContext {
	var p = new(ExpIdentifierContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *ExpIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpIdentifier(s)
	}
}

func (s *ExpIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpIdentifier(s)
	}
}

func (s *ExpIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpSumSubContext struct {
	ExpContext
	op antlr.Token
}

func NewExpSumSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpSumSubContext {
	var p = new(ExpSumSubContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *ExpSumSubContext) GetOp() antlr.Token { return s.op }

func (s *ExpSumSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpSumSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpSumSubContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpSumSubContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpSumSubContext) Add() antlr.TerminalNode {
	return s.GetToken(PLANParserAdd, 0)
}

func (s *ExpSumSubContext) Subtract() antlr.TerminalNode {
	return s.GetToken(PLANParserSubtract, 0)
}

func (s *ExpSumSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterExpSumSub(s)
	}
}

func (s *ExpSumSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitExpSumSub(s)
	}
}

func (s *ExpSumSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitExpSumSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *PLANParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, PLANParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(258)
			p.Match(PLANParserInteger)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExpIntegerHexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.Match(PLANParserIntegerHex)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExpNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(PLANParserNull)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.Match(PLANParserBool)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Match(PLANParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(263)
			p.Match(PLANParserFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)
			p.Match(PLANParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpCsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Closure()
		}

	case 9:
		localctx = NewExpMethodInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)
			p.MethodInvoke()
		}

	case 10:
		localctx = NewExpFnInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(267)
			p.FnInvoke()
		}

	case 11:
		localctx = NewExpCsInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(268)
			p.CsInvoke()
		}

	case 12:
		localctx = NewExpNegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(269)
			p.Match(PLANParserSubtract)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.exp(13)
		}

	case 13:
		localctx = NewExpLogicalNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(271)
			p.Match(PLANParserNot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.exp(12)
		}

	case 14:
		localctx = NewExpParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)
			p.Match(PLANParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.exp(0)
		}
		{
			p.SetState(275)
			p.Match(PLANParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewExpListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.List()
		}

	case 16:
		localctx = NewExpDictContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Dict()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(307)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpPowContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(282)
					p.Match(PLANParserPow)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(283)
					p.exp(11)
				}

			case 2:
				localctx = NewExpMulDivModContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(285)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpMulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpMulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(286)
					p.exp(11)
				}

			case 3:
				localctx = NewExpSumSubContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(288)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpSumSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PLANParserAdd || _la == PLANParserSubtract) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpSumSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(289)
					p.exp(10)
				}

			case 4:
				localctx = NewExpComparisonContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(291)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&206963736576) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(292)
					p.exp(9)
				}

			case 5:
				localctx = NewExpEqualContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(294)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpEqualContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PLANParserEq || _la == PLANParserNeq) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpEqualContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(295)
					p.exp(8)
				}

			case 6:
				localctx = NewExpXorContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(297)
					p.Match(PLANParserXor)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(298)
					p.exp(7)
				}

			case 7:
				localctx = NewExpLogicalAndContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(300)
					p.Match(PLANParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(301)
					p.exp(6)
				}

			case 8:
				localctx = NewExpLogicalOrContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(303)
					p.Match(PLANParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(304)
					p.exp(5)
				}

			case 9:
				localctx = NewExpIdxContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(306)
					p.Idx()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) CopyAll(ctx *IfBlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfBlockStmtContext struct {
	IfBlockContext
}

func NewIfBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfBlockStmtContext {
	var p = new(IfBlockStmtContext)

	InitEmptyIfBlockContext(&p.IfBlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfBlockContext))

	return p
}

func (s *IfBlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockStmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IfBlockStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IfBlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIfBlockStmt(s)
	}
}

func (s *IfBlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIfBlockStmt(s)
	}
}

func (s *IfBlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIfBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PLANParserRULE_ifBlock)
	var _la int

	localctx = NewIfBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(PLANParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.exp(0)
	}
	{
		p.SetState(314)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(315)
			p.Stmt()
		}

		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(321)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElifBlockContext is an interface to support dynamic dispatch.
type IElifBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElifBlockContext differentiates from other interfaces.
	IsElifBlockContext()
}

type ElifBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifBlockContext() *ElifBlockContext {
	var p = new(ElifBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_elifBlock
	return p
}

func InitEmptyElifBlockContext(p *ElifBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_elifBlock
}

func (*ElifBlockContext) IsElifBlockContext() {}

func NewElifBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifBlockContext {
	var p = new(ElifBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_elifBlock

	return p
}

func (s *ElifBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifBlockContext) CopyAll(ctx *ElifBlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ElifBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElifBlockStmtContext struct {
	ElifBlockContext
}

func NewElifBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElifBlockStmtContext {
	var p = new(ElifBlockStmtContext)

	InitEmptyElifBlockContext(&p.ElifBlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElifBlockContext))

	return p
}

func (s *ElifBlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifBlockStmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ElifBlockStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElifBlockStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElifBlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterElifBlockStmt(s)
	}
}

func (s *ElifBlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitElifBlockStmt(s)
	}
}

func (s *ElifBlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitElifBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) ElifBlock() (localctx IElifBlockContext) {
	localctx = NewElifBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PLANParserRULE_elifBlock)
	var _la int

	localctx = NewElifBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(PLANParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.exp(0)
	}
	{
		p.SetState(325)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(326)
			p.Stmt()
		}

		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(332)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) CopyAll(ctx *ElseBlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseBlockStmtContext struct {
	ElseBlockContext
}

func NewElseBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseBlockStmtContext {
	var p = new(ElseBlockStmtContext)

	InitEmptyElseBlockContext(&p.ElseBlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElseBlockContext))

	return p
}

func (s *ElseBlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseBlockStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElseBlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterElseBlockStmt(s)
	}
}

func (s *ElseBlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitElseBlockStmt(s)
	}
}

func (s *ElseBlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitElseBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PLANParserRULE_elseBlock)
	var _la int

	localctx = NewElseBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(PLANParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(336)
			p.Stmt()
		}

		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(342)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfBlock() IIfBlockContext
	AllElifBlock() []IElifBlockContext
	ElifBlock(i int) IElifBlockContext
	ElseBlock() IElseBlockContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfStmtContext) AllElifBlock() []IElifBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElifBlockContext); ok {
			len++
		}
	}

	tst := make([]IElifBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElifBlockContext); ok {
			tst[i] = t.(IElifBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) ElifBlock(i int) IElifBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifBlockContext)
}

func (s *IfStmtContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PLANParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.IfBlock()
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PLANParserT__17 {
		{
			p.SetState(345)
			p.ElifBlock()
		}

		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PLANParserT__18 {
		{
			p.SetState(351)
			p.ElseBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnParamsContext is an interface to support dynamic dispatch.
type IFnParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	// IsFnParamsContext differentiates from other interfaces.
	IsFnParamsContext()
}

type FnParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnParamsContext() *FnParamsContext {
	var p = new(FnParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fnParams
	return p
}

func InitEmptyFnParamsContext(p *FnParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fnParams
}

func (*FnParamsContext) IsFnParamsContext() {}

func NewFnParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnParamsContext {
	var p = new(FnParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_fnParams

	return p
}

func (s *FnParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FnParamsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(PLANParserIdentifier)
}

func (s *FnParamsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, i)
}

func (s *FnParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterFnParams(s)
	}
}

func (s *FnParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitFnParams(s)
	}
}

func (s *FnParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitFnParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) FnParams() (localctx IFnParamsContext) {
	localctx = NewFnParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PLANParserRULE_fnParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(PLANParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PLANParserT__10 {
		{
			p.SetState(355)
			p.Match(PLANParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(PLANParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnBodyContext is an interface to support dynamic dispatch.
type IFnBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FnParams() IFnParamsContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsFnBodyContext differentiates from other interfaces.
	IsFnBodyContext()
}

type FnBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnBodyContext() *FnBodyContext {
	var p = new(FnBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fnBody
	return p
}

func InitEmptyFnBodyContext(p *FnBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fnBody
}

func (*FnBodyContext) IsFnBodyContext() {}

func NewFnBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnBodyContext {
	var p = new(FnBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_fnBody

	return p
}

func (s *FnBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FnBodyContext) FnParams() IFnParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnParamsContext)
}

func (s *FnBodyContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FnBodyContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FnBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterFnBody(s)
	}
}

func (s *FnBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitFnBody(s)
	}
}

func (s *FnBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitFnBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) FnBody() (localctx IFnBodyContext) {
	localctx = NewFnBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PLANParserRULE_fnBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PLANParserIdentifier {
		{
			p.SetState(363)
			p.FnParams()
		}

	}
	{
		p.SetState(366)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(368)
			p.Stmt()
		}

		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(374)
		p.Match(PLANParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnContext is an interface to support dynamic dispatch.
type IFnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	FnBody() IFnBodyContext
	Identifier() antlr.TerminalNode

	// IsFnContext differentiates from other interfaces.
	IsFnContext()
}

type FnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyFnContext() *FnContext {
	var p = new(FnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fn
	return p
}

func InitEmptyFnContext(p *FnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_fn
}

func (*FnContext) IsFnContext() {}

func NewFnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnContext {
	var p = new(FnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_fn

	return p
}

func (s *FnContext) GetParser() antlr.Parser { return s.parser }

func (s *FnContext) GetName() antlr.Token { return s.name }

func (s *FnContext) SetName(v antlr.Token) { s.name = v }

func (s *FnContext) FnBody() IFnBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnBodyContext)
}

func (s *FnContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *FnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterFn(s)
	}
}

func (s *FnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitFn(s)
	}
}

func (s *FnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitFn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Fn() (localctx IFnContext) {
	localctx = NewFnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PLANParserRULE_fn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(PLANParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*FnContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.FnBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosureContext is an interface to support dynamic dispatch.
type IClosureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FnBody() IFnBodyContext

	// IsClosureContext differentiates from other interfaces.
	IsClosureContext()
}

type ClosureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosureContext() *ClosureContext {
	var p = new(ClosureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_closure
	return p
}

func InitEmptyClosureContext(p *ClosureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_closure
}

func (*ClosureContext) IsClosureContext() {}

func NewClosureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosureContext {
	var p = new(ClosureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_closure

	return p
}

func (s *ClosureContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosureContext) FnBody() IFnBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnBodyContext)
}

func (s *ClosureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterClosure(s)
	}
}

func (s *ClosureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitClosure(s)
	}
}

func (s *ClosureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitClosure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Closure() (localctx IClosureContext) {
	localctx = NewClosureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PLANParserRULE_closure)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(PLANParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.FnBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_include
	return p
}

func InitEmptyIncludeContext(p *IncludeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_include
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (s *IncludeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitInclude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PLANParserRULE_include)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(PLANParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.exp(0)
	}
	{
		p.SetState(386)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Match(PLANParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PLANParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PLANParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
