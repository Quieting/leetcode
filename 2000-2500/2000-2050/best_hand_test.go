package leetcode

import (
	"testing"
)

func Test_bestHand(t *testing.T) {
	type args struct {
		ranks []int
		suits []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{
				ranks: []int{10, 10, 12, 2, 1},
				suits: []byte{'a', 'b', 'c', 'd'},
			},
			want: "Pair",
		},
		{
			name: "示例1",
			args: args{
				ranks: []int{4, 4, 2, 4, 4},
				suits: []byte{'a', 'b', 'c', 'd'},
			},
			want: "Pair",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestHand(tt.args.ranks, tt.args.suits); got != tt.want {
				t.Errorf("bestHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
