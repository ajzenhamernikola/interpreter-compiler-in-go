package parser

import (
	"ast"
	"lexer"
	"token"
)

type Parser struct {
	l *lexer.Lexer

	currToken	token.Token
	peekToken	token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens,
	// so currToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
