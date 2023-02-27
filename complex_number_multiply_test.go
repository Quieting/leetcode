package leetcode

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "自然数",
			args: args{
				num1: "3+0i",
				num2: "4+0i",
			},
			want: "12+0i",
		},
		{
			name: "正数虚部",
			args: args{
				num1: "3+2i",
				num2: "4+1i",
			},
			want: "10+11i",
		},
		{
			name: "负数虚部",
			args: args{
				num1: "3+-1i",
				num2: "4+-2i",
			},
			want: "10+-10i",
		},
		{
			name: "一正一负虚部",
			args: args{
				num1: "3+-1i",
				num2: "4+2i",
			},
			want: "14+2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
