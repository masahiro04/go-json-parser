package ast

type Json interface {
	TokenLiteral()
}

type Array struct {
	Values []Json
}

type Object struct {
	Properties map[string]Json
}

type String struct {
	Value string
}

type Integer struct {
	Value string
}
