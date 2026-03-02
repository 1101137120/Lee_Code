package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a := 1
	b := 2
	for i := 2; i < n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}
