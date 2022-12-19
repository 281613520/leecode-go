package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	ans := 0
	maxDepth := 0

	var dfs = func(node *TreeNode, depth int) {}

	dfs = func(node *TreeNode, depth int) {
		depth++
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				ans = 0
				ans = node.Val
				maxDepth = depth
			} else if depth == maxDepth {
				ans += node.Val
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, depth)
		}

		if node.Right != nil {
			dfs(node.Right, depth)
		}

	}

	dfs(root, 0)

	return ans
}
