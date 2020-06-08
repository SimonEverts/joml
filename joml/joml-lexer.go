// Copyright 2017 Simon Everts. All rights reserved.

package joml

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Token int

const (
	Illegal Token = iota
	EOF
	WS

	Ident

	Comma
	ObjectBegin
	ObjectEnd

	ArrayBegin
	ArrayEnd

	Import
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

type Lexer struct {
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: bufio.NewReader(r)}
}

func (s *Lexer) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (s *Lexer) unread() {
	_ = s.r.UnreadRune()
}

func (s *Lexer) Scan() (token Token, literal string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return EOF, ""
	case ',':
		return Comma, string(ch)
	case '{':
		return ObjectBegin, ""
	case '}':
		return ObjectEnd, ""
	case '[':
		return ArrayBegin, ""
	case ']':
		return ArrayEnd, ""
	}

	return Illegal, string(ch)
}

func (s *Lexer) scanWhitespace() (token Token, literal string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Lexer) scanIdent() (token Token, literal string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case "IMPORT":
		return Import, buf.String()
	}

	return Ident, buf.String()
}
