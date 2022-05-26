package leetcode

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例: 空字符串",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "示例1",
			args: args{
				s: "()",
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				s: "(()()",
			},
			want: 4,
		},
		{
			name: "示例3",
			args: args{
				s: "))(()()))((((((())))))))",
			},
			want: 14,
		},
		{
			name: "示例4",
			args: args{
				s: "))(()())((((((())))))))",
			},
			want: 20,
		},
		{
			name: "示例5",
			args: args{
				s: "(((()())((((((())))))))",
			},
			want: 22,
		},
		{
			name: "示例6",
			args: args{
				s: "(()())((((((()))))))()",
			},
			want: 22,
		},
		{
			name: "示例6",
			args: args{
				s: "(())(",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
