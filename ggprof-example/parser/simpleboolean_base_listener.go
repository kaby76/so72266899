// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleBooleanListener is a complete listener for a parse tree produced by SimpleBooleanParser.
type BaseSimpleBooleanListener struct{}

var _ SimpleBooleanListener = &BaseSimpleBooleanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleBooleanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleBooleanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleBooleanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleBooleanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseSimpleBooleanListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseSimpleBooleanListener) ExitParse(ctx *ParseContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleBooleanListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleBooleanListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseSimpleBooleanListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseSimpleBooleanListener) ExitComparator(ctx *ComparatorContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseSimpleBooleanListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseSimpleBooleanListener) ExitBinary(ctx *BinaryContext) {}

// EnterBoolean_ is called when production boolean_ is entered.
func (s *BaseSimpleBooleanListener) EnterBoolean_(ctx *Boolean_Context) {}

// ExitBoolean_ is called when production boolean_ is exited.
func (s *BaseSimpleBooleanListener) ExitBoolean_(ctx *Boolean_Context) {}
