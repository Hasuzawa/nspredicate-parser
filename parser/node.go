package parser

import (
	"fmt"

	"github.com/Hasuzawa/nspredicate-parser/parser/scanner/token"
)

type Node interface {
	Eval() string
}

type ExprNode struct {
	Left     Node
	Operator Operator
	Right    Node
}

func (n *ExprNode) Eval() string {
	return fmt.Sprintf("%s %s %s", n.Left.Eval(), n.Operator.String(), n.Right.Eval())
}

type StmtNode struct {
	Left  Node
	Token token.Token
	Right Node
}

func (n *StmtNode) Eval() string {
	return fmt.Sprintf("%s %s %s", n.Left.Eval(), n.Token.Literal(), n.Right.Eval())
}

var (
	_ Node = new(ExprNode)
	_ Node = new(StmtNode)
)
