package leetcode

import (
	"reflect"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		m []int
		n []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "示例: 只有一个中位数，且在较短的数组中",
			args: args{
				m: []int{10, 20, 30},
				n: []int{11, 12},
			},
			want: 12,
		},
		{
			name: "示例: 只有一个中位数，且在较长的数组中",
			args: args{
				m: []int{10, 20, 30},
				n: []int{8, 9},
			},
			want: 10,
		},
		{
			name: "示例: 只有一个中位数，且为较短的数组的第一位数",
			args: args{
				m: []int{10, 12, 13, 30},
				n: []int{15, 16, 17},
			},
			want: 15,
		},
		{
			name: "示例: 只有一个中位数，且为较短的数组的最后一位数",
			args: args{
				m: []int{10, 18, 20, 30},
				n: []int{15, 16, 17},
			},
			want: 17,
		},

		{
			name: "示例: 有两个中位数，分别在两个数组中",
			args: args{
				m: []int{10, 20, 23, 31, 32},
				n: []int{8, 21, 25},
			},
			want: 22,
		},
		{
			name: "示例: 有两个中位数，且在较长的数组中",
			args: args{
				m: []int{10, 20, 30, 35, 36, 37},
				n: []int{8, 9},
			},
			want: 25,
		},
		{
			name: "示例: 有两个中位数，且在较短的数组中",
			args: args{
				m: []int{10, 12, 13, 30, 35, 36},
				n: []int{15, 16, 17, 18},
			},
			want: 16.5,
		},
		{
			name: "示例: 有两个中位数，其中一个为较短的数组的最后一位数",
			args: args{
				m: []int{10, 18, 20, 30},
				n: []int{15, 16},
			},
			want: 17,
		},
		{
			name: "示例: 有两个中位数，其中一个为较短的数组的第一位数",
			args: args{
				m: []int{10, 16, 20, 30},
				n: []int{17, 23},
			},
			want: 18.5,
		},
		{
			name: "示例: 有两个中位数，其中一个数组为空",
			args: args{
				m: []int{10, 16},
				n: []int{},
			},
			want: 13,
		},
		{
			name: "示例: 有两个中位数，其中一个数组比另一个数组第一个值都要小",
			args: args{
				m: []int{10, 16},
				n: []int{17, 18, 19, 20, 21, 22},
			},
			want: 18.5,
		},
		{
			name: "示例: 有两个中位数，两个数组元素为同一个值",
			args: args{
				m: []int{10, 10},
				n: []int{10, 10},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSmaller(t *testing.T) {
	type args struct {
		m []int
		n []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例: m 数组为空",
			args: args{
				m: []int{},
				n: []int{1, 2, 3, 4},
			},
			want: []int{1, 2},
		},
		{
			name: "示例: n 数组为空",
			args: args{
				m: []int{4, 7, 8},
				n: []int{},
			},
			want: []int{4, 7},
		},
		{
			name: "示例: 最小两个数在 m 数组中",
			args: args{
				m: []int{4, 7, 8},
				n: []int{9},
			},
			want: []int{4, 7},
		},

		{
			name: "示例: 最小两个数在 n 数组中",
			args: args{
				m: []int{4, 7, 8},
				n: []int{2, 3},
			},
			want: []int{2, 3},
		},
		{
			name: "示例: 最小两个数一个数组一个",
			args: args{
				m: []int{4, 7, 8},
				n: []int{2, 5},
			},
			want: []int{2, 4},
		},
		{
			name: "示例: 一个数组一个数",
			args: args{
				m: []int{4},
				n: []int{2},
			},
			want: []int{2, 4},
		},
		{
			name: "示例: 只有一个数",
			args: args{
				m: []int{},
				n: []int{2},
			},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSmaller(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
