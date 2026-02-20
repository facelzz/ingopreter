package lexer

import (
	"log/slog"
	"slices"
	"unicode"
)

type Token struct {
	Value string
}

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

// Tokenize TODO: line-by-line might no work because of multi-line tokens
func Tokenize(chars []rune) []Token {
	var tokens []Token // TODO: here's some random capacity

	var v rune
	for i := 0; i < len(chars); i++ {
		v = chars[i]
		s := string(v)
		slog.Debug("processing rune",
			"runeIndex", i,
			"runeValue", s)

		charsToCapture := 0
		switch {
		// single char token
		case slices.Contains(singleCharTokens, v):
			charsToCapture = 1
			slog.Debug("single-char token captured")

		// operators `c`, `c=`
		case slices.Contains(simpleAssignStart, v):
			charsToCapture = 1
			if chars[i+1] == '=' {
				charsToCapture = 2
			}
			slog.Debug("operator/punctuation captured")

		// operators `c`, `cc`, `c=`
		case slices.Contains(simpleDoubleAssignStart, v):
			charsToCapture = 1
			if chars[i+1] == '=' || chars[i+1] == v {
				charsToCapture = 2
			}
			slog.Debug("operator/punctuation captured")

		// &...
		case v == '&':
			charsToCapture = 1
			if chars[i+1] == '^' {
				if chars[i+2] == '=' {
					charsToCapture = 3
				} else {
					charsToCapture = 2
				}
			} else if chars[i+1] == '=' || chars[i+1] == v {
				charsToCapture = 2
			}
			slog.Debug("operator/punctuation captured")

		// <...
		case v == '<':
			charsToCapture = 1
			if chars[i+1] == v {
				if chars[i+2] == '=' {
					charsToCapture = 3
				} else {
					charsToCapture = 2
				}
			} else if chars[i+1] == '=' || chars[i+1] == '-' {
				charsToCapture = 2
			}
			slog.Debug("operator/punctuation captured")

		// >...
		case v == '>':
			charsToCapture = 1
			if chars[i+1] == v {
				if chars[i+2] == '=' {
					charsToCapture = 3
				} else {
					charsToCapture = 2
				}
			} else if chars[i+1] == '=' {
				charsToCapture = 2
			}
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
			slog.Debug("operator/punctuation captured")

		// skip
		case unicode.IsSpace(v):
			slog.Debug("space skipped")

		// starts with letter - identifier ar keywork
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
			slog.Debug("identifier or keyword captured")

		// unexpected
		default:
			panic("Unknown char: " + string(v))
		}

		// capture token
		if charsToCapture > 0 {
			tokens = append(tokens, Token{
				Value: string(chars[i : i+charsToCapture]),
			})
			i += charsToCapture - 1
		}
	}

	if len(tokens) == 0 {
		return nil
	}

	return tokens
}
