// Code generated from ./grammar/PLAN.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PLANLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PLANLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func planlexerLexerInit() {
	staticData := &PLANLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'fn'", "'('", "')'", "'{'", "'}'", "','", "';'", "'while'", "'for'",
		"'if'", "'elif'", "'else'", "'include'", "'return'", "'continue'", "'break'",
		"'='", "'['", "']'", "':'", "'.'", "", "'=='", "'!='", "'||'", "'&&'",
		"'**'", "'>='", "'<='", "'+='", "'-='", "'*='", "'/='", "'%='", "'**='",
		"'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'^'", "'!'", "'@'",
		"", "'null'",
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "WS", "Eq", "Neq", "Or", "And",
		"Pow", "GtEq", "LtEq", "AssSum", "AssSub", "AssMul", "AssDiv", "AssMod",
		"AssPow", "Gt", "Lt", "Multiply", "Division", "Modulus", "Add", "Subtract",
		"Xor", "Not", "Closure", "Bool", "Null", "Identifier", "Integer", "IntegerHex",
		"Float", "String", "Comment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 53, 349, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 3, 45, 259, 8, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 47, 1, 47, 5, 47, 268, 8, 47, 10, 47, 12, 47, 271, 9, 47, 1, 48, 4,
		48, 274, 8, 48, 11, 48, 12, 48, 275, 1, 49, 1, 49, 1, 49, 4, 49, 281, 8,
		49, 11, 49, 12, 49, 282, 1, 50, 5, 50, 286, 8, 50, 10, 50, 12, 50, 289,
		9, 50, 1, 50, 1, 50, 4, 50, 293, 8, 50, 11, 50, 12, 50, 294, 1, 51, 1,
		51, 1, 51, 1, 51, 5, 51, 301, 8, 51, 10, 51, 12, 51, 304, 9, 51, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 311, 8, 51, 10, 51, 12, 51, 314, 9,
		51, 1, 51, 3, 51, 317, 8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 323, 8,
		52, 10, 52, 12, 52, 326, 9, 52, 1, 52, 1, 52, 5, 52, 330, 8, 52, 10, 52,
		12, 52, 333, 9, 52, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 339, 8, 52, 10,
		52, 12, 52, 342, 9, 52, 1, 52, 1, 52, 3, 52, 346, 8, 52, 1, 52, 1, 52,
		1, 340, 0, 53, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45,
		91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 1, 0,
		12, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 48, 48, 2, 0, 88, 88,
		120, 120, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 34, 34, 4, 0, 10, 10, 13,
		13, 34, 34, 92, 92, 2, 0, 10, 10, 13, 13, 1, 0, 39, 39, 4, 0, 10, 10, 13,
		13, 39, 39, 92, 92, 364, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0,
		0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0,
		0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105,
		1, 0, 0, 0, 1, 107, 1, 0, 0, 0, 3, 110, 1, 0, 0, 0, 5, 112, 1, 0, 0, 0,
		7, 114, 1, 0, 0, 0, 9, 116, 1, 0, 0, 0, 11, 118, 1, 0, 0, 0, 13, 120, 1,
		0, 0, 0, 15, 122, 1, 0, 0, 0, 17, 128, 1, 0, 0, 0, 19, 132, 1, 0, 0, 0,
		21, 135, 1, 0, 0, 0, 23, 140, 1, 0, 0, 0, 25, 145, 1, 0, 0, 0, 27, 153,
		1, 0, 0, 0, 29, 160, 1, 0, 0, 0, 31, 169, 1, 0, 0, 0, 33, 175, 1, 0, 0,
		0, 35, 177, 1, 0, 0, 0, 37, 179, 1, 0, 0, 0, 39, 181, 1, 0, 0, 0, 41, 183,
		1, 0, 0, 0, 43, 185, 1, 0, 0, 0, 45, 189, 1, 0, 0, 0, 47, 192, 1, 0, 0,
		0, 49, 195, 1, 0, 0, 0, 51, 198, 1, 0, 0, 0, 53, 201, 1, 0, 0, 0, 55, 204,
		1, 0, 0, 0, 57, 207, 1, 0, 0, 0, 59, 210, 1, 0, 0, 0, 61, 213, 1, 0, 0,
		0, 63, 216, 1, 0, 0, 0, 65, 219, 1, 0, 0, 0, 67, 222, 1, 0, 0, 0, 69, 225,
		1, 0, 0, 0, 71, 229, 1, 0, 0, 0, 73, 231, 1, 0, 0, 0, 75, 233, 1, 0, 0,
		0, 77, 235, 1, 0, 0, 0, 79, 237, 1, 0, 0, 0, 81, 239, 1, 0, 0, 0, 83, 241,
		1, 0, 0, 0, 85, 243, 1, 0, 0, 0, 87, 245, 1, 0, 0, 0, 89, 247, 1, 0, 0,
		0, 91, 258, 1, 0, 0, 0, 93, 260, 1, 0, 0, 0, 95, 265, 1, 0, 0, 0, 97, 273,
		1, 0, 0, 0, 99, 277, 1, 0, 0, 0, 101, 287, 1, 0, 0, 0, 103, 316, 1, 0,
		0, 0, 105, 345, 1, 0, 0, 0, 107, 108, 5, 102, 0, 0, 108, 109, 5, 110, 0,
		0, 109, 2, 1, 0, 0, 0, 110, 111, 5, 40, 0, 0, 111, 4, 1, 0, 0, 0, 112,
		113, 5, 41, 0, 0, 113, 6, 1, 0, 0, 0, 114, 115, 5, 123, 0, 0, 115, 8, 1,
		0, 0, 0, 116, 117, 5, 125, 0, 0, 117, 10, 1, 0, 0, 0, 118, 119, 5, 44,
		0, 0, 119, 12, 1, 0, 0, 0, 120, 121, 5, 59, 0, 0, 121, 14, 1, 0, 0, 0,
		122, 123, 5, 119, 0, 0, 123, 124, 5, 104, 0, 0, 124, 125, 5, 105, 0, 0,
		125, 126, 5, 108, 0, 0, 126, 127, 5, 101, 0, 0, 127, 16, 1, 0, 0, 0, 128,
		129, 5, 102, 0, 0, 129, 130, 5, 111, 0, 0, 130, 131, 5, 114, 0, 0, 131,
		18, 1, 0, 0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 102, 0, 0, 134, 20,
		1, 0, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 108, 0, 0, 137, 138, 5,
		105, 0, 0, 138, 139, 5, 102, 0, 0, 139, 22, 1, 0, 0, 0, 140, 141, 5, 101,
		0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5, 115, 0, 0, 143, 144, 5, 101,
		0, 0, 144, 24, 1, 0, 0, 0, 145, 146, 5, 105, 0, 0, 146, 147, 5, 110, 0,
		0, 147, 148, 5, 99, 0, 0, 148, 149, 5, 108, 0, 0, 149, 150, 5, 117, 0,
		0, 150, 151, 5, 100, 0, 0, 151, 152, 5, 101, 0, 0, 152, 26, 1, 0, 0, 0,
		153, 154, 5, 114, 0, 0, 154, 155, 5, 101, 0, 0, 155, 156, 5, 116, 0, 0,
		156, 157, 5, 117, 0, 0, 157, 158, 5, 114, 0, 0, 158, 159, 5, 110, 0, 0,
		159, 28, 1, 0, 0, 0, 160, 161, 5, 99, 0, 0, 161, 162, 5, 111, 0, 0, 162,
		163, 5, 110, 0, 0, 163, 164, 5, 116, 0, 0, 164, 165, 5, 105, 0, 0, 165,
		166, 5, 110, 0, 0, 166, 167, 5, 117, 0, 0, 167, 168, 5, 101, 0, 0, 168,
		30, 1, 0, 0, 0, 169, 170, 5, 98, 0, 0, 170, 171, 5, 114, 0, 0, 171, 172,
		5, 101, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 107, 0, 0, 174, 32, 1,
		0, 0, 0, 175, 176, 5, 61, 0, 0, 176, 34, 1, 0, 0, 0, 177, 178, 5, 91, 0,
		0, 178, 36, 1, 0, 0, 0, 179, 180, 5, 93, 0, 0, 180, 38, 1, 0, 0, 0, 181,
		182, 5, 58, 0, 0, 182, 40, 1, 0, 0, 0, 183, 184, 5, 46, 0, 0, 184, 42,
		1, 0, 0, 0, 185, 186, 7, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 6, 21,
		0, 0, 188, 44, 1, 0, 0, 0, 189, 190, 5, 61, 0, 0, 190, 191, 5, 61, 0, 0,
		191, 46, 1, 0, 0, 0, 192, 193, 5, 33, 0, 0, 193, 194, 5, 61, 0, 0, 194,
		48, 1, 0, 0, 0, 195, 196, 5, 124, 0, 0, 196, 197, 5, 124, 0, 0, 197, 50,
		1, 0, 0, 0, 198, 199, 5, 38, 0, 0, 199, 200, 5, 38, 0, 0, 200, 52, 1, 0,
		0, 0, 201, 202, 5, 42, 0, 0, 202, 203, 5, 42, 0, 0, 203, 54, 1, 0, 0, 0,
		204, 205, 5, 62, 0, 0, 205, 206, 5, 61, 0, 0, 206, 56, 1, 0, 0, 0, 207,
		208, 5, 60, 0, 0, 208, 209, 5, 61, 0, 0, 209, 58, 1, 0, 0, 0, 210, 211,
		5, 43, 0, 0, 211, 212, 5, 61, 0, 0, 212, 60, 1, 0, 0, 0, 213, 214, 5, 45,
		0, 0, 214, 215, 5, 61, 0, 0, 215, 62, 1, 0, 0, 0, 216, 217, 5, 42, 0, 0,
		217, 218, 5, 61, 0, 0, 218, 64, 1, 0, 0, 0, 219, 220, 5, 47, 0, 0, 220,
		221, 5, 61, 0, 0, 221, 66, 1, 0, 0, 0, 222, 223, 5, 37, 0, 0, 223, 224,
		5, 61, 0, 0, 224, 68, 1, 0, 0, 0, 225, 226, 5, 42, 0, 0, 226, 227, 5, 42,
		0, 0, 227, 228, 5, 61, 0, 0, 228, 70, 1, 0, 0, 0, 229, 230, 5, 62, 0, 0,
		230, 72, 1, 0, 0, 0, 231, 232, 5, 60, 0, 0, 232, 74, 1, 0, 0, 0, 233, 234,
		5, 42, 0, 0, 234, 76, 1, 0, 0, 0, 235, 236, 5, 47, 0, 0, 236, 78, 1, 0,
		0, 0, 237, 238, 5, 37, 0, 0, 238, 80, 1, 0, 0, 0, 239, 240, 5, 43, 0, 0,
		240, 82, 1, 0, 0, 0, 241, 242, 5, 45, 0, 0, 242, 84, 1, 0, 0, 0, 243, 244,
		5, 94, 0, 0, 244, 86, 1, 0, 0, 0, 245, 246, 5, 33, 0, 0, 246, 88, 1, 0,
		0, 0, 247, 248, 5, 64, 0, 0, 248, 90, 1, 0, 0, 0, 249, 250, 5, 116, 0,
		0, 250, 251, 5, 114, 0, 0, 251, 252, 5, 117, 0, 0, 252, 259, 5, 101, 0,
		0, 253, 254, 5, 102, 0, 0, 254, 255, 5, 97, 0, 0, 255, 256, 5, 108, 0,
		0, 256, 257, 5, 115, 0, 0, 257, 259, 5, 101, 0, 0, 258, 249, 1, 0, 0, 0,
		258, 253, 1, 0, 0, 0, 259, 92, 1, 0, 0, 0, 260, 261, 5, 110, 0, 0, 261,
		262, 5, 117, 0, 0, 262, 263, 5, 108, 0, 0, 263, 264, 5, 108, 0, 0, 264,
		94, 1, 0, 0, 0, 265, 269, 7, 1, 0, 0, 266, 268, 7, 2, 0, 0, 267, 266, 1,
		0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0,
		0, 270, 96, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 274, 7, 3, 0, 0, 273,
		272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 98, 1, 0, 0, 0, 277, 278, 7, 4, 0, 0, 278, 280, 7, 5,
		0, 0, 279, 281, 7, 6, 0, 0, 280, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0,
		282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 100, 1, 0, 0, 0, 284,
		286, 7, 3, 0, 0, 285, 284, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285,
		1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0, 289, 287, 1, 0,
		0, 0, 290, 292, 5, 46, 0, 0, 291, 293, 7, 3, 0, 0, 292, 291, 1, 0, 0, 0,
		293, 294, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295,
		102, 1, 0, 0, 0, 296, 302, 7, 7, 0, 0, 297, 301, 8, 8, 0, 0, 298, 299,
		5, 92, 0, 0, 299, 301, 8, 9, 0, 0, 300, 297, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0,
		303, 305, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 317, 7, 7, 0, 0, 306,
		312, 7, 10, 0, 0, 307, 311, 8, 11, 0, 0, 308, 309, 5, 92, 0, 0, 309, 311,
		8, 9, 0, 0, 310, 307, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 314, 1, 0,
		0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0,
		314, 312, 1, 0, 0, 0, 315, 317, 7, 10, 0, 0, 316, 296, 1, 0, 0, 0, 316,
		306, 1, 0, 0, 0, 317, 104, 1, 0, 0, 0, 318, 319, 5, 47, 0, 0, 319, 320,
		5, 47, 0, 0, 320, 324, 1, 0, 0, 0, 321, 323, 8, 9, 0, 0, 322, 321, 1, 0,
		0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0,
		325, 346, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 331, 5, 35, 0, 0, 328,
		330, 8, 9, 0, 0, 329, 328, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329,
		1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 346, 1, 0, 0, 0, 333, 331, 1, 0,
		0, 0, 334, 335, 5, 47, 0, 0, 335, 336, 5, 42, 0, 0, 336, 340, 1, 0, 0,
		0, 337, 339, 9, 0, 0, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340,
		341, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341, 343, 1, 0, 0, 0, 342, 340,
		1, 0, 0, 0, 343, 344, 5, 42, 0, 0, 344, 346, 5, 47, 0, 0, 345, 318, 1,
		0, 0, 0, 345, 327, 1, 0, 0, 0, 345, 334, 1, 0, 0, 0, 346, 347, 1, 0, 0,
		0, 347, 348, 6, 52, 0, 0, 348, 106, 1, 0, 0, 0, 16, 0, 258, 269, 275, 282,
		287, 294, 300, 302, 310, 312, 316, 324, 331, 340, 345, 1, 0, 1, 0,
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

