package leetcode

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
				nums: []int{1, 2, 4, 5},
			},
			want: 0,
		}, {
			name: "示例2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "示例3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 7, 8, 9},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
