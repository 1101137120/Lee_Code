package structures

import "fmt"

// TreeNode 二元樹節點 (LeetCode 預設結構)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const null = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode (LeetCode 層序表示法)
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: ints[0]}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != null {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != null {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
		if i >= n {
			break
		}
	}
	return root
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	panic(fmt.Sprintf("%d 不存在於 %v 中", val, nums))
}

// PreIn2Tree 把 preorder 和 inorder 切片轉換成二元樹
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中兩個切片長度不相等")
	}
	if len(in) == 0 {
		return nil
	}
	res := &TreeNode{Val: pre[0]}
	if len(in) == 1 {
		return res
	}
	idx := indexOf(res.Val, in)
	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])
	return res
}

// Tree2Preorder 把二元樹轉換成 preorder 切片
func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)
	return res
}
