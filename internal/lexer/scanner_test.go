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
