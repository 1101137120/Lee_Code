package arrays

// TwoSum 兩數之和
// LeetCode 1: https://leetcode.com/problems/two-sum/
//
// 作法：雜湊表 (Hash Map)
// 遍歷陣列時，用 map 記錄「數值 -> 索引」。對每個 nums[i]，
// 檢查 target - nums[i] 是否已在 map 中，若有則回傳兩者索引。
// 時間 O(n)，空間 O(n)。
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
