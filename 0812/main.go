package main

import (
	"fmt"
	"strings"
)

func longString(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	if len(s) == 2 && s[1] != s[0] {
		return 2
	}
	pre := 0
	max := 1
	for index := 2; index < len(s); index++ {
		if strings.Contains(s[pre:index], s[index:index+1]) {
			if max < len(s[pre:index]) {
				max = len(s[pre:index])
			}
			pre = index
		}
	}
	return max
}

func main() {
	var s = "aab"
	fmt.Println(longString(s))
}
