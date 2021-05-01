package daily

import (
	"testing"
)

func TestWordFilter_F(t *testing.T) {
	type fields struct {
		words []string
	}
	type args struct {
		prefix string
		suffix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "1",
			fields: fields{
				words: []string{"apple"},
			},
			args: args{
				prefix: "a",
				suffix: "e",
			},
			want: 0,
		},
		{
			name: "2",
			fields: fields{
				words: []string{"cabaabaaaa","ccbcababac","bacaabccba","bcbbcbacaa","abcaccbcaa","accabaccaa","cabcbbbcca","ababccabcb","caccbbcbab","bccbacbcba"},
			},
			args: args{
				prefix: "a",
				suffix: "aa",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &WordFilter{
				words: tt.fields.words,
			}
			if got := this.F(tt.args.prefix, tt.args.suffix); got != tt.want {
				t.Errorf("F() = %v, want %v", got, tt.want)
			}
		})
	}
}
