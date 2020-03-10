/*
@date : 2020/03/10
@author : YaPi
@desc : 二叉树的直径
*/
package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定一棵二叉树，你需要计算它的直径长度。
// 一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
// 			1
//         / \
//        2   3
//       / \
//      4   5

// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLength = 0.0

// 最大长度及为两棵子树高度
func diameterOfBinaryTree(root *TreeNode) int {
	// LeetCode 所有用例用的全局存储，不设置全局变量maxLength为上一次的值
	maxLength = 0.0
	high(root)
	return int(maxLength)
}

func high(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	highL := high(node.Left)
	highR := high(node.Right)

	maxLength = math.Max(maxLength, highR+highL)

	return math.Max(highR, highL) + 1
}

func main() {
	tNode := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 9,
					},
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(diameterOfBinaryTree(tNode))
}
