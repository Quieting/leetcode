package leetcode

import (
	"testing"
)

func Test_minOperationsMaxProfit(t *testing.T) {
	type args struct {
		customers    []int
		boardingCost int
		runningCost  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				customers:    []int{8, 3},
				boardingCost: 5,
				runningCost:  6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperationsMaxProfit(tt.args.customers, tt.args.boardingCost, tt.args.runningCost); got != tt.want {
				t.Errorf("minOperationsMaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
