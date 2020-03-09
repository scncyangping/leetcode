/*
@date : 2020/03/09
@author : YaPi
@desc : 买卖股票的最佳时机
*/
package main

import "fmt"

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//
//注意你不能在买入股票前卖出股票。
//
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

// 输入: [7,8,5,3,6,1]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

// 一、暴力破解
func maxProfit2(prices []int) int {
	// 要找到最好的买卖点，其实就是找两个位置的差值最大值
	// 必须满足买入价格小于卖出价格，其实就是判断两位置的值前者大于后者
	baseNum := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				// 差值
				num := prices[j] - prices[i]
				if num > baseNum {
					baseNum = num
				}
			}
		}
	}
	return baseNum
}

// 二、一次遍历
func maxProfit(prices []int) int {
	baseIndex := 0
	maxNum := 0
	for i := 0; i < len(prices); i++ {
		// 当前的值小于历史最低的值
		// 更新历史最低值的index
		if prices[i] < prices[baseIndex] {
			baseIndex = i
		} else {
			// 判断当前的值和历史最低值的差值是不是最大
			// 若是最大的，就更新最大利润
			if prices[i]-prices[baseIndex] > maxNum {
				maxNum = prices[i] - prices[baseIndex]
			}
		}

	}
	return maxNum
}

func main() {
	arr := []int{17, 16, 15, 8, 2}
	fmt.Println("money:", maxProfit(arr))
}
