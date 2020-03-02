/*
@date : 2020/03/02
@author : YaPi
@desc : 反转链表
*/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for nil != cur {
		// 获取下一次遍历的节点
		tem := cur.Next
		// 当前节点反转后应为头节点
		cur.Next = pre
		// 设置得到到反转链表为前面设置的链表
		pre = cur
		// 设置下一次进行遍历的数据为前面暂存的数据
		cur = tem
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {

	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	head := reverse(cur, cur.Next)
	cur.Next = pre

	return head
}

func main() {
	r := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	rr := reverseList2(r)
	fmt.Println("test start ")
	for rr != nil {
		fmt.Println(rr.Val)
		rr = rr.Next
	}
}
