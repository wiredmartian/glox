package printer

import "glox/expression"

type Printer struct{}

// VisitAssignExpr implements expression.Visitor.
func (*Printer) VisitAssignExpr(expr *expression.Assign) interface{} {
	panic("unimplemented")
}

// VisitCallExpr implements expression.Visitor.
func (*Printer) VisitCallExpr(expr *expression.Call) interface{} {
	panic("unimplemented")
}

// VisitCommaExpr implements expression.Visitor.
func (*Printer) VisitCommaExpr(expr *expression.Comma) interface{} {
	panic("unimplemented")
}

// VisitLogicalExpr implements expression.Visitor.
func (*Printer) VisitLogicalExpr(expr *expression.Logical) interface{} {
	panic("unimplemented")
}

// VisitTernaryExpr implements expression.Visitor.
func (*Printer) VisitTernaryExpr(expr *expression.Ternary) interface{} {
	panic("unimplemented")
}

// VisitVariableExpr implements expression.Visitor.
func (*Printer) VisitVariableExpr(expr *expression.Variable) interface{} {
	panic("unimplemented")
}

// Printer needs to implement the Visitor interface
var _ expression.Visitor = &Printer{}

func (p *Printer) VisitBinaryExpr(expr *expression.Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *Printer) VisitGroupingExpr(expr *expression.Grouping) interface{} {
	return p.parenthesize("group", expr.Expression)
}

func (p *Printer) VisitLiteralExpr(expr *expression.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return expr.Value.(string)
}

func (p *Printer) VisitUnaryExpr(expr *expression.Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *Printer) parenthesize(name string, exprs ...expression.Expr) string {
	builder := ""
	builder += "(" + name
	for _, expr := range exprs {
		builder += " "
		builder += expr.Accept(p).(string)
	}
	builder += ")"
	return builder
}
