package leetcode

import "testing"

func Test_numIsLands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例: 一个岛屿",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
		{
			name: "示例: 多个岛屿",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			want: 3,
		},
		{
			name: "示例: 多个岛屿",
			args: args{
				grid: [][]byte{
					{'1', '1', '1'},
					{'0', '1', '0'},
					{'1', '1', '1'},
					{'0', '0', '0'},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIsLands(tt.args.grid); got != tt.want {
				t.Errorf("numIsLands() = %v, want %v", got, tt.want)
			}
		})
	}
}
