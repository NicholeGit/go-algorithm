package remove_nth_node_from_end_of_list

import "go-algorithm/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *list.Node, n int) *list.Node {
	dummyNode := &list.Node{
		Next: head,
	}

	//slow, fast := dummyNode, head
	//for i := 0; i < n; i++ { // fast 先往前走 n 步
	//	fast = fast.Next
	//}
	//for fast != nil {
	//	slow = slow.Next
	//	fast = fast.Next
	//}

	slow := dummyNode
	fast := head
	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n { // fast 先往前走 n 步
			slow = slow.Next
		}
		i++
	}

	if i > n {
		slow.Next = slow.Next.Next
	}

	return dummyNode.Next
}
