package leetcode

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "有相同的元素，但元素差值不等于k",
			args: args{
				nums: []int{1, 0, 2, 3, 4, 5, 6, 1},
				k:    1,
			},
			want: false,
		},
		{
			name: "没有相同的元素",
			args: args{
				nums: []int{1, 0, 2, 3, 4, 5, 6},
				k:    1,
			},
			want: false,
		},
		{
			name: "有相同的元素，元素差值等于k",
			args: args{
				nums: []int{1, 0, 2, 3, 4, 5, 6, 1},
				k:    7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
