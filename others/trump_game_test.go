package others

import "testing"

func Test_winnerPairOfCards(t *testing.T) {
	type args struct {
		player1 []string
		player2 []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				player1: []string{"♣4","♥8","♥8","♠8","♣J"},
				player2: []string{"♣4","♥J","♥J","♠Q","♣3"},
			},
			want: "player1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerPairOfCards(tt.args.player1, tt.args.player2); got != tt.want {
				t.Errorf("winnerPaairOfCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
