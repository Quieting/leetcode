package leetcode

import "testing"

func TestRecentCounter_Ping(t *testing.T) {
	type fields struct {
		pings []int
	}
	type args struct {
		t int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "示例: 请求查询全范围",
			fields: fields{
				pings: []int{1, 2, 3, 4, 5, 6, 7},
			},
			args: args{
				t: 8,
			},
			want: 8,
		},
		{
			name: "示例: 请求查询部分范围",
			fields: fields{
				pings: []int{1, 3002, 3008},
			},
			args: args{
				t: 3009,
			},
			want: 3,
		},
		{
			name: "示例: 请求查询部分范围",
			fields: fields{
				pings: []int{1, 9, 3002, 3008},
			},
			args: args{
				t: 3009,
			},
			want: 4,
		},
		{
			name: "示例: 请求查询部分范围",
			fields: fields{
				pings: []int{},
			},
			args: args{
				t: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RecentCounter{
				pings: tt.fields.pings,
			}
			if got := r.Ping(tt.args.t); got != tt.want {
				t.Errorf("Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}
