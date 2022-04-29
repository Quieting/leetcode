package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例: 无效整数-首字符非数字",
			args: args{
				s: "re32",
			},
			want: 0,
		},
		{
			name: "示例: 无效整数-只有'+'、'-'",
			args: args{
				s: "+",
			},
			want: 0,
		},
		{
			name: "示例: 有效整数-负数",
			args: args{
				s: "-32",
			},
			want: -32,
		},
		{
			name: "示例: 有效整数-正数",
			args: args{
				s: "+3890",
			},
			want: 3890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
