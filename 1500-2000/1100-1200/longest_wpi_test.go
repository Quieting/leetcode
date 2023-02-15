package leetcode

import (
	"testing"
)

func Test_longestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				hours: []int{6, 6, 6, 6},
			},
			want: 0,
		},

		{
			name: "示例2",
			args: args{
				hours: []int{9, 9, 6, 6},
			},
			want: 3,
		},
		{
			name: "示例3",
			args: args{
				hours: []int{9, 6, 6},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWPI(tt.args.hours); got != tt.want {
				t.Errorf("longestWPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
