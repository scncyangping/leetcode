/*
@date : 2020/03/17
@author : YaPi
@desc : 拼写单词
*/
package main

import (
	"fmt"
	"strings"
)

// 给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
//
//假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。
//
//注意：每次拼写时，chars 中的每个字母都只能用一次。
//
//返回词汇表 words 中你掌握的所有单词的 长度之和。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func countCharacters(words []string, chars string) int {
	if len(chars) < 1 {
		return 0
	}
	M := make(map[string]int)
	// 塞值进去
	for i := 1; i <= len(chars); i++ {
		M[chars[i-1:i]] += 1
	}
	b := strings.Builder{}
	for _, v := range words {
		arr := []rune(v)
		flag := false
		mm := cloneTags(M)
		for i := range arr {
			if mm[string(arr[i])] != 0 {
				mm[string(arr[i])] -= 1
				continue
			}
			flag = true
		}
		if !flag {
			b.WriteString(v)
		}
	}
	return b.Len()
}
func cloneTags(tags map[string]int) map[string]int {
	cloneTags := make(map[string]int)
	for k, v := range tags {
		cloneTags[k] = v
	}
	return cloneTags
}
func main() {
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}
