package lexer

import (
	"log"
	"unicode"
)

type Token struct {
	Value string
}

func Tokenize(chars []rune, verbose bool) []Token {
	var tokens []Token // TODO: here's some random capacity

	capturingToken := false
	tokenStart := 0
	for i, v := range chars {
		if verbose {
			log.Print("Char[", i, "]: ", v)
		}
		if !unicode.IsSpace(v) && !capturingToken {
			if verbose {
				log.Print("Word start!")
			}
			capturingToken = true
			continue
		}
		if unicode.IsSpace(v) && capturingToken {
			if verbose {
				log.Print("Word end!")
			}
			// Capture token
			tokens = append(tokens, Token{
				Value: string(chars[tokenStart:i]),
			})

			tokenStart = i + 1
			capturingToken = false
		}
	}

	if len(tokens) == 0 {
		return nil
	}

	return tokens
}
