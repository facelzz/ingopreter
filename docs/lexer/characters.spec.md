# Code characters

Go source code is not canonicalized case-sensitive Unicode text encoded in UTF-8,
but excluding NUL character (U+0000). UTF-8-encoded byte order mark (U+FEFF)
ignored at the start of the source text and disallowed in other places.

[Unicode 8.0.0 Standard Reference](https://www.unicode.org/versions/Unicode8.0.0/)

## Characters ([src](https://go.dev/ref/spec#Characters))

```ebnf
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

By Unicode 8.0.0 Standard Section 4.5 character categories definition:
- `Lu`, `Ll`, `Lt`, `Lm` and `Lo` as Letters
- `Nd` as digits.

## Letters and digits ([src](https://go.dev/ref/spec#Letters_and_digits))

```ebnf
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```
