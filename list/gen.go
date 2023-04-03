package list

// Node Definition for singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

func (l *Node) Gen(nums []int) *Node {
	size := len(nums)
	if size > 1 {
		l.Val = nums[0]
	}
	cur := l
	for i := 1; i < size; i++ {
		tmp := &Node{
			Val:  nums[i],
			Next: nil,
		}
		cur.Next = tmp
		cur = tmp
	}
	return l
}

func (l *Node) ToSlice() []int {
	var re []int
	cur := l
	for cur != nil {
		re = append(re, cur.Val)
		cur = cur.Next
	}
	return re
}

func (l *Node) String() string {
	return ""
}

func (l *Node) GetVal() int {
	return l.Val
}

func New(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	var n Node
	return n.Gen(nums)
}
