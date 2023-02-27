package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例: 相同的字符组成的字符串",
			args: args{
				s: "aaaaaa",
			},
			want: 1,
		},
		{
			name: "示例: 不相同的字符组成的字符串",
			args: args{
				s: "abcdefghikj",
			},
			want: 11,
		},
		{
			name: "示例: 有相同的字符组成的字符串",
			args: args{
				s: "abcdefaghikj",
			},
			want: 11,
		},
		{
			name: "示例: 有相同的字符组成的字符串",
			args: args{
				s: "abcdeefjh",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
