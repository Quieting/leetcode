package leetcode

import (
	"reflect"
	"testing"
)

func Test_findSolution(t *testing.T) {
	type args struct {
		customFunction func(int, int) int
		z              int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例1",
			args: args{
				customFunction: add,
				z:              5,
			},
			want: [][]int{{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSolution(tt.args.customFunction, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
