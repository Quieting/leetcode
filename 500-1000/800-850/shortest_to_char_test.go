package leetcode

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				s: "loveleetcode",
				c: 'e',
			},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "示例2",
			args: args{
				s: "aaab",
				c: 'b',
			},
			want: []int{3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
