package leetcode

import (
	"reflect"
	"testing"
)

func Test_prefixSum(t *testing.T) {
	type args struct {
		s [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例1",
			args: args{
				s: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
				},
			},
			want: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 3, 6, 10},
				{0, 6, 14, 24, 36},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixSum(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSumSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				matrix: [][]int{
					{1, 0, 1},
					{0, -2, 3},
				},
				k: 2,
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				matrix: [][]int{
					{2, 2, -1},
				},
				k: 3,
			},
			want: 3,
		},
		{
			name: "示例28",
			args: args{
				matrix: [][]int{
					{2, 2, -1},
				},
				k: 0,
			},
			want: -1,
		},

		{
			name: "示例31",
			args: args{
				matrix: [][]int{
					{5, -4, -3, 4},
					{-3, -4, 4, 5},
					{5, 1, 5, -4},
				},
				k: -2,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumSubmatrix(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("maxSumSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
