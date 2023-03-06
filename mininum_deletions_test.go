package leetcode

import (
	"testing"
)

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				s: "aababbab",
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				s: "bbaaaaabb",
			},
			want: 2,
		},
		{
			name: "示例3",
			args: args{
				s: "a",
			},
			want: 0,
		},
		{
			name: "示例4",
			args: args{
				s: "aaabababbaabbabbababbaaabaabbbbbbabbbabbabbaaabababbabbbabbbbabbabbabbbabaaaababababaababbaabaaababbbabbabbabababbbbbbaaaaaaabbbabababbbabbbaaaababbbbabaaabaaabababaaababaabbaabbababbaaabbabaaabbaaabbbaabaabbabbbbaabaaabaaaabaabaabaaaaaaabbaaaababbaaaaababbbbaaabbaaaaabbbbbababaaaabbaababbabbaababbbaababaaaaaabbbbaaaaabbaabbbbbaabbbabbaabbbaaaaabbbbaaababbaabaaabbababbaaaabbbaabbabaaaabbaaaaabbaaabbbababaaabaaabbababbbbbabababbabbbaabaaabbaabababbbabaaaaaa",
			},
			want: 220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
