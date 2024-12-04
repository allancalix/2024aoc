package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/allancalix/2024aoc/aoc"
)

type TokenTag int

const (
	TokenTagUnknown TokenTag = iota
	TokenTagKeyword
	TokenTagNumber
	TokenTagLParen
	TokenTagRParen
	TokenTagComma
)

type Token struct {
	tag   TokenTag
	start int
	end   int
}

type Lexer struct {
	input string
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) Lex() []Token {
	type state int
	const (
		stateInitial state = iota
		stateKeyword
		stateNumber
	)

	var tokens []Token
	var s state
	var result Token
	var idx int
	for idx < len(l.input) {
		c := l.input[idx]
		if s == stateInitial {
			switch {
			case c == 'm' || c == 'd':
				s = stateKeyword
				result.start = idx
				result.tag = TokenTagKeyword
			case c >= '0' && c <= '9':
				s = stateNumber
				result.start = idx
				result.tag = TokenTagNumber
			case c == '(':
				result.start = idx
				result.end = idx + 1
				result.tag = TokenTagLParen
				tokens = append(tokens, result)
				result = Token{}
			case c == ')':
				result.start = idx
				result.end = idx + 1
				result.tag = TokenTagRParen
				tokens = append(tokens, result)
				result = Token{}
			case c == ',':
				result.start = idx
				result.end = idx + 1
				result.tag = TokenTagComma
				tokens = append(tokens, result)
				result = Token{}
			default:
				result.tag = TokenTagUnknown
				result.start = idx
				result.end = idx + 1
				tokens = append(tokens, result)
			}

			idx++
			continue
		}

		if s == stateKeyword {
			if c == '\'' || (c >= 'a' && c <= 'z') {
				idx++
				continue
			}
		}

		if s == stateNumber {
			if c >= '0' && c <= '9' {
				idx++
				continue
			}
		}

		result.end = idx
		tokens = append(tokens, result)
		s = stateInitial
		result = Token{}
	}

	return tokens
}

type Expression interface {
	Kind() string
	Left() int
	Right() int
}

type DoExpression struct{}

func (d DoExpression) Kind() string {
	return "do"
}

func (d DoExpression) Left() int {
	return 0
}

func (d DoExpression) Right() int {
	return 0
}

type DoNotExpression struct{}

func (d DoNotExpression) Kind() string {
	return "do_not"
}

func (d DoNotExpression) Left() int {
	return 0
}

func (d DoNotExpression) Right() int {
	return 0
}

type MulExpression struct {
	left  int
	right int
}

func (m MulExpression) Kind() string {
	return "mul"
}

func (m MulExpression) Left() int {
	return m.left
}

func (m MulExpression) Right() int {
	return m.right
}

func parseExpressions(input string, tokens []Token) []Expression {
	type state int
	const (
		stateInitial state = iota
		stateMultiply
		stateMultiplyLeft
		stateMultiplyRightPre
		stateMultiplyRight
		stateMultiplyRightPost

		stateDo
		StateDoLeft

		stateDoNot
		StateDoNotLeft
	)

	var s state
	var exprs []Expression
	var result MulExpression
	for _, token := range tokens {
		if s == stateInitial {
			if token.tag == TokenTagKeyword {
				if input[token.start:token.end] == "mul" {
					s = stateMultiply
				}

				if input[token.start:token.end] == "do" {
					s = stateDo
				}

				if input[token.start:token.end] == "don't" {
					s = stateDoNot
				}
			}

			continue
		}

		if s == stateMultiply {
			if token.tag != TokenTagLParen {
				s = stateInitial
				continue
			}

			s = stateMultiplyLeft
			continue
		}

		if s == stateMultiplyLeft {
			if token.tag != TokenTagNumber {
				s = stateInitial
				continue
			}

			n := input[token.start:token.end]
			if len(n) >= 1 && len(n) <= 3 {
				value, err := strconv.Atoi(n)
				if err != nil {
					s = stateInitial
					continue
				}

				result.left = value
			}

			s = stateMultiplyRightPre
			continue
		}

		if s == stateMultiplyRightPre {
			if token.tag != TokenTagComma {
				s = stateInitial
				continue
			}

			s = stateMultiplyRight
			continue
		}

		if s == stateMultiplyRight {
			if token.tag != TokenTagNumber {
				s = stateInitial
				continue
			}

			n := input[token.start:token.end]
			if len(n) >= 1 && len(n) <= 3 {
				value, err := strconv.Atoi(n)
				if err != nil {
					s = stateInitial
					continue
				}

				result.right = value
			}

			s = stateMultiplyRightPost
			continue
		}

		if s == stateDo {
			if token.tag != TokenTagLParen {
				s = stateInitial
				continue
			}

			s = StateDoLeft
			continue
		}

		if s == StateDoLeft {
			if token.tag != TokenTagRParen {
				s = stateInitial
				continue
			}

			exprs = append(exprs, DoExpression{})
			s = stateInitial
			continue
		}

		if s == stateDoNot {
			if token.tag != TokenTagLParen {
				s = stateInitial
				continue
			}

			s = StateDoNotLeft
			continue
		}

		if s == StateDoNotLeft {
			if token.tag != TokenTagRParen {
				s = stateInitial
				continue
			}

			exprs = append(exprs, DoNotExpression{})
			s = stateInitial
			continue
		}

		if s == stateMultiplyRightPost {
			if token.tag != TokenTagRParen {
				s = stateInitial
				continue
			}
		}

		exprs = append(exprs, result)
		s = stateInitial
	}

	return exprs
}

func performConditionalProcessing() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	l := NewLexer(string(buf))
	expressions := parseExpressions(string(buf), l.Lex())

	enabled := true
	var sum int
	for _, expr := range expressions {
		switch expr.Kind() {
		case "mul":
			if !enabled {
				continue
			}

			sum += expr.Left() * expr.Right()
		case "do":
			enabled = true
		case "do_not":
			enabled = false
		default:
			log.Fatalf("Unknown expression type: %s", expr.Kind())
		}
	}

	fmt.Println("Total", sum)
}

func main() {
	aoc.Setup()

	if aoc.DoDayTwo() {
		performConditionalProcessing()

		return
	}

	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	l := NewLexer(string(buf))
	expressions := parseExpressions(string(buf), l.Lex())

	var sum int
	for _, expr := range expressions {
		sum += expr.Left() * expr.Right()
	}

	fmt.Println("Total", sum)
}
