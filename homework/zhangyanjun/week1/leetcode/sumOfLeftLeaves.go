package main

//问题：
//计算给定二叉树的所有左叶子之和。
//
//示例：
//
// 		   3
//		  / \
//		 9  20
//		   /  \
//		  15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24


//思路：
//这道题其实就是一个二叉树的遍历问题，在遍历的时候如果一个节点是其父节点的左孩子并且该节点没有左右孩子，那么这个节点就是我们想要的，
//把这个节点的val加进全局val中即可
//分析：时间复杂度为O(n)。空间复杂度为O(logn)（函数栈的空间）也就是树的高度。 n代表节点个数

//但是二叉树的遍历也分为很多种，最简单的就是递归形式的前中后序遍历，该题目难度为easy，并且没有强调实现方法，故可以用递归
//如果该题测试用例严格一些，树的高度增大，那么使用递归的做法很可能函数栈溢出，于是需要自己构建函数栈 来进行前中后序的遍历，或者是用queue按层次遍历
//分析：时间复杂度为O(n)， 空间复杂度为O(logn)（自己new的栈的空间），也就是树的高度。 n代表节点个数。比起递归解法的优势就是可以防止栈溢出

//如果这道题对空间复杂度的要求是O(1) 那么就要用到Morris遍历来做
// 分析：时间复杂度为O(n)。空间复杂度为O(1)（常数级别的变量）。 n代表节点个数

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return run(root, 2)
}
func run(root *TreeNode, flag int) int {
	if root.Left == nil && root.Right == nil && flag == 1 {
		return root.Val
	}
	res := 0
	if root.Left != nil {
		res = res + run(root.Left, 1)
	}
	if root.Right != nil {
		res = res + run(root.Right, 2)
	}
	return res;
}