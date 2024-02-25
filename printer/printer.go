package printer

import "glox/expr"

type Printer struct{}

// Printer needs to implement the Visitor interface
var _ expr.Visitor = &Printer{}

// VisitAssignExpr implements expression.Visitor.
func (*Printer) VisitAssignExpr(expr *expr.Assign) interface{} {
	panic("unimplemented")
}

// VisitCallExpr implements expr.Visitor.
func (*Printer) VisitCallExpr(expr *expr.Call) interface{} {
	panic("unimplemented")
}

// VisitCommaExpr implements expr.Visitor.
func (*Printer) VisitCommaExpr(expr *expr.Comma) interface{} {
	panic("unimplemented")
}

// VisitLogicalExpr implements expr.Visitor.
func (*Printer) VisitLogicalExpr(expr *expr.Logical) interface{} {
	panic("unimplemented")
}

// VisitTernaryExpr implements expr.Visitor.
func (*Printer) VisitTernaryExpr(expr *expr.Ternary) interface{} {
	panic("unimplemented")
}

// VisitVariableExpr implements expr.Visitor.
func (*Printer) VisitVariableExpr(expr *expr.Variable) interface{} {
	panic("unimplemented")
}

func (p *Printer) VisitBinaryExpr(expr *expr.Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *Printer) VisitGroupingExpr(expr *expr.Grouping) interface{} {
	return p.parenthesize("group", expr.Expression)
}

func (p *Printer) VisitLiteralExpr(expr *expr.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return expr.Value.(string)
}

func (p *Printer) VisitUnaryExpr(expr *expr.Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *Printer) parenthesize(name string, exprs ...expr.Expr) string {
	builder := ""
	builder += "(" + name
	for _, expr := range exprs {
		builder += " "
		builder += expr.Accept(p).(string) // This assertion is not safe
	}
	builder += ")"
	return builder
}

func (p *Printer) Print(expr expr.Expr) interface{} {
	return expr.Accept(p)
}
