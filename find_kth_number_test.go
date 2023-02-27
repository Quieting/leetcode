package leetcode

import (
	"testing"
)

func Test_findKthNumber1(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				n: 105,
				k: 10,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = findKthNumber1(tt.args.n, tt.args.k)
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSteps(t *testing.T) {
	type args struct {
		cur int
		n   int
	}
	tests := []struct {
		name      string
		args      args
		wantSteps int
	}{
		{
			name: "示例1",
			args: args{
				cur: 5,
				n:   20,
			},
			wantSteps: 1,
		},
		{
			name: "示例2",
			args: args{
				cur: 12,
				n:   20,
			},
			wantSteps: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSteps := getSteps(tt.args.cur, tt.args.n); gotSteps != tt.wantSteps {
				t.Errorf("getSteps() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
