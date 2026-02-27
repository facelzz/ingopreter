# Tokens

Tokens form vocabulary of the Go. They are:
- identifiers
- keywords
- operators & punctuation
- literals

## Identifiers ([src](https://go.dev/ref/spec#Identifiers))

Identifiers are names for language entities - custom and predeclared.

From Go Spec:
```
Types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap clear close complex copy delete imag len
	make max min new panic print println real recover
```

EBNF:
```ebnf
identifier = letter { letter | unicode_digit } .
```

## Keywords ([src](https://go.dev/ref/spec#Keywords))

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## Operators/Punctuations ([src](https://go.dev/ref/spec#Operators_and_punctuation))

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

## Literals

### Integer ([src](https://go.dev/ref/spec#Integer_literals))

```ebnf
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

### Floats ([src](https://go.dev/ref/spec#Floating-point_literals))

```ebnf
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
```

### Imaginary ([src](https://go.dev/ref/spec#Imaginary_literals))

```ebnf
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

### Rune ([src](https://go.dev/ref/spec#Rune_literals))

```ebnf
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```

### String ([src](https://go.dev/ref/spec#String_literals))

```ebnf
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

# Implementation

Tokens can be grouped like:
1. Single char tokens
This tokens can be easilly identified because they are always (almost, see 
comments) terminating by itself e.g.: `(`, `)`, `[`, `]`, `{`,
`}`, `,`, `;`, `!`, `~`, `"`, `'`. These are some of opernators&punctuation.
2. Continuous words
Sequence of non-terminating and non-space chars. These are identifiers, keywords,
and literals in simple non-validated representation.
3. Complex tokens
Sequence of chars that cannot be identifier as "continuous word" because of
some sort of exceptions, e.g. strings with spaces or multi-char operators like
`<-`, `*=`. This group also includes signle-char tokens that can't be identified
clearly as "signle char token" by the first token, e.g. `*` can be multiplication
or a first char of `*=` operation.

> [!NOTE]
> Should rune/string quotes be separate chars?
> They are always in context of their context, but later we need to validate it.

## Simple implementation

Implement "Single char tokens" and "Continuous words" as is, make lookups
(forward/backward) for "Complext tokens".

This may work but implementation guaranteed to be messy and hard to read
because of infinite "ifs".

## Very simple implementation

Use prioritised regexps :/

This is not even interesing...

## Very hard implementation

Build a by-char spec graph. In reality this is just a "prioritied regexps"
but with graph and handmade regexps :D

Nodes of this directed graph are char spec and edges are links to the next and
previous char specs.

Do not know if it is possible though