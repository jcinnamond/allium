package main

import (
	"testing"
)

func TestScanSingleToken(t *testing.T) {
	cases := []struct {
		input    string
		expected Token
	}{
		{"onions", Token{T_IDENTIFIER, "onions"}},
		{"(", Token{T_LPAREN, "("}},
		{"🐙", Token{T_IDENTIFIER, "🐙"}},
	}

	for _, c := range cases {
		s := NewScanner(c.input)
		got := s.Scan()

		if got != c.expected {
			t.Errorf("Scan %v returned %v, expected %v",
				c.input, got, c.expected)
		}
	}
}

func TestIsEof(t *testing.T) {
	s := NewScanner("")
	if !s.isEof() {
		t.Errorf("Expected eof with empty string")
	}

	s = NewScanner("a")
	s.advance()
	if !s.isEof() {
		t.Errorf("Expected eof after advance")
	}
}

func TestIsWhitespace(t *testing.T) {
	s := NewScanner(" ")
	if !s.isWhitespace() {
		t.Errorf("Expected whitespace with single space")
	}

	s = NewScanner("a ")
	s.advance()
	if !s.isWhitespace() {
		t.Errorf("Expected whitespace after advance")
	}
}

func TestScanMultipleTokens(t *testing.T) {
	expected := []Token{
		Token{T_IDENTIFIER, "onions"},
		Token{T_LPAREN, "("},
	}

	s := NewScanner("onions (")
	for _, tok := range expected {
		got := s.Scan()

		if got != tok {
			t.Errorf("Scan returned %v, expected %v", got, tok)
		}
	}
}
