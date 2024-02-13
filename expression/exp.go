package expression

import "glox/scanner"

// Visitor pattern
type Visitor interface {
	VisitAssignExpr(expr *AssignExpr) interface{}
	VisitBinaryExpr(expr *BinaryExpr) interface{}
	VisitCallExpr(expr *Call) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitLogicalExpr(expr *Logical) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
	VisitVariableExpr(expr *Variable) interface{}
	VisitCommaExpr(expr *Comma) interface{}
	VisitTernaryExpr(expr *Ternary) interface{}
}

type Expr interface {
	Accept(visitor Visitor) interface{}
}

// =, +=, -=, *=, /=
type AssignExpr struct {
	Name  scanner.Token
	Value Expr
}

func (a *AssignExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitAssignExpr(a)
}

// x + y
type BinaryExpr struct {
	Left     Expr
	Right    Expr
	Operator scanner.Token
}

func (b *BinaryExpr) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryExpr(b)
}

type Call struct {
	Callee    Expr
	Paren     scanner.Token
	Arguments []Expr
}

func (c *Call) Accept(visitor Visitor) interface{} {
	return visitor.VisitCallExpr(c)
}

// "(" expression ")"
type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitGroupingExpr(g)
}

// 123, "hello"
type Literal struct {
	Value interface{}
}

func (l *Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteralExpr(l)
}

// x && y
type Logical struct {
	Left     Expr
	Right    Expr
	Operator scanner.Token
}

func (l *Logical) Accept(visitor Visitor) interface{} {
	return visitor.VisitLogicalExpr(l)
}

// !x
type Unary struct {
	Right    Expr
	Operator scanner.Token
}

func (u *Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnaryExpr(u)
}

// x
type Variable struct {
	Name scanner.Token
}

func (v *Variable) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariableExpr(v)
}

// x, y
type Comma struct {
	Left  Expr
	Right Expr
}

func (c *Comma) Accept(visitor Visitor) interface{} {
	return visitor.VisitCommaExpr(c)
}

// x ? y : z
type Ternary struct {
	Expr1 Expr
	Expr2 Expr
	Expr3 Expr
}

func (t *Ternary) Accept(visitor Visitor) interface{} {
	return visitor.VisitTernaryExpr(t)
}
