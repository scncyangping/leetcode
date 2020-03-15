/*
@date : 2020/03/15
@author : YaPi
@desc : 岛屿的最大面积
*/
package main

import (
	"container/list"
	"fmt"
)

// 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
//
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-area-of-island
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//[[0,0,1,0,0,0,0,1,0,0,0,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,1,1,0,1,0,0,0,0,0,0,0,0],
// [0,1,0,0,1,1,0,0,1,0,1,0,0],
// [0,1,0,0,1,1,0,0,1,1,1,0,0],
// [0,0,0,0,0,0,0,0,0,0,1,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,0,0,0,0,0,0,1,1,0,0,0,0]]
//
// 对于上面这个给定矩阵应返回 6。注意答案不应该是11，
// 因为岛屿只能包含水平或垂直的四个方向的‘1’。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-area-of-island
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type land struct {
	row int
	col int
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	y := len(grid)
	x := len(grid[0])
	res := 0
	// i 为第几行
	for i, vi := range grid {
		// j 为第几列
		for j, kj := range vi {
			if kj == 0 {
				// 说明是海水
				continue
			}
			// 说明是岛屿
			// 是岛屿就需要把这个加入队列
			l := list.New()
			l.PushBack(&land{row: i, col: j})
			grid[i][j] = 0
			baseNum := 0
			for l.Front() != nil {
				// 取出一个陆地
				first := l.Front()
				l.Remove(first)
				lan := first.Value.(*land)
				// 把当前的陆地设置为海水 避免下一次再次计算
				// 同时把岛屿数加1
				// 判断这个陆地周围是否有陆地，有陆地就加入队列
				// 判断上面是否有陆地
				baseNum++
				if lan.row > 0 && grid[lan.row-1][lan.col] == 1 {
					l.PushBack(&land{lan.row - 1, lan.col})
					grid[lan.row-1][lan.col] = 0
				}
				// 判断下方是否有岛屿
				if lan.row < y-1 && grid[lan.row+1][lan.col] == 1 {
					l.PushBack(&land{lan.row + 1, lan.col})
					grid[lan.row+1][lan.col] = 0
				}
				// 判断左边是否有岛屿
				if lan.col > 0 && grid[lan.row][lan.col-1] == 1 {
					l.PushBack(&land{lan.row, lan.col - 1})
					grid[lan.row][lan.col-1] = 0
				}
				// 判断右边是否有岛屿
				if lan.col < x-1 && grid[lan.row][lan.col+1] == 1 {
					l.PushBack(&land{lan.row, lan.col + 1})
					grid[lan.row][lan.col+1] = 0

				}
			}
			if baseNum > res {
				res = baseNum
			}
		}
	}
	return res
}
func print(grid [][]int) {
	for _, iv := range grid {
		for _, jv := range iv {
			fmt.Print(jv, "   ")
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------")
}
func main() {

	// 		{0,0,1,0,0,0,0,1,0,0,0,0,0},
	//		{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//		{0,1,1,0,1,0,0,0,0,0,0,0,0},
	//		{0,1,0,0,1,1,0,0,1,1,1,1,0},
	//		{0,1,0,0,1,1,0,0,1,1,1,0,0},
	//		{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//		{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//		{0,0,0,0,0,0,0,1,1,0,0,0,0}
	arr := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	fmt.Println(maxAreaOfIsland(arr))
}
