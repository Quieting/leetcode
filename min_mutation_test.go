package leetcode

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				start: "AAAAACCC",
				end:   "AACCCCCC",
				bank:  []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
