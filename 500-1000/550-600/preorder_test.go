package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{root: &Node{
				Val: 0,
				Children: []*Node{
					{
						Val:      1,
						Children: nil,
					},
					{
						Val: 2,
						Children: []*Node{
							{
								Val:      3,
								Children: nil,
							},
							{
								Val:      4,
								Children: nil,
							},
						},
					},
				},
			}},
		},
		{
			name: "示例2",
			args: args{root: &Node{
				Val: 0,
				Children: []*Node{
					{
						Val: 1,
						Children: []*Node{
							{
								Val:      5,
								Children: nil,
							},
							{
								Val:      6,
								Children: nil,
							},
						},
					},
					{
						Val: 2,
						Children: []*Node{
							{
								Val:      3,
								Children: nil,
							},
							{
								Val:      4,
								Children: nil,
							},
						},
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = preorder1(tt.args.root)
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
