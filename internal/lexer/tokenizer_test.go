package lexer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	type testCase struct {
		input []rune
		want  []Token
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
					want: []Token{
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
					want:  []Token{{"+", Operator}},
				},
				{
					input: []rune("&"),
					want:  []Token{{"&", Operator}},
				},
				{
					input: []rune("+="),
					want:  []Token{{"+=", Operator}},
				},
				{
					input: []rune("&="),
					want:  []Token{{"&=", Operator}},
				},
				{
					input: []rune("&&"),
					want:  []Token{{"&&", Operator}},
				},
				{
					input: []rune("=="),
					want:  []Token{{"==", Operator}},
				},
				{
					input: []rune("!="),
					want:  []Token{{"!=", Operator}},
				},
				{
					input: []rune("("),
					want:  []Token{{"(", Punctuation}},
				},
				{
					input: []rune(")"),
					want:  []Token{{")", Punctuation}},
				},
				{
					input: []rune("-"),
					want:  []Token{{"-", Operator}},
				},
				{
					input: []rune("|"),
					want:  []Token{{"|", Operator}},
				},
				{
					input: []rune("-="),
					want:  []Token{{"-=", Operator}},
				},
				{
					input: []rune("|="),
					want:  []Token{{"|=", Operator}},
				},
				{
					input: []rune("||"),
					want:  []Token{{"||", Operator}},
				},
				{
					input: []rune("<"),
					want:  []Token{{"<", Operator}},
				},
				{
					input: []rune("<="),
					want:  []Token{{"<=", Operator}},
				},
				{
					input: []rune("["),
					want:  []Token{{"[", Punctuation}},
				},
				{
					input: []rune("]"),
					want:  []Token{{"]", Punctuation}},
				},
				{
					input: []rune("*"),
					want:  []Token{{"*", Operator}},
				},
				{
					input: []rune("^"),
					want:  []Token{{"^", Operator}},
				},
				{
					input: []rune("*="),
					want:  []Token{{"*=", Operator}},
				},
				{
					input: []rune("^="),
					want:  []Token{{"^=", Operator}},
				},
				{
					input: []rune("<-"),
					want:  []Token{{"<-", Operator}},
				},
				{
					input: []rune(">"),
					want:  []Token{{">", Operator}},
				},
				{
					input: []rune(">="),
					want:  []Token{{">=", Operator}},
				},
				{
					input: []rune("{"),
					want:  []Token{{"{", Punctuation}},
				},
				{
					input: []rune("}"),
					want:  []Token{{"}", Punctuation}},
				},
				{
					input: []rune("/"),
					want:  []Token{{"/", Operator}},
				},
				{
					input: []rune("<<"),
					want:  []Token{{"<<", Operator}},
				},
				{
					input: []rune("/="),
					want:  []Token{{"/=", Operator}},
				},
				{
					input: []rune("<<="),
					want:  []Token{{"<<=", Operator}},
				},
				{
					input: []rune("++"),
					want:  []Token{{"++", Operator}},
				},
				{
					input: []rune("="),
					want:  []Token{{"=", Operator}},
				},
				{
					input: []rune(":="),
					want:  []Token{{":=", Operator}},
				},
				{
					input: []rune(","),
					want:  []Token{{",", Punctuation}},
				},
				{
					input: []rune(";"),
					want:  []Token{{";", Punctuation}},
				},
				{
					input: []rune("%"),
					want:  []Token{{"%", Operator}},
				},
				{
					input: []rune(">>"),
					want:  []Token{{">>", Operator}},
				},
				{
					input: []rune("%="),
					want:  []Token{{"%=", Operator}},
				},
				{
					input: []rune(">>="),
					want:  []Token{{">>=", Operator}},
				},
				{
					input: []rune("--"),
					want:  []Token{{"--", Operator}},
				},
				{
					input: []rune("!"),
					want:  []Token{{"!", Operator}},
				},
				{
					input: []rune("..."),
					want:  []Token{{"...", Operator}},
				},
				{
					input: []rune("."),
					want:  []Token{{".", Operator}},
				},
				{
					input: []rune(":"),
					want:  []Token{{":", Punctuation}}, // todo: fix
				},
				{
					input: []rune("&^"),
					want:  []Token{{"&^", Operator}},
				},
				{
					input: []rune("&^="),
					want:  []Token{{"&^=", Operator}},
				},
				{
					input: []rune("~"),
					want:  []Token{{"~", Punctuation}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tc := range tt.cases {
				if got := Tokenize(tc.input); !reflect.DeepEqual(got, tc.want) {
					t.Errorf("Tokenize() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
