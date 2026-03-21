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
					input: []rune("some 'v' rune"),
					want: []Lexeme{
						{"some", Identifier},
						{"'v'", Literal},
						{"rune", Identifier},
					},
				},
				{
					input: []rune("12 13.3 12_000 13.3e+12 23.4E+31 0x2p-1 0XfP+1"),
					want: []Lexeme{
						{"12", Literal},
						{"13.3", Literal},
						{"12_000", Literal},
						{"13.3e+12", Literal},
						{"23.4E+31", Literal},
						{"0x2p-1", Literal},
						{"0XfP+1", Literal},
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
