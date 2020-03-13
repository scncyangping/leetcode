/*
@date : 2020/03/12
@author : YaPi
@desc : 字符串的最大公因子
*/
package main

import "fmt"

// 对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，
// 我们才认定 “T 能除尽 S”。
//
//返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"

// 输入：str1 = "ABABAB", str2 = "ABAB"
// 输出："AB"

// 输入：str1 = "LEET", str2 = "CODE"
// 输出：""
// 一 暴力破解
func gcdOfStrings2(str1 string, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	// 遍历第一个字符串
	// 依次设置值
	baseX := ""
	for i := 1; i <= len(str1); i++ {
		// 获取字符串长度
		x := str1[:i]
		if len(str2)%i != 0 {
			continue
		}
		if len(str1)%i != 0 {
			continue
		}
		baseJ := x
		for k := 1; k < len(str2)/i; k++ {
			baseJ += x
		}
		if baseJ == str2 {
			baseK := x
			for k := 1; k < len(str1)/i; k++ {
				baseK += x
			}
			if baseK == str1 {
				baseX = x
			}
		}
	}
	return baseX
}

// 二、辗转相除法
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	var remainder = modOfStrings(str1, str2)
	if remainder == "" {
		return str2
	} else if remainder == str1 {
		return ""
	}
	return gcdOfStrings(str2, remainder)
}
func modOfStrings(str1, str2 string) string {
	length := len(str2)
	var remainder = str1
	for {
		if len(remainder) < length || remainder[:length] != str2 {
			break
		}
		remainder = remainder[length:]
	}
	return remainder
}

func main() {
	fmt.Println(gcdOfStrings("ABABABCDEF", "ABABDFER"))
}
