### leetcode 206:反转链表

1. 题目
    
    反转一个单链表。
    
    示例:
    
    输入: 1->2->3->4->5->NULL
    
    输出: 5->4->3->2->1->NULL
    
    进阶:
    你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
    
2. 解题思路

    迭代法：
   
    递归法：


3. 代码实现

      迭代法：
       
       /**
        * Definition for singly-linked list.
        * type ListNode struct {
        *     Val int
        *     Next *ListNode
        * }
        */
       func reverseList(head *ListNode) *ListNode {
          	var pre *ListNode = nil
          	cur:=this.head
          	for nil!=cur{
          		tempNextNode:=cur.next
          		cur.next = pre
          		pre = cur
          		cur =  tempNextNode
          	}
          	return pre
       }         
                 
                 
                 
