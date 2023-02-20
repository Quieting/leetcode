package leetcode

import (
	"testing"
)

func Test_largest1BorderedSquare(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			},
			want: 9,
		},
		{
			name: "示例2: 元素0",
			args: args{
				grid: [][]int{{0}},
			},
			want: 0,
		},
		{
			name: "示例3: 元素1",
			args: args{
				grid: [][]int{{1}},
			},
			want: 1,
		},
		{
			name: "示例3: 元素1",
			args: args{
				grid: [][]int{{0, 0, 0, 1}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largest1BorderedSquare(tt.args.grid); got != tt.want {
				t.Errorf("largest1BorderedSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
