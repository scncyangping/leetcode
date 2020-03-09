/*
@date : 2020/03/04
@author : YaPi
@desc : 腐烂的橘子
*/
package main

import (
	"fmt"
)

// 在给定的网格中，每个单元格可以有以下三个值之一：
//
//值 0 代表空单元格；
//值 1 代表新鲜橘子；
//值 2 代表腐烂的橘子。
//每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
//
//返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。
//
//链接：https://leetcode-cn.com/problems/rotting-oranges

type Node struct {
	// 列
	col int
	// 行
	row int
}

func orangesRotting(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return -1
	}
	var (
		yL    int
		xL    int
		num   int
		time  int
		queue = make([]Node, 0)
	)
	// 多少行
	yL = len(grid)
	// 多少列
	xL = len(grid[0])
	// 遍历将所有节点添加到队列中
	for i, v := range grid {
		for j, s := range v {
			if s == 2 {
				node := Node{
					row: i,
					col: j,
				}
				queue = append(queue, node)
			}
			if s == 1 {
				num++
			}
		}
	}

	for {
		tmp := 0 //用来记录每分钟变坏的橘子数，当不再新增则结束
		newNode := make([]Node, 0)
		for len(queue) > 0 {
			// 取出一个被感染的橘子
			loc := queue[0]
			queue = queue[1:]
			// 判断它上方是否有橘子，及橘子的状态
			if loc.row > 0 {
				if grid[loc.row-1][loc.col] == 1 {
					tmp++
					num--
					grid[loc.row-1][loc.col] = 2
					var n Node
					n.row = loc.row - 1
					n.col = loc.col
					newNode = append(newNode, n)
				}
			}
			// 判断下方是否有橘子
			if loc.row < yL-1 {
				if grid[loc.row+1][loc.col] == 1 {
					tmp++
					num--
					grid[loc.row+1][loc.col] = 2
					var n Node
					n.row = loc.row + 1
					n.col = loc.col
					newNode = append(newNode, n)
				}
			}
			// 判断左边是否还有橘子
			if loc.col > 0 {
				if grid[loc.row][loc.col-1] == 1 {
					tmp++
					num--
					grid[loc.row][loc.col-1] = 2
					var n Node
					n.row = loc.row
					n.col = loc.col - 1
					newNode = append(newNode, n)
				}
			}
			// 判断右边是否还有橘子
			if loc.col < xL-1 {
				if grid[loc.row][loc.col+1] == 1 {
					tmp++
					num--
					grid[loc.row][loc.col+1] = 2
					var n Node
					n.row = loc.row
					n.col = loc.col + 1
					newNode = append(newNode, n)
				}
			}
		}
		if tmp == 0 {
			break
		}
		queue = newNode
		time++
	}
	if num > 0 {
		return -1
	} else {
		return time
	}
}

func main() {
	ass := [][]int{{1, 2}}

	fmt.Println(orangesRotting(ass))

}
