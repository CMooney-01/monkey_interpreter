package ast

import (
	"monkey_interpreter/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node // embedded anonymous field - Statement type has access to all Node fields (in this case, the .TokenLiteral() method)
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() // return root node for given program
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

// Methods have a reciever (ls *LetStatement)
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}


