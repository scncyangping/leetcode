/*
@date : 2020/02/28
@author : YaPi
@desc :
*/
package main

import "fmt"

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
// 不要使用额外的数组空间
// 你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 思路
	// 原地排序，由于是顺序数组
	no := 1
	for i, v := range nums {
		// 第一个位置直接跳过
		if i == 0 {
			continue
		}
		// 比较当前值和前一个值
		// 相同跳过
		if v == nums[i-1] {
			continue
		} else {
			nums[no] = v
			no++
		}
	}
	return no
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3}
	fmt.Printf("个数 : %d", removeDuplicates(arr))
	fmt.Println()
	for k, v := range arr {
		fmt.Printf("位置: %d,数: %d", k, v)
		fmt.Println()
	}
}
