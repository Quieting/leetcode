package leetcode

import (
	"reflect"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	type args struct {
		operates []string
		args     [][]int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "示例1",
			args: args{
				operates: []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
				args:     [][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}},
			},
			want: []interface{}{nil, true, false, true, 2, true, false, 2},
		},
		{
			name: "示例2",
			args: args{
				operates: []string{"RandomizedSet", "remove", "remove", "insert", "getRandom", "remove", "insert"},
				args:     [][]int{{}, {0}, {0}, {0}, {}, {0}, {0}},
			},
			want: []interface{}{nil, false, false, true, 0, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var set RandomizedSet
			var got []interface{}
			for i, opt := range tt.args.operates {
				switch opt {
				case "RandomizedSet":
					set = RandomSetConstructor()
					got = append(got, nil)
				case "insert":
					got = append(got, set.Insert(tt.args.args[i][0]))
				case "remove":
					got = append(got, set.Remove(tt.args.args[i][0]))
				case "getRandom":
					got = append(got, set.GetRandom())
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
