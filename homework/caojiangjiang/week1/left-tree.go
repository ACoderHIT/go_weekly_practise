// 计算给定二叉树的所有左叶子之和。

// 示例：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
// 思路：分为左右子树，然后获取左子树的值

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sum-of-left-leaves
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumOfLeftLeaves(root *TreeNode) int {
	var total int
	sum(root, &total, 0)
	return total
}

func sum(root *TreeNode, total *int, flag int) int {  
	if (root == nil) {
			return 0
	}
	sum(root.Left, total, 1)   
	if (root.Left == nil && root.Right == nil && flag == 1) {
		 *total = *total + root.Val   
	} 
	sum(root.Right, total, 2)
	return *total
}