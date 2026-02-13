package lexer

import (
	"log/slog"
	"unicode"
)

type Token struct {
	Value string
}

// Tokenize TODO: line-by-line might no work because of multi-line tokens
func Tokenize(chars []rune) []Token {
	var tokens []Token // TODO: here's some random capacity

	capturingToken := false
	tokenStart := 0
	for i, v := range chars {
		slog.Debug("processing rune",
			"runeIndex", i,
			"runeValue", string(v))

		// capture token start
		if !unicode.IsSpace(v) && !capturingToken {
			slog.Debug("Word start!")

			capturingToken = true
			continue
		}

		// capture terminated token
		if unicode.IsSpace(v) {
			if capturingToken {
				slog.Debug("Word end!")

				// Capture token
				tokens = append(tokens, Token{
					Value: string(chars[tokenStart:i]),
				})
			}

			tokenStart = i + 1
			capturingToken = false
			continue
		}
	}
	// capture non-closed token
	if capturingToken {
		slog.Debug("Word end!")

		// Capture token
		tokens = append(tokens, Token{
			Value: string(chars[tokenStart:]),
		})

		capturingToken = false
	}

	if len(tokens) == 0 {
		return nil
	}

	return tokens
}
