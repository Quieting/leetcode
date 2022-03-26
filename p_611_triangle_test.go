package leetcode

import "testing"

func Test_triangleNumber(t *testing.T) {
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
				nums: []int{2, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "示例2",
			args: args{
				nums: []int{4, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "示例3",
			args: args{
				nums: []int{2, 2, 3, 4, 10, 14, 17, 38, 26},
			},
			want: 9,
		},
		{
			name: "示例4",
			args: args{
				nums: []int{7, 0, 0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = triangleNumber1(tt.args.nums)
			if got := triangleNumber(tt.args.nums); got != tt.want {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
