package list

import (
	"testing"
)

func TestListNode_ToSlice(t *testing.T) {
	var l Node
	l.Gen([]int{1, 2, 3, 4, 5, 6, 10})
	t.Logf("%#v", l)
	t.Log(l.ToSlice())
}
