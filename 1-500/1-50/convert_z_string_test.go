package leetcode

import (
	"testing"
)

func Test_convertZString(t *testing.T) {
	type args struct {
		s    string
		rows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "原样输出，一行",
			args: args{
				s:    "PAYPALISHIRING",
				rows: 1,
			},
			want: "PAYPALISHIRING",
		},
		{
			name: "原样输出，行数大于字符串",
			args: args{
				s:    "PAYPALISHIRING",
				rows: 20,
			},
			want: "PAYPALISHIRING",
		},
		{
			name: "示例1",
			args: args{
				s:    "PAYPALISHIRING",
				rows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "示例2",
			args: args{
				s:    "PAYPALISHIRING",
				rows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = convertZString1(tt.args.s, tt.args.rows)
			if got := convertZString(tt.args.s, tt.args.rows); got != tt.want {
				t.Errorf("convertZString() = %v, want %v", got, tt.want)
			}
		})
	}
}
