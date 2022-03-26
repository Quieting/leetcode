package leetcode

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
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
			args: args{s: "ab-cd"},
			want: "dc-ba",
		},
		{
			name: "示例2",
			args: args{s: "a-bC-dEf-ghIj"},
			want: "j-Ih-gfE-dCba",
		},
		{
			name: "示例3",
			args: args{s: "Test1ng-Leet=code-Q!"},
			want: "Qedo1ct-eeLg=ntse-T!",
		},
		{
			name: "第一次执行失败",
			args: args{s: "-"},
			want: "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
