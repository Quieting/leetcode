package leetcode

import "testing"

func Test_countHighestScoreNodes(t *testing.T) {
	type args struct {
		parents []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				parents: []int{-1, 2, 0, 2, 0},
			},
			want: 3,
		},
		{
			name: "示例2",
			args: args{
				parents: []int{-1, 2, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHighestScoreNodes(tt.args.parents); got != tt.want {
				t.Errorf("countHighestScoreNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
