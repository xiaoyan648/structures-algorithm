/*
 * @lc app=leetcode.cn id=94 lang=golang
 * @lcpr version=30204
 *
 * [94] 二叉树的中序遍历
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var scanNode func(*TreeNode)

	scanNode = func(node *TreeNode) {
		if node == nil {
			return
		}

		scanNode(node.Left)
		result = append(result, node.Val)
		scanNode(node.Right)
	}

	scanNode(root)

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

