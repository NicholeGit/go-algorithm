package remove_linked_list_elements

import "go-algorithm/list"

func removeElements(head *list.Node, val int) *list.Node {
	first := head
	for first != nil && first.Val == val {
		first = first.Next
	}
	if first == nil {
		return nil
	}
	pre := first
	cur := first.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return first
}

// 使用虚拟头结点
func removeElementsByDummy(head *list.Node, val int) *list.Node {
	dummyNode := &list.Node{
		Next: head,
	}
	pre := dummyNode
	cur := dummyNode.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummyNode.Next
}
