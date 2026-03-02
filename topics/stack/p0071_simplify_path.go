package stack

import "strings"

// SimplifyPath 簡化路徑
// LeetCode 71: https://leetcode.com/problems/simplify-path/
//
// 作法：堆疊 (Stack)
// 以 "/" 分割路徑，遍歷每個 segment：空字串與 "." 略過；
// ".." 則從堆疊彈出上一層；其餘視為目錄名壓入堆疊。
// 最後以 "/" 連接堆疊內容，並加上開頭 "/"。
// 時間 O(n)，空間 O(n)。
func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	st := []string{}
	for _, p := range parts {
		switch p {
		case "", ".":
			continue
		case "..":
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
			continue
		default:
			st = append(st, p)
		}
	}
	return "/" + strings.Join(st, "/")
}
