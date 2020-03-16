/*
@date : 2020/03/16
@author : YaPi
@desc : 字符串压缩
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/compress-string-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//  输入："aabcccccaaa"
// 输出："a2b1c5a3"

func compressString(S string) string {
	var temp byte
	var res strings.Builder
	var count int

	for i := range S {
		if S[i] == temp {
			count++
		}
		if S[i] != temp {
			if temp != 0 {
				res.WriteByte(temp)
				res.WriteString(strconv.Itoa(count))
			}
			temp, count = S[i], 1
		}
	}
	res.WriteByte(temp)
	res.WriteString(strconv.Itoa(count))
	if res.Len() < len(S) {
		return res.String()
	}
	return S
}

func main() {
	fmt.Println(compressString("aaaqqqwwwwwee"))
}
