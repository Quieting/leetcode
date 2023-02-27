package leetcode

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{
				dominoes: "RR.L",
			},
		},
		{
			name: "示例2",
			args: args{
				dominoes: ".L.R...LR..L..",
			},
		},
		{
			name: "示例3",
			args: args{
				dominoes: "LLRRR..RL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = pushDominoes1(tt.args.dominoes)
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
