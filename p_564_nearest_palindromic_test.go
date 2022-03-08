package leetcode

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "原数不是回文数",
			args: args{s: "12"},
			want: "11",
		},
		{
			name: "一位数",
			args: args{s: "1"},
			want: "0",
		},
		{
			name: "原数是回文数",
			args: args{s: "11"},
			want: "9",
		},
		{
			name: "最接近的回文数由多个9组成",
			args: args{s: "101"},
			want: "99",
		},
		{
			name: "原数是回文数,返回的回文数比原数大",
			args: args{s: "202"},
			want: "212",
		},
		{
			name: "原数是回文数,返回的回文数比原数小",
			args: args{s: "292"},
			want: "282",
		},
		{
			name: "原数是回文数由9组成",
			args: args{s: "99"},
			want: "101",
		},
		{
			name: "1283",
			args: args{s: "1283"},
			want: "1331",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.s); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