// PLANLexerInit initializes any static state used to implement PLANLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPLANLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PLANLexerInit() {
	staticData := &PLANLexerLexerStaticData
	staticData.once.Do(planlexerLexerInit)
}

// NewPLANLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPLANLexer(input antlr.CharStream) *PLANLexer {
	PLANLexerInit()
	l := new(PLANLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PLANLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PLAN.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PLANLexer tokens.
const (
	PLANLexerT__0       = 1
	PLANLexerT__1       = 2
	PLANLexerT__2       = 3
	PLANLexerT__3       = 4
	PLANLexerT__4       = 5
	PLANLexerT__5       = 6
	PLANLexerT__6       = 7
	PLANLexerT__7       = 8
	PLANLexerT__8       = 9
	PLANLexerT__9       = 10
	PLANLexerT__10      = 11
	PLANLexerT__11      = 12
	PLANLexerT__12      = 13
	PLANLexerT__13      = 14
	PLANLexerT__14      = 15
	PLANLexerT__15      = 16
	PLANLexerT__16      = 17
	PLANLexerT__17      = 18
	PLANLexerT__18      = 19
	PLANLexerT__19      = 20
	PLANLexerT__20      = 21
	PLANLexerWS         = 22
	PLANLexerEq         = 23
	PLANLexerNeq        = 24
	PLANLexerOr         = 25
	PLANLexerAnd        = 26
	PLANLexerPow        = 27
	PLANLexerGtEq       = 28
	PLANLexerLtEq       = 29
	PLANLexerAssSum     = 30
	PLANLexerAssSub     = 31
	PLANLexerAssMul     = 32
	PLANLexerAssDiv     = 33
	PLANLexerAssMod     = 34
	PLANLexerAssPow     = 35
	PLANLexerGt         = 36
	PLANLexerLt         = 37
	PLANLexerMultiply   = 38
	PLANLexerDivision   = 39
	PLANLexerModulus    = 40
	PLANLexerAdd        = 41
	PLANLexerSubtract   = 42
	PLANLexerXor        = 43
	PLANLexerNot        = 44
	PLANLexerClosure    = 45
	PLANLexerBool       = 46
	PLANLexerNull       = 47
	PLANLexerIdentifier = 48
	PLANLexerInteger    = 49
	PLANLexerIntegerHex = 50
	PLANLexerFloat      = 51
	PLANLexerString_    = 52
	PLANLexerComment    = 53
)
