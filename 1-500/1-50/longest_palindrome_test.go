package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "示例2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "示例3",
			args: args{
				s: "abc",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = longestPalindrome1(tt.args.s)
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
