package lexer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want []Token
	}{
		{
			name: "returns empty for empty input",
			args: args{
				chars: []rune{},
			},
			want: nil,
		},
		{
			name: "splits by space",
			args: args{
				chars: []rune("some word\nhere"),
			},
			want: []Token{
				{"some"},
				{"word"},
				{"here"},
			},
		},
		{
			name: "captures operators and punctuation",
			args: args{
				chars: []rune(`
					+    &     +=    &=     &&    ==    !=    (    )
					-    |     -=    |=     ||    <     <=    [    ]
					*    ^     *=    ^=     <-    >     >=    {    }
					/    <<    /=    <<=    ++    =     :=    ,    ;
					%    >>    %=    >>=    --    !     ...   .    :
						 &^          &^=          ~`),
			},
			want: []Token{
				{"+"},
				{"&"},
				{"+="},
				{"&="},
				{"&&"},
				{"=="},
				{"!="},
				{"("},
				{")"},
				{"-"},
				{"|"},
				{"-="},
				{"|="},
				{"||"},
				{"<"},
				{"<="},
				{"["},
				{"]"},
				{"*"},
				{"^"},
				{"*="},
				{"^="},
				{"<-"},
				{">"},
				{">="},
				{"{"},
				{"}"},
				{"/"},
				{"<<"},
				{"/="},
				{"<<="},
				{"++"},
				{"="},
				{":="},
				{","},
				{";"},
				{"%"},
				{">>"},
				{"%="},
				{">>="},
				{"--"},
				{"!"},
				{"..."},
				{"."},
				{":"},
				{"&^"},
				{"&^="},
				{"~"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tokenize(tt.args.chars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
