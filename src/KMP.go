package main

import (
	"fmt"
)

func main() {
	str := "abcabcdxyz"
	fmt.Println(getNext(str))

}

func KMP(origin, pattern string) int {
	originSlice := []rune(origin)
	patternSlice := []rune(pattern)

	i := 0 // 主串的位置
	j := 0 // 模式串的位置

	next := getNext(pattern)

	while i < len(originSlice) && j < len(patternSlice) {
		if j == -1 || t[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}

	}
	if (j == len(patternSlice) {
		return i - j
	} else {
		return -1
	} 
}

// KMP 的next数组
func getNext(pattern string) []int {
	bytes := []rune(pattern)
	length := len(pattern)
	next := make([]int, length)
	next[0] = -1
	j := 0
	k := -1

	for j < length-1 {
		if k == -1 || bytes[j] == bytes[k] {
			j++
			k++
			// 	一般来说他们不相等才是好事
			if (bytes[j] == bytes[k]) {
				next[j] = next[k]
			} else {
				next[j] = k
			}
			fmt.Printf("next[%d] = %d\n", j, k)
		} else {
			fmt.Printf("%d = next[%d]\n", next[k], k)
			k = next[k]
		}
	}
	return next
}

