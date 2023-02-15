package leetcode

import (
	"testing"
)

func Test_balanceString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				s: "QWERERWEQWEWRQEW",
			},
		},
		{
			name: "示例2",
			args: args{
				s: "QQWE",
			},
		},
		{
			name: "示例2",
			args: args{
				s: "WQWRQQQW",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = balanceStringOfficial(tt.args.s)
			if got := balanceString(tt.args.s); got != tt.want {
				t.Errorf("balanceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
