package leetcode

import (
	"testing"
)

func Test_minTaps(t *testing.T) {
	type args struct {
		n      int
		ranges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1：一个水龙头",
			args: args{
				n:      5,
				ranges: []int{0, 0, 8, 0, 0},
			},
			want: 1,
		},
		{
			name: "示例2：不能给整个花园浇水",
			args: args{
				n:      4,
				ranges: []int{0, 0, 1, 0, 0},
			},
			want: -1,
		},
		{
			name: "示例3：需要一个水龙头",
			args: args{
				n:      5,
				ranges: []int{3, 4, 1, 1, 0, 0},
			},
			want: 1,
		},
		{
			name: "示例3：需要一个水龙头",
			args: args{
				n:      4,
				ranges: []int{2, 0, 2, 0, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTaps(tt.args.n, tt.args.ranges); got != tt.want {
				t.Errorf("minTaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
