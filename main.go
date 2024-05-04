package main

import "fmt"

func main() {
	fmt.Println(Divisors([]int{42, 12, 18}))
}

func Divisors(n []int) []int {
	min := n[0]
	for _, num := range n {
		if num < min {
			min = num
		}
	}

	var c []int

	for i := 2; i <= min; i++ {
		isCommon := true
		for _, num := range n {
			if num%i != 0 {
				isCommon = false
				break
			}
		}
		if isCommon {
			c = append(c, i)
		}
	}

	return c
}
