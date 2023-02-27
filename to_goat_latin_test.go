package leetcode

import "testing"

func Test_toGoatLatin(t *testing.T) {
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
				s: "I speak Goat Latin",
			},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		//
		{
			name: "示例1",
			args: args{
				s: "The quick brown fox jumped over the lazy dog",
			},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.s); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
