package parser

import (
	"fmt"
	"go-json-parser/ast"
	"go-json-parser/lexer"
	"go-json-parser/token"
)

type Parser struct {
	Lexer *lexer.Lexer
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{Lexer: l}
}

func (p *Parser) Parse() ast.Json {
	tok := p.Lexer.NewToken()
	switch tok.Type {
	case token.STRING:
		return &ast.String{string(tok.Lit)}
	case token.INTEGER:
		return &ast.Integer{string(tok.Lit)}
	case token.LBRACE:
		return parseObject(p)
	case token.LBRACKET:
		return parseArray(p)
	case token.EOF:
		return nil
	}
	return nil
}

func parseArray(p *Parser) ast.Json {
	array := []ast.Json{}
	tok := p.Lexer.PeakToken()

	if tok.Type == token.RBRACKET {
		return &ast.Array{Values: array}
	} else {
		array = append(array, p.Parse())
		tok = p.Lexer.NewToken()
		if tok.Type == token.RBRACKET {
			return &ast.Array{Values: array}
		}
	}

	for {
		array = append(array, p.Parse())
		tok = p.Lexer.NewToken()
		if tok.Type == token.RBRACKET {
			break
		}
		if tok.Type != token.COMMA {
			panic(fmt.Sprintf("was expecting ',' got %s in array parse", string(tok.Lit)))

		}
	}
	return &ast.Array{Values: array}
}

func parseObject(p *Parser) ast.Json {
	object := map[string]ast.Json{}
	tok := p.Lexer.NewToken()

	if tok.Type == token.RBRACE {
		return &ast.Object{Properties: object}
	} else {
		key := string(tok.Lit)
		p.Lexer.NewToken()
		object[key] = p.Parse()
		tok = p.Lexer.NewToken()
		if tok.Type == token.RBRACE {
			return &ast.Object{Properties: object}
		}
	}

	for {
		key := string(p.Lexer.NewToken().Lit)
		tok = p.Lexer.NewToken()
		if tok.Type != token.COLON {
			panic(fmt.Sprintf("was expecting ':' got %s", string(tok.Lit)))
		}
		object[key] = p.Parse()
		tok = p.Lexer.NewToken()

		if tok.Type == token.RBRACE {
			break
		}

		if tok.Type != token.COMMA {
			panic(fmt.Sprintf("was expecting ',' got %s", string(tok.Lit)))
		}
	}
	return &ast.Object{Properties: object}
}
