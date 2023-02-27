package leetcode

import "testing"

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "数组长度为1",
			args: args{
				nums: []int{1},
			},
			want: "1",
		},
		{
			name: "数组长度为2",
			args: args{
				nums: []int{1, 2},
			},
			want: "1/2",
		},

		{
			name: "数组长度为3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: "1/(2/3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}
