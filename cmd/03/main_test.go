package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
	// l := NewLexer("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	l := NewLexer("mul(2,434)")
	tokens := l.Lex()

	expected := []Token{
		{tag: TokenTagKeyword, start: 0, end: 3},
		{tag: TokenTagLParen, start: 3, end: 4},
		{tag: TokenTagNumber, start: 4, end: 5},
		{tag: TokenTagComma, start: 5, end: 6},
		{tag: TokenTagNumber, start: 6, end: 9},
		{tag: TokenTagRParen, start: 9, end: 10},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, token := range tokens {
		if token.tag != expected[i].tag {
			t.Fatalf("Expected token %d to be %d, got %d", i, expected[i].tag, token.tag)
		}

		if token.start != expected[i].start {
			t.Fatalf("Expected token %d to start at %d, got %d", i, expected[i].start, token.start)
		}

		if token.end != expected[i].end {
			t.Fatalf("Expected token %d to end at %d, got %d", i, expected[i].end, token.end)
		}
	}
}

func TestParseExpressions(t *testing.T) {
	src := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+\nmul(32,64]then(mul(11,8)mul(8,5))"
	l := NewLexer(src)
	tokens := l.Lex()
	exprs := parseExpressions(src, tokens)

	expected := []MulExpression{
		{left: 2, right: 4},
		{left: 5, right: 5},
		{left: 11, right: 8},
		{left: 8, right: 5},
	}

	if len(exprs) != len(expected) {
		t.Fatalf("Expected %d expressions, got %d", len(expected), len(exprs))
	}

	for i, expr := range exprs {
		if expr.left != expected[i].left {
			t.Fatalf("Expected expression %d to have left %d, got %d", i, expected[i].left, expr.left)
		}

		if expr.right != expected[i].right {
			t.Fatalf("Expected expression %d to have right %d, got %d", i, expected[i].right, expr.right)
		}
	}
}
