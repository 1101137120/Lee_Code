package dp

// ClimbStairs 爬樓梯
// LeetCode 70: https://leetcode.com/problems/climbing-stairs/
//
// 作法：動態規劃 + 空間優化
// 到第 n 階的方法數 = 到第 n-1 階 + 到第 n-2 階（一次爬 1 或 2 階）。
// 等同費氏數列，用兩個變數 a, b 滾動更新，避免開陣列。
// 時間 O(n)，空間 O(1)。
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
