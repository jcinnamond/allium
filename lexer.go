package main

import (
	"unicode/utf8"
)

type Scanner struct {
	input    string
	pos      int
	lastRune rune
}

func NewScanner(input string) *Scanner {
	return &Scanner{input: input, pos: 0}
}

func (s *Scanner) Scan() Token {
	s.skipWhitespace()
	switch s.nextChar() {
	case '(':
		s.advance()
		return Token{T_LPAREN, "("}
	default:
		identifier := s.scanIdentifier()
		return Token{T_IDENTIFIER, identifier}
	}
}

func (s *Scanner) backup() {
	s.pos -= utf8.RuneLen(s.lastRune)
}

func (s *Scanner) advance() {
	r, width := utf8.DecodeRuneInString(s.input[s.pos:])
	s.lastRune = r
	s.pos += width
}

func (s *Scanner) nextChar() rune {
	r, _ := utf8.DecodeRuneInString(s.input[s.pos:])
	return r
}

func (s *Scanner) skipWhitespace() {
	for s.isWhitespace() {
		s.advance()
	}
}

func (s *Scanner) scanIdentifier() string {
	start := s.pos
	for !(s.isEof() || s.isWhitespace()) {
		s.advance()
	}

	return s.input[start:s.pos]
}

func (s *Scanner) isWhitespace() bool {
	ch := s.nextChar()
	if ch == ' ' || ch == '\t' || ch == '\n' {
		return true
	}
	return false
}

func (s *Scanner) isEof() bool {
	return s.pos == len(s.input)
}
