package leetcode

import "testing"

func Test_binaryCap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例: 不存在相邻二进制1",
			args: args{
				n: 0b10000,
			},
			want: 0,
		},
		{
			name: "示例: 存在相邻二进制1",
			args: args{
				n: 0b10010,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.n); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
