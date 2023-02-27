package leetcode

import "testing"

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				n: 2,
			},
		},
		{
			name: "示例2",
			args: args{
				n: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = countNumbersWithUniqueDigits1(tt.args.n)
			if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
