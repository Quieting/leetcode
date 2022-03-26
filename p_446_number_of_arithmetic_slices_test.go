package leetcode

import "testing"

func Test_numberOfArithmeticSlices2(t *testing.T) {
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
				nums: []int{2, 4, 6, 8, 10},
			},
			want: 7,
		},
		{
			name: "示例2",
			args: args{
				nums: []int{7, 7, 7, 7, 7},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices2(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices2() = %v, want %v", got, tt.want)
			}
		})
	}
}
