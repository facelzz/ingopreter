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
	'(', ')', '[', ']', '{', '}', ',', ';', '!', '~', '"',
}

// Tokenize TODO: line-by-line might no work because of multi-line tokens
func Tokenize(chars []rune) []Token {
	var tokens []Token // TODO: here's some random capacity

	var v rune
	for i := 0; i < len(chars); i++ {
		v = chars[i]
		slog.Debug("processing rune",
			"runeIndex", i,
			"runeValue", string(v))

		switch {
		// single char token
		case slices.Contains(singleCharTokens, v):
			tokens = append(tokens, Token{
				Value: string(chars[i : i+1]),
			})
			slog.Debug("single-char token captured")

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

			// capture token
			tokens = append(tokens, Token{
				Value: string(chars[i:j]),
			})
			i = j - 1
			slog.Debug("identifier or keyword captured")

		// unexpected
		default:
			// Unknown
			//panic("Unknown char: " + string(v))
		}
	}

	if len(tokens) == 0 {
		return nil
	}

	return tokens
}
