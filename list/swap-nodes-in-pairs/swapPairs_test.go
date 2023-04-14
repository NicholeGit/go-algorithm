package swap_nodes_in_pairs

import (
	"go-algorithm/list"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *list.Node
	}
	tests := []struct {
		name string
		args args
		want *list.Node
	}{
		{
			name: "1",
			args: args{
				head: list.New([]int{1, 2, 3, 4}),
			},
			want: list.New([]int{2, 1, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got.ToSlice(), tt.want.ToSlice())
			}
		})
	}
}
