package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例：多个元素",
			args: args{
				nums: []int{3, 1, 2, 4},
			},
			want: []int{4, 2, 1, 3},
		},
		{
			name: "示例：1个元素",
			args: args{
				nums: []int{3},
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
