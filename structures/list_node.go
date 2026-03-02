package structures

import "fmt"

// ListNode 單向鏈結串列節點 (LeetCode 預設結構)
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints 將鏈結串列轉為 []int
func List2Ints(head *ListNode) []int {
	limit := 100
	times := 0
	res := []int{}
	for head != nil {
		times++
		if times > limit {
			panic(fmt.Sprintf("鏈條深度超過 %d，可能出現環狀鏈條", limit))
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// Ints2List 將 []int 轉為鏈結串列
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{Val: nums[0]}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}
	return res
}
