package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
}

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

