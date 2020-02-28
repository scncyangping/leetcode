/*
@date : 2020/02/29
@author : YaPi
@desc :
*/
package main

import "fmt"

// 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return len(nums)
}

func main() {
	arr := []int{1}
	fmt.Printf("个数 : %d", removeElement(arr, 1))
	fmt.Println()
	for k, v := range arr {
		fmt.Printf("位置: %d,数: %d", k, v)
		fmt.Println()
	}
}
