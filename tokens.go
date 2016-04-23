package main

type Token struct {
	t       int
	literal string
}

const (
	T_WHITESPACE = iota

	T_IDENTIFIER

	T_LPAREN
	T_RPAREN
)
