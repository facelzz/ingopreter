package lexer

import "testing"

func TestChspNewLine(t *testing.T) {
	for _, v := range []rune{'\n', rune(`
`[0]), rune("\n"[0])} {
		pass := chsp(v, newLine)
		if !pass {
			t.Errorf("%q is 'new line'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'\r', '\t', '\f'} {
		pass := chsp(v, newLine)
		if pass {
			t.Errorf("%q is not 'new line')", v)
		}
	}
}

func TestUnicodeChar(t *testing.T) {
	// random unicode chars
	for _, v := range []rune{'\t', 'F', '%', 'ќ', ';', 'ї', ' '} {
		pass := chsp(v, unicodeChar)
		if !pass {
			t.Errorf("%q is 'unicode char'", v)
		}
	}

	// no new line
	for _, v := range []rune{'\n', rune(`
`[0]), rune("\n"[0])} {
		pass := chsp(v, unicodeChar)
		if pass {
			t.Errorf("%q is not 'unicode char'", v)
		}
	}
}

func TestUnicodeLetter(t *testing.T) {
	for _, v := range []rune{'A', 'n', 'ќ', 'Ї'} {
		pass := chsp(v, unicodeLetter)
		if !pass {
			t.Errorf("%q is 'unicode letter'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'_', '\n', '\t', '\f', '.', ';', ' ', '5'} {
		pass := chsp(v, unicodeLetter)
		if pass {
			t.Errorf("%q is not 'unicode letter'", v)
		}
	}
}

func TestUnicodeDigit(t *testing.T) {
	for _, v := range []rune{'𑵔', '५', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		pass := chsp(v, unicodeDigit)
		if !pass {
			t.Errorf("%q is 'unicode digit'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'A', 'n', 'ќ', 'Ї', '\n', '\t', '\f', '.', ';', ' '} {
		pass := chsp(v, unicodeDigit)
		if pass {
			t.Errorf("%q is not 'unicode digit'", v)
		}
	}
}

func TestLetter(t *testing.T) {
	for _, v := range []rune{'_', 'A', 'n', 'ќ', 'Ї'} {
		pass := chsp(v, letter)
		if !pass {
			t.Errorf("%q is 'letter'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'\n', '\t', '\f', '.', ';', ' ', '5'} {
		pass := chsp(v, letter)
		if pass {
			t.Errorf("%q is not 'letter'", v)
		}
	}
}

func TestDecimalDigit(t *testing.T) {
	for _, v := range []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		pass := chsp(v, decimalDigit)
		if !pass {
			t.Errorf("%q is 'decimal digit'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'𑵔', '५', 'A', 'n', 'ќ', 'Ї', '\n', '\t', '\f', '.', ';', ' '} {
		pass := chsp(v, decimalDigit)
		if pass {
			t.Errorf("%q is not 'decimal digit'", v)
		}
	}
}

func TestBinaryDigit(t *testing.T) {
	for _, v := range []rune{'0', '1'} {
		pass := chsp(v, binaryDigit)
		if !pass {
			t.Errorf("%q is 'binary digit'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'3', '5', '𑵔', '५', 'A', 'n', 'ќ', 'Ї', '\n', '\t', '\f', '.', ';', ' '} {
		pass := chsp(v, binaryDigit)
		if pass {
			t.Errorf("%q is not 'binary digit'", v)
		}
	}
}

func TestOctalDigit(t *testing.T) {
	for _, v := range []rune{'0', '1', '2', '3', '4', '5', '6', '7'} {
		pass := chsp(v, octalDigit)
		if !pass {
			t.Errorf("%q is 'octal digit'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'8', '9', '𑵔', '५', 'A', 'n', 'ќ', 'Ї', '\n', '\t', '\f', '.', ';', ' '} {
		pass := chsp(v, octalDigit)
		if pass {
			t.Errorf("%q is not 'octal digit'", v)
		}
	}
}

func TestHexDigit(t *testing.T) {
	for _, v := range []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f'} {
		pass := chsp(v, hexDigit)
		if !pass {
			t.Errorf("%q is 'hex digit'", v)
		}
	}

	// negative cases
	for _, v := range []rune{'𑵔', '५', 'G', 'g', 'n', 'ќ', 'Ї', '\n', '\t', '\f', '.', ';', ' '} {
		pass := chsp(v, hexDigit)
		if pass {
			t.Errorf("%q is not 'hex digit'", v)
		}
	}
}
