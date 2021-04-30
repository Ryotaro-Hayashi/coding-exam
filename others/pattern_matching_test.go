package others

import (
	"testing"
)

func Test_hasSameType(t *testing.T) {
	type args struct {
		user1 string
		user2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				user1: "aabb",
				user2: "yyza",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				user1: "aappl",
				user2: "bbtte",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				user1: "aappl",
				user2: "bbttb",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				user1: "aappl",
				user2: "bktte",
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				user1: "aaapppl",
				user2: "bbbttke",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				user1: "abcd",
				user2: "tso",
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				user1: "abcd",
				user2: "jklm",
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				user1: "aaabbccdddaa",
				user2: "jjjddkkpppjj",
			},
			want: true,
		},
		{
			name: "9",
			args: args{
				user1: "jjjddkkpppjj",
				user2: "jjjddkkpppjd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSameType(tt.args.user1, tt.args.user2); got != tt.want {
				t.Errorf("hasSameType() = %v, want %v", got, tt.want)
			}
		})
	}
}
