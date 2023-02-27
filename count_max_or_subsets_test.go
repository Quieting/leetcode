package leetcode

import (
	"reflect"
	"testing"
)

func Test_countMaxOrSubsets(t *testing.T) {
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
				nums: []int{3, 1},
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				nums: []int{2, 2, 2},
			},
			want: 7,
		},
		{
			name: "示例3",
			args: args{
				nums: []int{3, 2, 1, 5},
			},
			want: 6,
		},
		{
			name: "示例4",
			args: args{
				nums: []int{87, 4, 41, 42, 100, 50, 93, 48, 56, 80, 2, 22, 77, 49, 37},
			},
			want: 29436,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = countMaxOrSubsets1(tt.args.nums)
			if got := countMaxOrSubsets(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitForOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				n: 0xf0,
			},
			want: []int{4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitForOne(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bitForOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
