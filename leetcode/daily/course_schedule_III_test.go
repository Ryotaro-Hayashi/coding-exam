package daily

import "testing"

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				courses: [][]int{{100,200},{200,1300},{1000,1250},{2000,3200}},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				courses: [][]int{{1,2}},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				courses: [][]int{{3,2},{4,3}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scheduleCourse(tt.args.courses); got != tt.want {
				t.Errorf("scheduleCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
