package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left, lok := getHeight(root.Left)
	right, rok := getHeight(root.Right)

	if !lok || !rok {
		return false
	}

	if left > right {
		return (left - right) <= 1
	}

	return (right - left) <= 1
}

func getHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	left, lok := getHeight(root.Left)
	right, rok := getHeight(root.Right)
	if !lok || !rok {
		return 0, false
	}
	if left > right {
		return left + 1, (left - right) <= 1
	}
	return right + 1, (right - left) <= 1
}
