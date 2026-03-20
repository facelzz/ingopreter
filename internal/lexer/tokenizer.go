package lexer

import (
	"log/slog"
	"slices"
	"unicode"
)

type Token struct {
	Value string
	Type  TokenType
}

type TokenType int

const (
	Unknown TokenType = iota
	Identifier
	Keyword
	Operator
	Punctuation
	Literal
)

func (t TokenType) String() string {
	switch t {
	case Identifier:
		return "identifier"
	case Keyword:
		return "keyword"
	case Operator:
		return "operator"
	case Punctuation:
		return "punctuation"
	case Literal:
		return "literal"
	default:
		return "unknown"
	}
}

// TODO: mabe use binary masks for it?
var singleCharTokens = []rune{
	'(', ')', '[', ']', '{', '}', ',', ';', '~', '"',
}

// c, c=
var simpleAssignStart = []rune{
	'!', '^', '/', ':', '=', '*', '%',
}

// c, cc, c=
var simpleDoubleAssignStart = []rune{
	'+', '-', '|',
}

var keywords = []string{
	"break", "default", "func", "interface", "select",
	"case", "defer", "go", "map", "struct",
	"chan", "else", "goto", "package", "switch",
	"const", "fallthrough", "if", "range", "type",
	"continue", "for", "import", "return", "var",
}

// Tokenize TODO: line-by-line might no work because of multi-line tokens
func Tokenize(chars []rune) []Token {
	var tokens []Token // TODO: here's some random capacity

	var v rune
	for i := 0; i < len(chars); i++ {
		v = chars[i]
		hasNext := i+1 < len(chars)
		hasOverNext := i+2 < len(chars)
		s := string(v)
		slog.Debug("processing rune",
			"runeIndex", i,
			"runeValue", s)

		charsToCapture := 0
		tokenType := Unknown
		switch {
		// single char token
		case slices.Contains(singleCharTokens, v):
			charsToCapture = 1
			tokenType = Punctuation
			slog.Debug("single-char token captured")

		// operators `c`, `c=`
		case slices.Contains(simpleAssignStart, v):
			charsToCapture = 1
			if hasNext && chars[i+1] == '=' {
				charsToCapture = 2
			}
			tokenType = Operator
			if charsToCapture == 1 && v == ':' {
				tokenType = Punctuation
			}
			slog.Debug("operator/punctuation captured")

		// operators `c`, `cc`, `c=`
		case slices.Contains(simpleDoubleAssignStart, v):
			charsToCapture = 1
			if hasNext && (chars[i+1] == '=' || chars[i+1] == v) {
				charsToCapture = 2
			}
			tokenType = Operator
			slog.Debug("operator/punctuation captured")

		// &...
		case v == '&':
			charsToCapture = 1
			if hasNext {
				if chars[i+1] == '^' {
					if hasOverNext && chars[i+2] == '=' {
						charsToCapture = 3
					} else {
						charsToCapture = 2
					}
				} else if chars[i+1] == '=' || chars[i+1] == v {
					charsToCapture = 2
				}
			}
			tokenType = Operator
			slog.Debug("operator/punctuation captured")

		// <...
		case v == '<':
			charsToCapture = 1
			if hasNext {
				if chars[i+1] == v {
					if hasOverNext && chars[i+2] == '=' {
						charsToCapture = 3
					} else {
						charsToCapture = 2
					}
				} else if chars[i+1] == '=' || chars[i+1] == '-' {
					charsToCapture = 2
				}
			}
			tokenType = Operator
			slog.Debug("operator/punctuation captured")

		// >...
		case v == '>':
			charsToCapture = 1
			if hasNext {
				if chars[i+1] == v {
					if hasOverNext && chars[i+2] == '=' {
						charsToCapture = 3
					} else {
						charsToCapture = 2
					}
				} else if chars[i+1] == '=' {
					charsToCapture = 2
				}
			}
			tokenType = Operator
			slog.Debug("operator/punctuation captured")

		// dots
		case v == '.':
			// ... greedy capture
			j := i + 1
			for {
				if j >= len(chars) || chars[j] != '.' {
					break
				}
				j++
			}
			charsToCapture = j - i
			tokenType = Operator
			slog.Debug("operator/punctuation captured")

		// skip
		case unicode.IsSpace(v):
			slog.Debug("space skipped")

		// starts with letter - identifier or keyword
		// EBNF: identifier = letter { letter | unicode_digit } .
		case chsp(v, letter):
			// greedily capture token
			j := i + 1
			for {
				if j >= len(chars) || !chsp(chars[j], letter) && !chsp(chars[j], unicodeDigit) {
					break
				}
				j++
			}
			charsToCapture = j - i
			if slices.Contains(keywords, string(chars[i:i+charsToCapture])) {
				tokenType = Keyword
			} else {
				tokenType = Identifier
			}
			slog.Debug("identifier or keyword captured")

		// starts with digit - numeric literal
		case chsp(v, decimalDigit):
			// greedily capture token
			j := i + 1
			for {
				if j >= len(chars) ||
					!chsp(chars[j], decimalDigit) &&
						!chsp(chars[j], hexDigit) &&
						// TODO: process exponent part
						// TODO: process imaginary
						chars[j] != '_' && chars[j] != '.' &&
						chars[j] != 'b' && chars[j] != 'B' &&
						chars[j] != 'o' && chars[j] != 'O' &&
						chars[j] != 'x' && chars[j] != 'X' {
					break
				}
				j++
			}
			charsToCapture = j - i
			tokenType = Literal
			slog.Debug("identifier or keyword captured")

		// unexpected
		default:
			panic("Unknown char: " + string(v))
		}

		// capture token
		if charsToCapture > 0 {
			tokens = append(tokens, Token{
				Value: string(chars[i : i+charsToCapture]),
				Type:  tokenType,
			})
			i += charsToCapture - 1
		}
	}

	if len(tokens) == 0 {
		return nil
	}

	return tokens
}
