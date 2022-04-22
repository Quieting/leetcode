package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    0,
			},
			want: 40,
		},
		{
			name: "示例2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxRotateFunction(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{4, 3, 2, 6},
			},
		},
		{
			name: "示例2",
			args: args{
				nums: []int{4, 3, 2, 6, 0, -3},
			},
		},
		{
			name: "示例3",
			args: args{
				nums: []int{4, 3, 5, 10, 2099, 2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = maxRotateFunction1(tt.args.nums)
			if got := maxRotateFunction(tt.args.nums); got != tt.want {
				t.Errorf("maxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
