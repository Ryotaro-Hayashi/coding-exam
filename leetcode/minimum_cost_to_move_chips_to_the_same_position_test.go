package leetcode

import "testing"

func Test_minCostToMoveChips(t *testing.T) {
	type args struct {
		position []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				position: []int{1,2,3},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				position: []int{2,2,2,3,3},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				position: []int{1,1000000000},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				position: []int{2,3,3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostToMoveChips(tt.args.position); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
