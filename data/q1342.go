/*
@date : 2020/02/29
@author : YaPi
@desc : 将数字变成 0 的操作次数
*/
package main

import "fmt"

// 给你一个非负整数 num ，请你返回将它变成 0
// 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1
// 输入：num = 14
//输出：6
//解释：
//步骤 1) 14 是偶数，除以 2 得到 7 。
//步骤 2） 7 是奇数，减 1 得到 6 。
//步骤 3） 6 是偶数，除以 2 得到 3 。
//步骤 4） 3 是奇数，减 1 得到 2 。
//步骤 5） 2 是偶数，除以 2 得到 1 。
//步骤 6） 1 是奇数，减 1 得到 0 。
//链接：https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero

func numberOfSteps(num int) int {
	var (
		// 步数
		flag int
	)
	for {
		if num == 0 {
			break
		}
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		flag++
	}
	return flag
}

func main() {
	fmt.Println(numberOfSteps(0))
	fmt.Println(numberOfSteps(1))
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))

}
