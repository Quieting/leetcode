package interview

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例: 一次替换两个字符串相同",
			args: args{
				first:  "abc",
				second: "aac",
			},
			want: true,
		},
		{
			name: "示例: 一次插入或者删除两个字符串相同",
			args: args{
				first:  "abc",
				second: "ac",
			},
			want: true,
		},
		{
			name: "示例: 一次插入或者删除两个字符串相同, 一个字符串为空串",
			args: args{
				first:  "",
				second: "a",
			},
			want: true,
		},
		{
			name: "示例: 两个字符串长度相差超过1",
			args: args{
				first:  "abc",
				second: "abccc",
			},
			want: false,
		},
		{
			name: "示例: 两个字符串长度相同, 需要替换的字符超过1",
			args: args{
				first:  "abcd",
				second: "abdc",
			},
			want: false,
		},
		{
			name: "示例: 两个字符串长度相差1, 一次操作不能成为相同字符串",
			args: args{
				first:  "abcd",
				second: "adc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
