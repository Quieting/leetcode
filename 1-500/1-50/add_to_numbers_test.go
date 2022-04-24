package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例: 列表长度不一致",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 8,
					},
				},
				l2: &ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "示例: 相加结果最高位提高一位",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 8,
					},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = addTwoNumbers1(tt.args.l1, tt.args.l2)
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
