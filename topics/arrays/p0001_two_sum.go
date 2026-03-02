package arrays

// TwoSum 兩數之和
// LeetCode 1: https://leetcode.com/problems/two-sum/
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
