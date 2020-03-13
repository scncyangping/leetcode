/*
@date : 2020/03/13
@author : YaPi
@desc : 多数元素
*/
package main

import "fmt"

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/majority-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 输入: [3,2,3]
// 输出: 3
// 双指针遍历
func majorityElement1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int)
	i := 0
	j := len(nums) - 1
	for i < j {
		m[nums[i]] += 1
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
		m[nums[j]] += 1
		if m[nums[j]] > len(nums)/2 {
			return nums[j]
		}

		// 遍历完了
		if j-i == 1 {
			break
		} else if j-i == 2 {
			i++
			m[nums[i]] += 1
			if m[nums[i]] > len(nums)/2 {
				return nums[i]
			}
			break
		}
		i++
		j--
	}
	return 0
}

// 先对数据进行排序,排序过后的列表中间的那一个数一定是多数元素
func majorityElement(nums []int) int {
	// sort(nums)

	return nums[len(nums)/2]
}

func main() {
	fmt.Println(majorityElement([]int{2, 2}))
}
