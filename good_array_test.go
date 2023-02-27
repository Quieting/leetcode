package leetcode

import (
	"testing"
)

func Test_isGoodArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1：好数组",
			args: args{
				nums: []int{3, 5, 7, 9},
			},
			want: true,
		},
		{
			name: "示例2：好数组,有重复数字",
			args: args{
				nums: []int{5, 5, 7, 9},
			},
			want: true,
		},
		{
			name: "示例3：不是好数组",
			args: args{
				nums: []int{2, 4, 6},
			},
			want: false,
		},
		{
			name: "示例4：好数组",
			args: args{
				nums: []int{3, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGoodArray(tt.args.nums); got != tt.want {
				t.Errorf("isGoodArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
