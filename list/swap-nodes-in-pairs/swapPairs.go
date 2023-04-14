package swap_nodes_in_pairs

import "go-algorithm/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *list.Node) *list.Node {
	dummyNode := &list.Node{
		Next: head,
	}
	pre := dummyNode
	for head != nil && head.Next != nil {
		next := head.Next.Next // 临时保存
		pre.Next = head.Next   // 第一步
		head.Next.Next = head  // 第二步
		head.Next = next       // 第三步

		pre = head  // 遍历的前一个
		head = next // 当前遍历

	}
	return dummyNode.Next
}
