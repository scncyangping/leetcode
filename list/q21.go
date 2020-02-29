/*
@date : 2020/02/29
@author : YaPi
@desc : 合并两个有序链表
*/
package main

import "fmt"

// 将两个有序链表合并为一个新的有序链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 未说明是两个长度相等的链表
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	var (
//		l *ListNode
//		arr = make([]int,0)
//	)
//	if l1 == nil && l2 != nil {
//		return l2
//	}
//	if l2 == nil && l1 != nil {
//		return l1
//	}
//
//	p1 := l1
//	p2 := l2
//	for p1 != nil || p2 != nil {
//		if p1 != nil && p2 != nil {
//			// 比较两个指针对应的值，取最大的新建
//			if p1.Val <= p2.Val {
//				arr = append(arr, p1.Val)
//				p1 = p1.Next
//			}else {
//				arr = append(arr, p2.Val)
//				p2 = p2.Next
//			}
//		}else if p1 != nil && p2 == nil {
//			arr = append(arr, p1.Val)
//			p1 = p1.Next
//
//		}else {
//			arr = append(arr, p2.Val)
//			p2 = p2.Next
//		}
//	}
//
//	for i :=len(arr)-1 ;i >= 0 ;i-- {
//		ne := &ListNode{Val:arr[i],Next:l}
//		if i == len(arr) -1 {
//			ne = &ListNode{Val:arr[i]}
//		}
//		l = ne
//	}
//	return l
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	l3 := mergeTwoLists(l1, l2)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
