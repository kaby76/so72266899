// Code generated from parser/SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleBooleanListener is a complete listener for a parse tree produced by SimpleBooleanParser.
type SimpleBooleanListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterBoolean_ is called when entering the boolean_ production.
	EnterBoolean_(c *Boolean_Context)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitBoolean_ is called when exiting the boolean_ production.
	ExitBoolean_(c *Boolean_Context)
}
