/*
 * @lc app=leetcode.cn id=257 lang=golang
 * @lcpr version=30204
 *
 * [257] 二叉树的所有路径
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
var results []string

func binaryTreePaths(root *TreeNode) []string {
	results = []string{}
	scanNode(root, "")
	return results
}

func scanNode(node *TreeNode, path string) {
	if node == nil {
		return
	}

	if path == "" {
		path = strconv.Itoa(node.Val)
	} else {
		path += fmt.Sprintf("->%d", node.Val)
	}
	if node.Left == nil && node.Right == nil {
		results = append(results, path)
		return
	}

	scanNode(node.Left, path)
	scanNode(node.Right, path)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

