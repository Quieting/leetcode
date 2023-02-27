package leetcode

import (
	"testing"
)

func Test_circularPermutation(t *testing.T) {
	type args struct {
		n     int
		start int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "示例1",
			args: args{
				n:     7,
				start: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := circularPermutation(tt.args.n, tt.args.start)
			for i := 1; i < len(got); i++ {
				n := got[i] ^ got[i-1]
				if n&(n-1) != 0 {
					t.Errorf("当前结果不符合要求: got: %v, i-1:%b, i: %b", got, got[i-1], got[i])
				}
			}
		})
	}
}
