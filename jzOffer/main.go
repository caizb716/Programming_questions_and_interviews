package main

import (
	"fmt"
	sort "jzOffer/10sort"
)

func main() {
	num := []int{3, 32, 321}
	res := sort.MinNumber(num)
	fmt.Println(res)
}
