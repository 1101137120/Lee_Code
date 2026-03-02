package dp

// ClimbStairs 爬樓梯
// LeetCode 70: https://leetcode.com/problems/climbing-stairs/
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
