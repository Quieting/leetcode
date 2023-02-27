package leetcode

import (
	"reflect"
	"testing"
)

func Test_numsOfPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1:一个元素",
			args: args{nums: []int{1}},
			want: []int{0, 1},
		},

		{
			name: "示例1:一对数对元素",
			args: args{nums: []int{1, 1, 1, 3, 4, 5}},
			want: []int{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numsOfPairs(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
