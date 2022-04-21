package leetcode

import "testing"

func Test_lengthLongestPath(t *testing.T) {
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
				s: "file1.txt\nfile2.txt\nlongfile.txt",
			},
			want: 12,
		},
		{
			name: "示例2",
			args: args{
				s: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			},
			want: 20,
		},
		{
			name: "示例3",
			args: args{
				s: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			},
			want: 32,
		},
		//
		{
			name: "示例4",
			args: args{
				s: "a\n\tb\n\t\tc.txt",
			},
			want: 9,
		},
		{
			name: "示例5",
			args: args{
				s: "dir\n        file.txt",
			},
			want: 16,
		},
		{
			name: "示例6",
			args: args{
				s: "a\n\tb.txt\na2\n\tb2.txt",
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = lengthLongestPath1(tt.args.s)
			if got := lengthLongestPath(tt.args.s); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
