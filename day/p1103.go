/*
@date : 2020/03/05
@author : YaPi
@desc : 分糖果
*/
package main

import (
	"fmt"
)

// 我们买了一些糖果 candies，打算把它们分给排好队的 n = num_people 个小朋友。
//
//给第一个小朋友 1 颗糖果，第二个小朋友 2 颗，依此类推，直到给最后一个小朋友 n 颗糖果。
//
//然后，我们再回到队伍的起点，给第一个小朋友 n + 1 颗糖果，第二个小朋友 n + 2 颗，依此类推，直到给最后一个小朋友 2 * n 颗糖果。
//
//重复上述过程（每次都比上一次多给出一颗糖果，当到达队伍终点后再次从队伍起点开始），直到我们分完所有的糖果。注意，就算我们手中的剩下糖果数不够（不比前一次发出的糖果多），这些糖果也会全部发给当前的小朋友。
//
//返回一个长度为 num_people、元素之和为 candies 的数组，以表示糖果的最终分发情况（即 ans[i] 表示第 i 个小朋友分到的糖果数）
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/distribute-candies-to-people

//输入：candies = 7, num_people = 4
//输出：[1,2,3,1]
//解释：
//第一次，ans[0] += 1，数组变为 [1,0,0,0]。
//第二次，ans[1] += 2，数组变为 [1,2,0,0]。
//第三次，ans[2] += 3，数组变为 [1,2,3,0]。
//第四次，ans[3] += 1（因为此时只剩下 1 颗糖果），最终数组变为 [1,2,3,1]。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/distribute-candies-to-people
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func distributeCandies2(candies int, num_people int) []int {
	arr := make([]int, num_people)

	num := 1
out:
	for {
		// 当前分配数量
		// 依次给每一个分配
		newNum := 0
		for i := 0; i < num_people; i++ {
			n := i + num
			// 分完了
			if candies-n <= 0 {
				arr[i] = arr[i] + candies
				break out
			}
			// 分一次 num + 1
			newNum = n
			candies -= n
			arr[i] = arr[i] + n
		}
		num = newNum + 1

	}
	return arr
}

// 10 3
func distributeCandies(candies int, num_people int) []int {
	arr := make([]int, num_people)
	num := 1
	for candies > 0 {
		n := 0
		if candies-num >= 0 {
			n = num
		} else {
			n = candies
		}
		arr[(num-1)%num_people] += n
		candies -= num
		num++
	}
	return arr
}

func main() {
	candies := 10
	numPeople := 3
	arr := distributeCandies(candies, numPeople)
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
