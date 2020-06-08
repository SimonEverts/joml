// Copyright 2017 Simon Everts. All rights reserved.

package joml

import "io"

type Parser struct {
	lexer *Lexer
	buf   struct {
		token   Token
		literal string
		n       int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{lexer: NewLexer(r)}
}

func (p *Parser) scan() (token Token, literal string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.token, p.buf.literal
	}

	token, literal = p.lexer.Scan()

	p.buf.token, p.buf.literal = token, literal

	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhitespace() (token Token, literal string) {
	token, literal = p.scan()
	if token == WS {
		token, literal = p.scanIgnoreWhitespace()
	}
	return
}

func (p *Parser) ParseRootObject() map[string]interface{} {
	item := make(map[string]interface{})

	if token, literal := p.scanIgnoreWhitespace(); token != Ident {
		panic("Expected identifier, found: " + literal)
	}

	return item
}
