/*
@date : 2019/10/29
@author : YaPi
@desc : 两数相加
*/

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 近位
	flag := 0
	resNode := &ListNode{}
	// 都不为空继续向下执行
	for l1 != nil || l2 != nil || flag != 0 {
		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		num := v1 + v2 + flag
		// 进位数
		if num >= 10 {
			flag = 1
		} else {
			flag = 0
		}
		gum := num % 10
		// 创建子节点
		add(gum, resNode)
	}

	return resNode.Next
}

func add(v int, node *ListNode) {
	if node.Next == nil {
		node.Next = &ListNode{Val: v}
		return
	}
	add(v, node.Next)
}

func main() {
	//l1 := &ListNode{Val:2,Next:&ListNode{Val:4,Next:&ListNode{Val:3}}}
	//l2 := &ListNode{Val:5,Next:&ListNode{Val:6}}
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}

	res := addTwoNumbers(l1, l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
