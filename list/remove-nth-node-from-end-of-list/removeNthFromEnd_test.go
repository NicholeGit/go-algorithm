package remove_nth_node_from_end_of_list

import (
	"go-algorithm/list"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *list.Node
		n    int
	}
	tests := []struct {
		name string
		args args
		want *list.Node
	}{
		{
			name: "1",
			args: args{
				head: list.New([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: list.New([]int{1, 2, 3, 5}),
		},
		{
			name: "2",
			args: args{
				head: list.New([]int{1}),
				n:    1,
			},
			want: list.New([]int{}),
		},
		{
			name: "3",
			args: args{
				head: list.New([]int{1}),
				n:    2,
			},
			want: list.New([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got.ToSlice(), tt.want.ToSlice())
			}
		})
	}
}
