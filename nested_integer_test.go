package leetcode

import (
	"reflect"
	"testing"
)

func Test_deserialize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *NestedInteger
	}{
		{
			name: "示例1",
			args: args{
				s: "[123,[456,[789]]]",
			},
			want: &NestedInteger{
				vals: []int{123, 456, 789},
			},
		},
		{
			name: "示例2",
			args: args{
				s: "12",
			},
			want: &NestedInteger{
				vals: []int{12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
