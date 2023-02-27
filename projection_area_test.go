package leetcode

import "testing"

func Test_projectionArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例: 多行多列",
			args: args{
				grid: [][]int{{1, 4}, {1, 1}},
			},
		},
		{
			name: "示例: 只有一行一列",
			args: args{
				grid: [][]int{{2}},
			},
		},
		{
			name: "示例: 有空的正方体叠放在单元格",
			args: args{
				grid: [][]int{{2, 0}, {3, 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = projectionArea1(tt.args.grid)
			if got := projectionArea(tt.args.grid); got != tt.want {
				t.Errorf("projectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
