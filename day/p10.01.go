/*
@date : 2020/03/03
@author : YaPi
@desc : 合并排序的数组
*/
package main

// 给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
//
// 初始化 A 和 B 的元素数量分别为 m 和 n。
//
//链接：https://leetcode-cn.com/problems/sorted-merge-lcci

func merge(A []int, m int, B []int, n int) {
	// 直接将b添加到A
	// 然后进行一次冒泡排序就行了
	A = append(A[:m], B...)

	for i := 0; i < m+n-1; i++ {
		for j := 0; j < m+n-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}
}

func main() {
	a := []int{2, 0}
	b := []int{1}

	merge(a, 1, b, 1)

}
