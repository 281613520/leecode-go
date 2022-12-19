package context292

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ans := 0

	// sum nums
	var dfs func(node *TreeNode) (int, int)

	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftSum, leftNums := dfs(node.Left)
		rightSum, rightNums := dfs(node.Right)

		avager := (leftSum + rightSum + node.Val) / (leftNums + rightNums + 1)
		if avager == node.Val {
			ans++
		}

		return leftSum + rightSum + node.Val, leftNums + rightNums + 1
	}

	dfs(root)

	return ans
}
