# 203. 移除链表元素

[203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/)  
[203. 移除链表元素 解题思路](https://www.programmercarl.com/0203.%E7%A7%BB%E9%99%A4%E9%93%BE%E8%A1%A8%E5%85%83%E7%B4%A0.html)  

- 移除头结点和移除其他节点的操作是不一样的，因为链表的其他节点都是通过前一个节点来移除当前节点，而头结点没有前一个节点。
- 其实可以设置一个虚拟头结点，这样原链表的所有节点就都可以按照统一的方式进行移除了。  
![203. 移除链表元素.png](../../doc/977.有序数组的平方.gif)