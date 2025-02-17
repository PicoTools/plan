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
		"breakStmt", "assignment", "list", "dictUnit", "dict", "idx", "idxs",
		"methodInvoke", "fnInvoke", "csInvoke", "exp", "ifBlock", "elifBlock",
		"elseBlock", "ifStmt", "fnParams", "fnBody", "fn", "closure", "include",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 53, 399, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 0, 5, 0, 56, 8, 0, 10, 0, 12, 0, 59, 9, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 87,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 93, 8, 2, 10, 2, 12, 2, 96, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 108, 8,
		3, 10, 3, 12, 3, 111, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 117, 8, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3,
		7, 179, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 185, 8, 8, 10, 8, 12, 8, 188,
		9, 8, 3, 8, 190, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 5, 10, 202, 8, 10, 10, 10, 12, 10, 205, 9, 10, 3, 10, 207,
		8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 217,
		8, 12, 10, 12, 12, 12, 220, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 5, 13, 229, 8, 13, 10, 13, 12, 13, 232, 9, 13, 3, 13, 234, 8,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 243, 8, 14,
		10, 14, 12, 14, 246, 9, 14, 3, 14, 248, 8, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 258, 8, 15, 10, 15, 12, 15, 261,
		9, 15, 3, 15, 263, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 289, 8, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 317, 8, 16, 10, 16, 12, 16, 320,
		9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 326, 8, 17, 10, 17, 12, 17, 329,
		9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 337, 8, 18, 10,
		18, 12, 18, 340, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 347,
		8, 19, 10, 19, 12, 19, 350, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 356,
		8, 20, 10, 20, 12, 20, 359, 9, 20, 1, 20, 3, 20, 362, 8, 20, 1, 21, 1,
		21, 1, 21, 5, 21, 367, 8, 21, 10, 21, 12, 21, 370, 9, 21, 1, 22, 1, 22,
		3, 22, 374, 8, 22, 1, 22, 1, 22, 1, 22, 5, 22, 379, 8, 22, 10, 22, 12,
		22, 382, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 0, 1, 32, 26, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 0, 4, 1, 0, 38, 40, 1, 0, 41, 42, 2, 0, 28, 29,
		36, 37, 1, 0, 23, 24, 443, 0, 57, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 88,
		1, 0, 0, 0, 6, 99, 1, 0, 0, 0, 8, 114, 1, 0, 0, 0, 10, 118, 1, 0, 0, 0,
		12, 120, 1, 0, 0, 0, 14, 178, 1, 0, 0, 0, 16, 180, 1, 0, 0, 0, 18, 193,
		1, 0, 0, 0, 20, 197, 1, 0, 0, 0, 22, 210, 1, 0, 0, 0, 24, 214, 1, 0, 0,
		0, 26, 221, 1, 0, 0, 0, 28, 237, 1, 0, 0, 0, 30, 251, 1, 0, 0, 0, 32, 288,
		1, 0, 0, 0, 34, 321, 1, 0, 0, 0, 36, 332, 1, 0, 0, 0, 38, 343, 1, 0, 0,
		0, 40, 353, 1, 0, 0, 0, 42, 363, 1, 0, 0, 0, 44, 371, 1, 0, 0, 0, 46, 385,
		1, 0, 0, 0, 48, 389, 1, 0, 0, 0, 50, 392, 1, 0, 0, 0, 52, 56, 3, 2, 1,
		0, 53, 56, 3, 46, 23, 0, 54, 56, 3, 50, 25, 0, 55, 52, 1, 0, 0, 0, 55,
		53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0,
		0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61,
		5, 0, 0, 1, 61, 1, 1, 0, 0, 0, 62, 63, 3, 14, 7, 0, 63, 64, 5, 1, 0, 0,
		64, 87, 1, 0, 0, 0, 65, 66, 3, 26, 13, 0, 66, 67, 5, 1, 0, 0, 67, 87, 1,
		0, 0, 0, 68, 69, 3, 30, 15, 0, 69, 70, 5, 1, 0, 0, 70, 87, 1, 0, 0, 0,
		71, 72, 3, 28, 14, 0, 72, 73, 5, 1, 0, 0, 73, 87, 1, 0, 0, 0, 74, 75, 3,
		12, 6, 0, 75, 76, 5, 1, 0, 0, 76, 87, 1, 0, 0, 0, 77, 78, 3, 10, 5, 0,
		78, 79, 5, 1, 0, 0, 79, 87, 1, 0, 0, 0, 80, 81, 3, 8, 4, 0, 81, 82, 5,
		1, 0, 0, 82, 87, 1, 0, 0, 0, 83, 87, 3, 4, 2, 0, 84, 87, 3, 6, 3, 0, 85,
		87, 3, 40, 20, 0, 86, 62, 1, 0, 0, 0, 86, 65, 1, 0, 0, 0, 86, 68, 1, 0,
		0, 0, 86, 71, 1, 0, 0, 0, 86, 74, 1, 0, 0, 0, 86, 77, 1, 0, 0, 0, 86, 80,
		1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0,
		87, 3, 1, 0, 0, 0, 88, 89, 5, 2, 0, 0, 89, 90, 3, 32, 16, 0, 90, 94, 5,
		3, 0, 0, 91, 93, 3, 2, 1, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94,
		92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0,
		0, 97, 98, 5, 4, 0, 0, 98, 5, 1, 0, 0, 0, 99, 100, 5, 5, 0, 0, 100, 101,
		3, 14, 7, 0, 101, 102, 5, 1, 0, 0, 102, 103, 3, 32, 16, 0, 103, 104, 5,
		1, 0, 0, 104, 105, 3, 14, 7, 0, 105, 109, 5, 3, 0, 0, 106, 108, 3, 2, 1,
		0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109,
		110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113,
		5, 4, 0, 0, 113, 7, 1, 0, 0, 0, 114, 116, 5, 6, 0, 0, 115, 117, 3, 32,
		16, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 9, 1, 0, 0, 0,
		118, 119, 5, 7, 0, 0, 119, 11, 1, 0, 0, 0, 120, 121, 5, 8, 0, 0, 121, 13,
		1, 0, 0, 0, 122, 123, 5, 48, 0, 0, 123, 124, 5, 9, 0, 0, 124, 179, 3, 32,
		16, 0, 125, 126, 5, 48, 0, 0, 126, 127, 5, 30, 0, 0, 127, 179, 3, 32, 16,
		0, 128, 129, 5, 48, 0, 0, 129, 130, 5, 31, 0, 0, 130, 179, 3, 32, 16, 0,
		131, 132, 5, 48, 0, 0, 132, 133, 5, 32, 0, 0, 133, 179, 3, 32, 16, 0, 134,
		135, 5, 48, 0, 0, 135, 136, 5, 33, 0, 0, 136, 179, 3, 32, 16, 0, 137, 138,
		5, 48, 0, 0, 138, 139, 5, 34, 0, 0, 139, 179, 3, 32, 16, 0, 140, 141, 5,
		48, 0, 0, 141, 142, 5, 35, 0, 0, 142, 179, 3, 32, 16, 0, 143, 144, 5, 48,
		0, 0, 144, 145, 3, 24, 12, 0, 145, 146, 5, 9, 0, 0, 146, 147, 3, 32, 16,
		0, 147, 179, 1, 0, 0, 0, 148, 149, 5, 48, 0, 0, 149, 150, 3, 24, 12, 0,
		150, 151, 5, 30, 0, 0, 151, 152, 3, 32, 16, 0, 152, 179, 1, 0, 0, 0, 153,
		154, 5, 48, 0, 0, 154, 155, 3, 24, 12, 0, 155, 156, 5, 31, 0, 0, 156, 157,
		3, 32, 16, 0, 157, 179, 1, 0, 0, 0, 158, 159, 5, 48, 0, 0, 159, 160, 3,
		24, 12, 0, 160, 161, 5, 32, 0, 0, 161, 162, 3, 32, 16, 0, 162, 179, 1,
		0, 0, 0, 163, 164, 5, 48, 0, 0, 164, 165, 3, 24, 12, 0, 165, 166, 5, 33,
		0, 0, 166, 167, 3, 32, 16, 0, 167, 179, 1, 0, 0, 0, 168, 169, 5, 48, 0,
		0, 169, 170, 3, 24, 12, 0, 170, 171, 5, 34, 0, 0, 171, 172, 3, 32, 16,
		0, 172, 179, 1, 0, 0, 0, 173, 174, 5, 48, 0, 0, 174, 175, 3, 24, 12, 0,
		175, 176, 5, 35, 0, 0, 176, 177, 3, 32, 16, 0, 177, 179, 1, 0, 0, 0, 178,
		122, 1, 0, 0, 0, 178, 125, 1, 0, 0, 0, 178, 128, 1, 0, 0, 0, 178, 131,
		1, 0, 0, 0, 178, 134, 1, 0, 0, 0, 178, 137, 1, 0, 0, 0, 178, 140, 1, 0,
		0, 0, 178, 143, 1, 0, 0, 0, 178, 148, 1, 0, 0, 0, 178, 153, 1, 0, 0, 0,
		178, 158, 1, 0, 0, 0, 178, 163, 1, 0, 0, 0, 178, 168, 1, 0, 0, 0, 178,
		173, 1, 0, 0, 0, 179, 15, 1, 0, 0, 0, 180, 189, 5, 10, 0, 0, 181, 186,
		3, 32, 16, 0, 182, 183, 5, 11, 0, 0, 183, 185, 3, 32, 16, 0, 184, 182,
		1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 181, 1, 0, 0, 0,
		189, 190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 192, 5, 12, 0, 0, 192,
		17, 1, 0, 0, 0, 193, 194, 3, 32, 16, 0, 194, 195, 5, 13, 0, 0, 195, 196,
		3, 32, 16, 0, 196, 19, 1, 0, 0, 0, 197, 206, 5, 3, 0, 0, 198, 203, 3, 18,
		9, 0, 199, 200, 5, 11, 0, 0, 200, 202, 3, 18, 9, 0, 201, 199, 1, 0, 0,
		0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204,
		207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 198, 1, 0, 0, 0, 206, 207,
		1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 5, 4, 0, 0, 209, 21, 1, 0,
		0, 0, 210, 211, 5, 10, 0, 0, 211, 212, 3, 32, 16, 0, 212, 213, 5, 12, 0,
		0, 213, 23, 1, 0, 0, 0, 214, 218, 3, 22, 11, 0, 215, 217, 3, 22, 11, 0,
		216, 215, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218,
		219, 1, 0, 0, 0, 219, 25, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 222, 5,
		48, 0, 0, 222, 223, 5, 14, 0, 0, 223, 224, 5, 48, 0, 0, 224, 233, 5, 15,
		0, 0, 225, 230, 3, 32, 16, 0, 226, 227, 5, 11, 0, 0, 227, 229, 3, 32, 16,
		0, 228, 226, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230,
		231, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 225,
		1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 5, 16,
		0, 0, 236, 27, 1, 0, 0, 0, 237, 238, 5, 48, 0, 0, 238, 247, 5, 15, 0, 0,
		239, 244, 3, 32, 16, 0, 240, 241, 5, 11, 0, 0, 241, 243, 3, 32, 16, 0,
		242, 240, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 247, 239,
		1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 5, 16,
		0, 0, 250, 29, 1, 0, 0, 0, 251, 252, 5, 45, 0, 0, 252, 253, 5, 48, 0, 0,
		253, 262, 5, 15, 0, 0, 254, 259, 3, 32, 16, 0, 255, 256, 5, 11, 0, 0, 256,
		258, 3, 32, 16, 0, 257, 255, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257,
		1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0,
		0, 0, 262, 254, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0,
		264, 265, 5, 16, 0, 0, 265, 31, 1, 0, 0, 0, 266, 267, 6, 16, -1, 0, 267,
		289, 5, 49, 0, 0, 268, 289, 5, 50, 0, 0, 269, 289, 5, 47, 0, 0, 270, 289,
		5, 46, 0, 0, 271, 289, 5, 48, 0, 0, 272, 289, 5, 51, 0, 0, 273, 289, 5,
		52, 0, 0, 274, 289, 3, 48, 24, 0, 275, 289, 3, 26, 13, 0, 276, 289, 3,
		28, 14, 0, 277, 289, 3, 30, 15, 0, 278, 279, 5, 42, 0, 0, 279, 289, 3,
		32, 16, 13, 280, 281, 5, 44, 0, 0, 281, 289, 3, 32, 16, 12, 282, 283, 5,
		15, 0, 0, 283, 284, 3, 32, 16, 0, 284, 285, 5, 16, 0, 0, 285, 289, 1, 0,
		0, 0, 286, 289, 3, 16, 8, 0, 287, 289, 3, 20, 10, 0, 288, 266, 1, 0, 0,
		0, 288, 268, 1, 0, 0, 0, 288, 269, 1, 0, 0, 0, 288, 270, 1, 0, 0, 0, 288,
		271, 1, 0, 0, 0, 288, 272, 1, 0, 0, 0, 288, 273, 1, 0, 0, 0, 288, 274,
		1, 0, 0, 0, 288, 275, 1, 0, 0, 0, 288, 276, 1, 0, 0, 0, 288, 277, 1, 0,
		0, 0, 288, 278, 1, 0, 0, 0, 288, 280, 1, 0, 0, 0, 288, 282, 1, 0, 0, 0,
		288, 286, 1, 0, 0, 0, 288, 287, 1, 0, 0, 0, 289, 318, 1, 0, 0, 0, 290,
		291, 10, 11, 0, 0, 291, 292, 5, 27, 0, 0, 292, 317, 3, 32, 16, 11, 293,
		294, 10, 10, 0, 0, 294, 295, 7, 0, 0, 0, 295, 317, 3, 32, 16, 11, 296,
		297, 10, 9, 0, 0, 297, 298, 7, 1, 0, 0, 298, 317, 3, 32, 16, 10, 299, 300,
		10, 8, 0, 0, 300, 301, 7, 2, 0, 0, 301, 317, 3, 32, 16, 9, 302, 303, 10,
		7, 0, 0, 303, 304, 7, 3, 0, 0, 304, 317, 3, 32, 16, 8, 305, 306, 10, 6,
		0, 0, 306, 307, 5, 43, 0, 0, 307, 317, 3, 32, 16, 7, 308, 309, 10, 5, 0,
		0, 309, 310, 5, 26, 0, 0, 310, 317, 3, 32, 16, 6, 311, 312, 10, 4, 0, 0,
		312, 313, 5, 25, 0, 0, 313, 317, 3, 32, 16, 5, 314, 315, 10, 14, 0, 0,
		315, 317, 3, 22, 11, 0, 316, 290, 1, 0, 0, 0, 316, 293, 1, 0, 0, 0, 316,
		296, 1, 0, 0, 0, 316, 299, 1, 0, 0, 0, 316, 302, 1, 0, 0, 0, 316, 305,
		1, 0, 0, 0, 316, 308, 1, 0, 0, 0, 316, 311, 1, 0, 0, 0, 316, 314, 1, 0,
		0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0,
		319, 33, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 322, 5, 17, 0, 0, 322,
		323, 3, 32, 16, 0, 323, 327, 5, 3, 0, 0, 324, 326, 3, 2, 1, 0, 325, 324,
		1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327, 328, 1, 0,
		0, 0, 328, 330, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 331, 5, 4, 0, 0,
		331, 35, 1, 0, 0, 0, 332, 333, 5, 18, 0, 0, 333, 334, 3, 32, 16, 0, 334,
		338, 5, 3, 0, 0, 335, 337, 3, 2, 1, 0, 336, 335, 1, 0, 0, 0, 337, 340,
		1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 341, 1, 0,
		0, 0, 340, 338, 1, 0, 0, 0, 341, 342, 5, 4, 0, 0, 342, 37, 1, 0, 0, 0,
		343, 344, 5, 19, 0, 0, 344, 348, 5, 3, 0, 0, 345, 347, 3, 2, 1, 0, 346,
		345, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349,
		1, 0, 0, 0, 349, 351, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351, 352, 5, 4,
		0, 0, 352, 39, 1, 0, 0, 0, 353, 357, 3, 34, 17, 0, 354, 356, 3, 36, 18,
		0, 355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357,
		358, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 362,
		3, 38, 19, 0, 361, 360, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 41, 1, 0,
		0, 0, 363, 368, 5, 48, 0, 0, 364, 365, 5, 11, 0, 0, 365, 367, 5, 48, 0,
		0, 366, 364, 1, 0, 0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368,
		369, 1, 0, 0, 0, 369, 43, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 373, 5,
		15, 0, 0, 372, 374, 3, 42, 21, 0, 373, 372, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 5, 16, 0, 0, 376, 380, 5, 3, 0, 0,
		377, 379, 3, 2, 1, 0, 378, 377, 1, 0, 0, 0, 379, 382, 1, 0, 0, 0, 380,
		378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 383, 1, 0, 0, 0, 382, 380,
		1, 0, 0, 0, 383, 384, 5, 4, 0, 0, 384, 45, 1, 0, 0, 0, 385, 386, 5, 20,
		0, 0, 386, 387, 5, 48, 0, 0, 387, 388, 3, 44, 22, 0, 388, 47, 1, 0, 0,
		0, 389, 390, 5, 20, 0, 0, 390, 391, 3, 44, 22, 0, 391, 49, 1, 0, 0, 0,
		392, 393, 5, 21, 0, 0, 393, 394, 5, 15, 0, 0, 394, 395, 3, 32, 16, 0, 395,
		396, 5, 16, 0, 0, 396, 397, 5, 1, 0, 0, 397, 51, 1, 0, 0, 0, 29, 55, 57,
		86, 94, 109, 116, 178, 186, 189, 203, 206, 218, 230, 233, 244, 247, 259,
		262, 288, 316, 318, 327, 338, 348, 357, 361, 368, 373, 380,
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
	PLANParserRULE_idxs         = 12
	PLANParserRULE_methodInvoke = 13
	PLANParserRULE_fnInvoke     = 14
	PLANParserRULE_csInvoke     = 15
	PLANParserRULE_exp          = 16
	PLANParserRULE_ifBlock      = 17
	PLANParserRULE_elifBlock    = 18
	PLANParserRULE_elseBlock    = 19
	PLANParserRULE_ifStmt       = 20
	PLANParserRULE_fnParams     = 21
	PLANParserRULE_fnBody       = 22
	PLANParserRULE_fn           = 23
	PLANParserRULE_closure      = 24
	PLANParserRULE_include      = 25
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659352076772) != 0 {
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PLANParserT__1, PLANParserT__4, PLANParserT__5, PLANParserT__6, PLANParserT__7, PLANParserT__16, PLANParserClosure, PLANParserIdentifier:
			{
				p.SetState(52)
				p.Stmt()
			}

		case PLANParserT__19:
			{
				p.SetState(53)
				p.Fn()
			}

		case PLANParserT__20:
			{
				p.SetState(54)
				p.Include()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Assignment()
		}
		{
			p.SetState(63)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.MethodInvoke()
		}
		{
			p.SetState(66)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.CsInvoke()
		}
		{
			p.SetState(69)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.FnInvoke()
		}
		{
			p.SetState(72)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(74)
			p.BreakStmt()
		}
		{
			p.SetState(75)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.ContinueStmt()
		}
		{
			p.SetState(78)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(80)
			p.ReturnStmt()
		}
		{
			p.SetState(81)
			p.Match(PLANParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(83)
			p.WhileStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(84)
			p.ForStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(85)
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
		p.SetState(88)
		p.Match(PLANParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.exp(0)
	}
	{
		p.SetState(90)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(91)
			p.Stmt()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
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
		p.SetState(99)
		p.Match(PLANParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Assignment()
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
		p.exp(0)
	}
	{
		p.SetState(103)
		p.Match(PLANParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Assignment()
	}
	{
		p.SetState(105)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(106)
			p.Stmt()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
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
		p.SetState(114)
		p.Match(PLANParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(115)
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
		p.SetState(118)
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
		p.SetState(120)
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

type AssignIdxsRegularContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsRegularContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsRegularContext {
	var p = new(AssignIdxsRegularContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsRegularContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsRegularContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsRegularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsRegularContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsRegularContext) Exp() IExpContext {
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

func (s *AssignIdxsRegularContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsRegularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsRegular(s)
	}
}

func (s *AssignIdxsRegularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsRegular(s)
	}
}

func (s *AssignIdxsRegularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsRegular(s)

	default:
		return t.VisitChildren(s)
	}
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

type AssignIdxsPowContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsPowContext {
	var p = new(AssignIdxsPowContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsPowContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsPowContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsPowContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsPowContext) AssPow() antlr.TerminalNode {
	return s.GetToken(PLANParserAssPow, 0)
}

func (s *AssignIdxsPowContext) Exp() IExpContext {
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

func (s *AssignIdxsPowContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsPowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsPow(s)
	}
}

func (s *AssignIdxsPowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsPow(s)
	}
}

func (s *AssignIdxsPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxsModContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsModContext {
	var p = new(AssignIdxsModContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsModContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsModContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsModContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsModContext) AssMod() antlr.TerminalNode {
	return s.GetToken(PLANParserAssMod, 0)
}

func (s *AssignIdxsModContext) Exp() IExpContext {
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

func (s *AssignIdxsModContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsMod(s)
	}
}

func (s *AssignIdxsModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsMod(s)
	}
}

func (s *AssignIdxsModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsMod(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIdxsMulContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsMulContext {
	var p = new(AssignIdxsMulContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsMulContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsMulContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsMulContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsMulContext) AssMul() antlr.TerminalNode {
	return s.GetToken(PLANParserAssMul, 0)
}

func (s *AssignIdxsMulContext) Exp() IExpContext {
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

func (s *AssignIdxsMulContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsMul(s)
	}
}

func (s *AssignIdxsMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsMul(s)
	}
}

func (s *AssignIdxsMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsMul(s)

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

type AssignIdxsSubContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsSubContext {
	var p = new(AssignIdxsSubContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsSubContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsSubContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsSubContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsSubContext) AssSub() antlr.TerminalNode {
	return s.GetToken(PLANParserAssSub, 0)
}

func (s *AssignIdxsSubContext) Exp() IExpContext {
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

func (s *AssignIdxsSubContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsSub(s)
	}
}

func (s *AssignIdxsSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsSub(s)
	}
}

func (s *AssignIdxsSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsSub(s)

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

type AssignIdxsSumContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsSumContext {
	var p = new(AssignIdxsSumContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsSumContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsSumContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsSumContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsSumContext) AssSum() antlr.TerminalNode {
	return s.GetToken(PLANParserAssSum, 0)
}

func (s *AssignIdxsSumContext) Exp() IExpContext {
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

func (s *AssignIdxsSumContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsSum(s)
	}
}

func (s *AssignIdxsSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsSum(s)
	}
}

func (s *AssignIdxsSumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsSum(s)

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

type AssignIdxsDivContext struct {
	AssignmentContext
	name antlr.Token
}

func NewAssignIdxsDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIdxsDivContext {
	var p = new(AssignIdxsDivContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignIdxsDivContext) GetName() antlr.Token { return s.name }

func (s *AssignIdxsDivContext) SetName(v antlr.Token) { s.name = v }

func (s *AssignIdxsDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIdxsDivContext) Idxs() IIdxsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdxsContext)
}

func (s *AssignIdxsDivContext) AssDiv() antlr.TerminalNode {
	return s.GetToken(PLANParserAssDiv, 0)
}

func (s *AssignIdxsDivContext) Exp() IExpContext {
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

func (s *AssignIdxsDivContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PLANParserIdentifier, 0)
}

func (s *AssignIdxsDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterAssignIdxsDiv(s)
	}
}

func (s *AssignIdxsDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitAssignIdxsDiv(s)
	}
}

func (s *AssignIdxsDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitAssignIdxsDiv(s)

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

func (p *PLANParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PLANParserRULE_assignment)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignRegularContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignRegularContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(PLANParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.exp(0)
		}

	case 2:
		localctx = NewAssignSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignSumContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(PLANParserAssSum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.exp(0)
		}

	case 3:
		localctx = NewAssignSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignSubContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(PLANParserAssSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.exp(0)
		}

	case 4:
		localctx = NewAssignMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignMulContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(PLANParserAssMul)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.exp(0)
		}

	case 5:
		localctx = NewAssignDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(134)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignDivContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(PLANParserAssDiv)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.exp(0)
		}

	case 6:
		localctx = NewAssignModContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(137)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignModContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(PLANParserAssMod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.exp(0)
		}

	case 7:
		localctx = NewAssignPowContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(140)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignPowContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(PLANParserAssPow)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.exp(0)
		}

	case 8:
		localctx = NewAssignIdxsRegularContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(143)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsRegularContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Idxs()
		}
		{
			p.SetState(145)
			p.Match(PLANParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.exp(0)
		}

	case 9:
		localctx = NewAssignIdxsSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(148)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsSumContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Idxs()
		}
		{
			p.SetState(150)
			p.Match(PLANParserAssSum)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.exp(0)
		}

	case 10:
		localctx = NewAssignIdxsSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(153)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsSubContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Idxs()
		}
		{
			p.SetState(155)
			p.Match(PLANParserAssSub)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.exp(0)
		}

	case 11:
		localctx = NewAssignIdxsMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(158)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsMulContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Idxs()
		}
		{
			p.SetState(160)
			p.Match(PLANParserAssMul)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.exp(0)
		}

	case 12:
		localctx = NewAssignIdxsDivContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(163)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsDivContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Idxs()
		}
		{
			p.SetState(165)
			p.Match(PLANParserAssDiv)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.exp(0)
		}

	case 13:
		localctx = NewAssignIdxsModContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(168)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsModContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Idxs()
		}
		{
			p.SetState(170)
			p.Match(PLANParserAssMod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.exp(0)
		}

	case 14:
		localctx = NewAssignIdxsPowContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(173)

			var _m = p.Match(PLANParserIdentifier)

			localctx.(*AssignIdxsPowContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Idxs()
		}
		{
			p.SetState(175)
			p.Match(PLANParserAssPow)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
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
		p.SetState(180)
		p.Match(PLANParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
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

		for _la == PLANParserT__10 {
			{
				p.SetState(182)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(183)
				p.exp(0)
			}

			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(191)
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
		p.SetState(193)
		p.exp(0)
	}
	{
		p.SetState(194)
		p.Match(PLANParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
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
		p.SetState(197)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
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

		for _la == PLANParserT__10 {
			{
				p.SetState(199)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.DictUnit()
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(208)
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
		p.SetState(210)
		p.Match(PLANParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.exp(0)
	}
	{
		p.SetState(212)
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

// IIdxsContext is an interface to support dynamic dispatch.
type IIdxsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdx() []IIdxContext
	Idx(i int) IIdxContext

	// IsIdxsContext differentiates from other interfaces.
	IsIdxsContext()
}

type IdxsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdxsContext() *IdxsContext {
	var p = new(IdxsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_idxs
	return p
}

func InitEmptyIdxsContext(p *IdxsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PLANParserRULE_idxs
}

func (*IdxsContext) IsIdxsContext() {}

func NewIdxsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdxsContext {
	var p = new(IdxsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLANParserRULE_idxs

	return p
}

func (s *IdxsContext) GetParser() antlr.Parser { return s.parser }

func (s *IdxsContext) AllIdx() []IIdxContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdxContext); ok {
			len++
		}
	}

	tst := make([]IIdxContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdxContext); ok {
			tst[i] = t.(IIdxContext)
			i++
		}
	}

	return tst
}

func (s *IdxsContext) Idx(i int) IIdxContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdxContext); ok {
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

	return t.(IIdxContext)
}

func (s *IdxsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdxsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdxsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.EnterIdxs(s)
	}
}

func (s *IdxsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLANListener); ok {
		listenerT.ExitIdxs(s)
	}
}

func (s *IdxsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PLANVisitor:
		return t.VisitIdxs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PLANParser) Idxs() (localctx IIdxsContext) {
	localctx = NewIdxsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PLANParserRULE_idxs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Idx()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PLANParserT__9 {
		{
			p.SetState(215)
			p.Idx()
		}

		p.SetState(220)
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
	p.EnterRule(localctx, 26, PLANParserRULE_methodInvoke)
	var _la int

	localctx = NewIdentifierMethodInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierMethodInvokeContext).var_ = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(PLANParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierMethodInvokeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(225)
			p.exp(0)
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(226)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(227)
				p.exp(0)
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(235)
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
	p.EnterRule(localctx, 28, PLANParserRULE_fnInvoke)
	var _la int

	localctx = NewIdentifierFnInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierFnInvokeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(239)
			p.exp(0)
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(240)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(241)
				p.exp(0)
			}

			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(249)
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
	p.EnterRule(localctx, 30, PLANParserRULE_csInvoke)
	var _la int

	localctx = NewIdentifierCsInvokeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(PLANParserClosure)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*IdentifierCsInvokeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8994005116290056) != 0 {
		{
			p.SetState(254)
			p.exp(0)
		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PLANParserT__10 {
			{
				p.SetState(255)
				p.Match(PLANParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(256)
				p.exp(0)
			}

			p.SetState(261)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(264)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, PLANParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(267)
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
			p.SetState(268)
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
			p.SetState(269)
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
			p.SetState(270)
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
			p.SetState(271)
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
			p.SetState(272)
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
			p.SetState(273)
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
			p.SetState(274)
			p.Closure()
		}

	case 9:
		localctx = NewExpMethodInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			p.MethodInvoke()
		}

	case 10:
		localctx = NewExpFnInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(276)
			p.FnInvoke()
		}

	case 11:
		localctx = NewExpCsInvokeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.CsInvoke()
		}

	case 12:
		localctx = NewExpNegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Match(PLANParserSubtract)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.exp(13)
		}

	case 13:
		localctx = NewExpLogicalNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(280)
			p.Match(PLANParserNot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.exp(12)
		}

	case 14:
		localctx = NewExpParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(282)
			p.Match(PLANParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.exp(0)
		}
		{
			p.SetState(284)
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
			p.SetState(286)
			p.List()
		}

	case 16:
		localctx = NewExpDictContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(287)
			p.Dict()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpPowContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(291)
					p.Match(PLANParserPow)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(292)
					p.exp(11)
				}

			case 2:
				localctx = NewExpMulDivModContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(294)

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
					p.SetState(295)
					p.exp(11)
				}

			case 3:
				localctx = NewExpSumSubContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(297)

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
					p.SetState(298)
					p.exp(10)
				}

			case 4:
				localctx = NewExpComparisonContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(300)

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
					p.SetState(301)
					p.exp(9)
				}

			case 5:
				localctx = NewExpEqualContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(303)

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
					p.SetState(304)
					p.exp(8)
				}

			case 6:
				localctx = NewExpXorContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(306)
					p.Match(PLANParserXor)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(307)
					p.exp(7)
				}

			case 7:
				localctx = NewExpLogicalAndContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(309)
					p.Match(PLANParserAnd)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(310)
					p.exp(6)
				}

			case 8:
				localctx = NewExpLogicalOrContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(312)
					p.Match(PLANParserOr)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(313)
					p.exp(5)
				}

			case 9:
				localctx = NewExpIdxContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLANParserRULE_exp)
				p.SetState(314)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(315)
					p.Idx()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, PLANParserRULE_ifBlock)
	var _la int

	localctx = NewIfBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(PLANParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.exp(0)
	}
	{
		p.SetState(323)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(324)
			p.Stmt()
		}

		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(330)
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
	p.EnterRule(localctx, 36, PLANParserRULE_elifBlock)
	var _la int

	localctx = NewElifBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(PLANParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.exp(0)
	}
	{
		p.SetState(334)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(335)
			p.Stmt()
		}

		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(341)
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
	p.EnterRule(localctx, 38, PLANParserRULE_elseBlock)
	var _la int

	localctx = NewElseBlockStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(PLANParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(345)
			p.Stmt()
		}

		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(351)
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
	p.EnterRule(localctx, 40, PLANParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.IfBlock()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PLANParserT__17 {
		{
			p.SetState(354)
			p.ElifBlock()
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PLANParserT__18 {
		{
			p.SetState(360)
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
	p.EnterRule(localctx, 42, PLANParserRULE_fnParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(PLANParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PLANParserT__10 {
		{
			p.SetState(364)
			p.Match(PLANParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(PLANParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(370)
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
	p.EnterRule(localctx, 44, PLANParserRULE_fnBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PLANParserIdentifier {
		{
			p.SetState(372)
			p.FnParams()
		}

	}
	{
		p.SetState(375)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Match(PLANParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&316659348931044) != 0 {
		{
			p.SetState(377)
			p.Stmt()
		}

		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(383)
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
	p.EnterRule(localctx, 46, PLANParserRULE_fn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(PLANParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)

		var _m = p.Match(PLANParserIdentifier)

		localctx.(*FnContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
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
	p.EnterRule(localctx, 48, PLANParserRULE_closure)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(PLANParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
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
	p.EnterRule(localctx, 50, PLANParserRULE_include)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(PLANParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(PLANParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.exp(0)
	}
	{
		p.SetState(395)
		p.Match(PLANParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
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
	case 16:
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
