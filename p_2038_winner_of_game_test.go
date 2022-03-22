package leetcode

import "testing"

func Test_winnerOfGame(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				colors: "AA",
			},
			want: false,
		},
		{
			name: "示例2",
			args: args{
				colors: "AAABB",
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				colors: "ABBBBBBBAAA",
			},
			want: false,
		},
		{
			name: "示例4",
			args: args{
				colors: "AAAAAAABAAAABBBBBBBBBBB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
