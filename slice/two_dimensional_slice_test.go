package slice

import (
	"reflect"
	"testing"
)

func Test_addColumn(t *testing.T) {
	type args struct {
		matrix [][]int
		index  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				index: 0,
			},
			want: [][]int{
				{0, 1, 2, 3},
				{0, 4, 5, 6},
			},
		},
		{
			name: "示例2",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				index: 2,
			},
			want: [][]int{
				{1, 2, 0, 3},
				{4, 5, 0, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddColumn(tt.args.matrix, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addRow(t *testing.T) {
	type args struct {
		matrix [][]int
		index  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				index: 0,
			},
			want: [][]int{
				{0, 0, 0},
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{
			name: "示例2",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				index: 1,
			},
			want: [][]int{
				{1, 2, 3},
				{0, 0, 0},
				{4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRow(tt.args.matrix, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := PrefixSum(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
