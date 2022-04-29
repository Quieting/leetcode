package leetcode

import (
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "示例：值相等",
			args: args{
				grid: [][]int{
					[]int{1, 1},
					[]int{1, 1},
				},
			},
			want: &Node{
				IsLeaf: true,
				Val:    true,
			},
		},
		{
			name: "示例：值不相等",
			args: args{
				grid: [][]int{
					[]int{1, 0},
					[]int{0, 1},
				},
			},
			want: &Node{
				IsLeaf: false,
				Val:    false,
				TopLeft: &Node{
					IsLeaf: true,
					Val:    true,
				},
				TopRight: &Node{
					IsLeaf: true,
					Val:    false,
				},
				BottomLeft: &Node{
					IsLeaf: true,
					Val:    false,
				},
				BottomRight: &Node{
					IsLeaf: true,
					Val:    true,
				},
			},
		},
		{
			name: "示例",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0, 0},
				},
			},
			want: &Node{
				IsLeaf: false,
				Val:    false,
				TopLeft: &Node{
					IsLeaf: true,
					Val:    true,
				},
				TopRight: &Node{
					TopLeft: &Node{
						IsLeaf: true,
						Val:    false,
					},
					TopRight: &Node{
						IsLeaf: true,
						Val:    false,
					},
					BottomLeft: &Node{
						IsLeaf: true,
						Val:    true,
					},
					BottomRight: &Node{
						IsLeaf: true,
						Val:    true,
					},
				},
				BottomLeft: &Node{
					IsLeaf: true,
					Val:    true,
				},
				BottomRight: &Node{
					IsLeaf: true,
					Val:    false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}
