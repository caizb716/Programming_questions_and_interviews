package main

import (
	"fmt"
	"interview/cider"
)

func main() {
	str := "asfdfggfdroc"
	res := cider.LongestPalindrome(str)
	fmt.Println(res)
	// str := "aaf"
	// fmt.Println(str[0] == str[1])
}
