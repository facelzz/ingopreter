package lexer

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	type testCase struct {
		input []rune
		want  []Lexeme
	}
	tests := []struct {
		name  string
		cases []testCase
	}{
		{
			name: "returns empty for empty input",
			cases: []testCase{
				{
					input: []rune{},
					want:  nil,
				},
			},
		},
		{
			name: "splits by space",
			cases: []testCase{
				{
					input: []rune("some word\nhere"),
					want: []Lexeme{
						{"some", Identifier},
						{"word", Identifier},
						{"here", Identifier},
					},
				},
			},
		},
		{
			name: "captures operators and punctuation",
			cases: []testCase{
				{
					input: []rune("+"),
					want:  []Lexeme{{"+", Operator}},
				},
				{
					input: []rune("&"),
					want:  []Lexeme{{"&", Operator}},
				},
				{
					input: []rune("+="),
					want:  []Lexeme{{"+=", Operator}},
				},
				{
					input: []rune("&="),
					want:  []Lexeme{{"&=", Operator}},
				},
				{
					input: []rune("&&"),
					want:  []Lexeme{{"&&", Operator}},
				},
				{
					input: []rune("=="),
					want:  []Lexeme{{"==", Operator}},
				},
				{
					input: []rune("!="),
					want:  []Lexeme{{"!=", Operator}},
				},
				{
					input: []rune("("),
					want:  []Lexeme{{"(", Punctuation}},
				},
				{
					input: []rune(")"),
					want:  []Lexeme{{")", Punctuation}},
				},
				{
					input: []rune("-"),
					want:  []Lexeme{{"-", Operator}},
				},
				{
					input: []rune("|"),
					want:  []Lexeme{{"|", Operator}},
				},
				{
					input: []rune("-="),
					want:  []Lexeme{{"-=", Operator}},
				},
				{
					input: []rune("|="),
					want:  []Lexeme{{"|=", Operator}},
				},
				{
					input: []rune("||"),
					want:  []Lexeme{{"||", Operator}},
				},
				{
					input: []rune("<"),
					want:  []Lexeme{{"<", Operator}},
				},
				{
					input: []rune("<="),
					want:  []Lexeme{{"<=", Operator}},
				},
				{
					input: []rune("["),
					want:  []Lexeme{{"[", Punctuation}},
				},
				{
					input: []rune("]"),
					want:  []Lexeme{{"]", Punctuation}},
				},
				{
					input: []rune("*"),
					want:  []Lexeme{{"*", Operator}},
				},
				{
					input: []rune("^"),
					want:  []Lexeme{{"^", Operator}},
				},
				{
					input: []rune("*="),
					want:  []Lexeme{{"*=", Operator}},
				},
				{
					input: []rune("^="),
					want:  []Lexeme{{"^=", Operator}},
				},
				{
					input: []rune("<-"),
					want:  []Lexeme{{"<-", Operator}},
				},
				{
					input: []rune(">"),
					want:  []Lexeme{{">", Operator}},
				},
				{
					input: []rune(">="),
					want:  []Lexeme{{">=", Operator}},
				},
				{
					input: []rune("{"),
					want:  []Lexeme{{"{", Punctuation}},
				},
				{
					input: []rune("}"),
					want:  []Lexeme{{"}", Punctuation}},
				},
				{
					input: []rune("/"),
					want:  []Lexeme{{"/", Operator}},
				},
				{
					input: []rune("<<"),
					want:  []Lexeme{{"<<", Operator}},
				},
				{
					input: []rune("/="),
					want:  []Lexeme{{"/=", Operator}},
				},
				{
					input: []rune("<<="),
					want:  []Lexeme{{"<<=", Operator}},
				},
				{
					input: []rune("++"),
					want:  []Lexeme{{"++", Operator}},
				},
				{
					input: []rune("="),
					want:  []Lexeme{{"=", Operator}},
				},
				{
					input: []rune(":="),
					want:  []Lexeme{{":=", Operator}},
				},
				{
					input: []rune(","),
					want:  []Lexeme{{",", Punctuation}},
				},
				{
					input: []rune(";"),
					want:  []Lexeme{{";", Punctuation}},
				},
				{
					input: []rune("%"),
					want:  []Lexeme{{"%", Operator}},
				},
				{
					input: []rune(">>"),
					want:  []Lexeme{{">>", Operator}},
				},
				{
					input: []rune("%="),
					want:  []Lexeme{{"%=", Operator}},
				},
				{
					input: []rune(">>="),
					want:  []Lexeme{{">>=", Operator}},
				},
				{
					input: []rune("--"),
					want:  []Lexeme{{"--", Operator}},
				},
				{
					input: []rune("!"),
					want:  []Lexeme{{"!", Operator}},
				},
				{
					input: []rune("..."),
					want:  []Lexeme{{"...", Operator}},
				},
				{
					input: []rune("."),
					want:  []Lexeme{{".", Operator}},
				},
				{
					input: []rune(":"),
					want:  []Lexeme{{":", Punctuation}},
				},
				{
					input: []rune("&^"),
					want:  []Lexeme{{"&^", Operator}},
				},
				{
					input: []rune("&^="),
					want:  []Lexeme{{"&^=", Operator}},
				},
				{
					input: []rune("~"),
					want:  []Lexeme{{"~", Punctuation}},
				},
			},
		},
		{
			name: "Literals",
			cases: []testCase{
				{
					input: []rune(`some "string with 12 _ +$'" string`),
					want: []Lexeme{
						{"some", Identifier},
						{"\"string with 12 _ +$'\"", Literal},
						{"string", Identifier},
					},
				},
				{
					input: []rune(`'v' '\377' 'xFF' '\uF913' '\UF913F832' '\r' '\v'`),
					want: []Lexeme{
						{"'v'", Literal},
						{"'\\377'", Literal},
						{"'xFF'", Literal},
						{`'\uF913'`, Literal},
						{`'\UF913F832'`, Literal},
						{`'\r'`, Literal},
						{`'\v'`, Literal},
					},
				},
				{
					input: []rune("12 13.3 12_000 13.3e+12 23.4E+31 0x2p-1 0XfP+1 0o71 0O_71 0b01 0B_01 12i 13.2i 0x2i 0X3i 0o7i 0O7i 0b10i 0B11i"), // TODO: add octal and binary
					want: []Lexeme{
						{"12", Literal},
						{"13.3", Literal},
						{"12_000", Literal},
						{"13.3e+12", Literal},
						{"23.4E+31", Literal},
						{"0x2p-1", Literal},
						{"0XfP+1", Literal},
						{"0o71", Literal},
						{"0O_71", Literal},
						{"0b01", Literal},
						{"0B_01", Literal},
						{"12i", Literal},
						{"13.2i", Literal},
						{"0x2i", Literal},
						{"0X3i", Literal},
						{"0o7i", Literal},
						{"0O7i", Literal},
						{"0b10i", Literal},
						{"0B11i", Literal},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tc := range tt.cases {
				if got := Scan(tc.input); !reflect.DeepEqual(got, tc.want) {
					t.Errorf("Scan() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
