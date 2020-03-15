/*
@date : 2020/03/14
@author : YaPi
@desc : 给定一个无序的整数数组，找到其中最长上升子序列的长度
*/
package main

import "fmt"

// 输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// [1,3,6,7,9,4,10,5,6]
	arr := []int{10, 9, 2, 5, 3, 4}
	fmt.Println(lengthOfLIS(arr))
}
