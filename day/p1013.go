/*
@date : 2020/03/11
@author : YaPi
@desc : 将数组分成和相等的三个部分
*/
package main

import "fmt"

// 给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
//
//形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
//
//链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum

// 输出：[0,2,1,-6,6,-7,9,1,2,0,1]
//输出：true
//解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
//
//链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum

// 方法1 暴力破解 -- 超出时间限制
func canThreePartsEqualSum2(A []int) bool {
	if A == nil || len(A) < 3 {
		return false
	}
	allNum := 0
	for i := 0; i < len(A); i++ {
		allNum += A[i]
	}
	baseI := 0
	// 选定标定点 i  i > 0 i < j-1  j <len(A)
	for i := 0; i < len(A)-2; i++ {
		baseI += A[i]
		baseJ := 0
		for j := i + 1; j < len(A)-1; j++ {
			baseJ += A[j]
			if baseI == baseJ && allNum-baseJ*2 == baseI {
				return true
			}
		}
	}

	return false
}

func canThreePartsEqualSum(A []int) bool {
	if A == nil {
		return false
	}
	allNum := 0
	for i := 0; i < len(A); i++ {
		allNum += A[i]
	}
	if allNum%3 != 0 {
		return false
	}
	target := allNum / 3
	cur := 0
	i := 0
	for i < len(A) {
		cur += A[i]
		if cur == target {
			break
		}
		i++
	}
	if cur != target {
		return false
	}
	j := i + 1
	for j < len(A)-1 {
		cur += A[j]
		if cur == target*2 {
			return true
		}
		j++
	}
	return false
}

func main() {
	arr := []int{3, 3}
	fmt.Println(canThreePartsEqualSum(arr))

}
