# 977. 有序数组的平方

[977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/)  
[977. 有序数组的平方 解题思路](https://www.programmercarl.com/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.html)  

![移除元素-双指针法.gif](../../doc/977.有序数组的平方.gif)
- 因为最大的平方根值只会出现在左右两边，所以使用双指针左右比较放入结果集中
- 因为需要从大到小排序，所以从数组最后开始放入
