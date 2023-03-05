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
		want *FourTreeNode
	}{
		{
			name: "示例：值相等",
			args: args{
				grid: [][]int{
					[]int{1, 1},
					[]int{1, 1},
				},
			},
			want: &FourTreeNode{
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
			want: &FourTreeNode{
				IsLeaf: false,
				Val:    false,
				TopLeft: &FourTreeNode{
					IsLeaf: true,
					Val:    true,
				},
				TopRight: &FourTreeNode{
					IsLeaf: true,
					Val:    false,
				},
				BottomLeft: &FourTreeNode{
					IsLeaf: true,
					Val:    false,
				},
				BottomRight: &FourTreeNode{
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
			want: &FourTreeNode{
				IsLeaf: false,
				Val:    false,
				TopLeft: &FourTreeNode{
					IsLeaf: true,
					Val:    true,
				},
				TopRight: &FourTreeNode{
					TopLeft: &FourTreeNode{
						IsLeaf: true,
						Val:    false,
					},
					TopRight: &FourTreeNode{
						IsLeaf: true,
						Val:    false,
					},
					BottomLeft: &FourTreeNode{
						IsLeaf: true,
						Val:    true,
					},
					BottomRight: &FourTreeNode{
						IsLeaf: true,
						Val:    true,
					},
				},
				BottomLeft: &FourTreeNode{
					IsLeaf: true,
					Val:    true,
				},
				BottomRight: &FourTreeNode{
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
