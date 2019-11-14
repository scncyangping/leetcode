/*
@date : 2019/11/14
@author : YaPi
@desc : 二叉树的最大深度
*/
package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	fmt.Println(maxDepth(nil))
}

func maxDepth(root *TreeNode) int {

	return depth(root, 0)
}

func depth(node *TreeNode, dep int) int {
	if node == nil || &node == nil {
		return 0
	}
	lDep, rDep := 0, 0
	if node.Left != nil {
		lDep = depth(node.Left, dep)
	}
	if node.Right != nil {
		rDep = depth(node.Right, dep)
	}
	if lDep > rDep {
		return lDep + 1
	} else {
		return rDep + 1
	}
}
