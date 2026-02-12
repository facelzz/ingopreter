package lexer

import "unicode"

type CharSpec int

const (
	newLine CharSpec = iota
	unicodeChar
	unicodeLetter
	unicodeDigit
	letter
	decimalDigit
	binaryDigit
	octalDigit
	hexDigit
)

// check char spec
func chsp(char rune, spec CharSpec) bool {
	switch spec {
	case newLine:
		return char == '\u000A'
	case unicodeChar:
		return !chsp(char, newLine)
	case unicodeLetter:
		return unicode.IsLetter(char)
	case unicodeDigit:
		return unicode.IsDigit(char)
	case letter:
		return chsp(char, unicodeLetter) || char == '_'
	case decimalDigit:
		return char >= '0' && char <= '9'
	case binaryDigit:
		return char == '0' || char == '1'
	case octalDigit:
		return char >= '0' && char <= '7'
	case hexDigit:
		return (char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')
	}

	return false
}
