package remove_linked_list_elements

import (
	"go-algorithm/list"
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *list.Node
		val  int
	}
	tests := []struct {
		name string
		args args
		want *list.Node
	}{
		{
			name: "1",
			args: args{
				head: list.New([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: list.New([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "2",
			args: args{
				head: list.New([]int{1, 1, 1, 1, 1}),
				val:  1,
			},
			want: list.New([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got.ToSlice(), tt.want.ToSlice())
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementsByDummy(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got.ToSlice(), tt.want.ToSlice())
			}
		})
	}
}
