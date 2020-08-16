package leetcode

/*
110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedSub(isBalancedDepth(root.Left), isBalancedDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalancedDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return isBalancedMax(isBalancedDepth(root.Left), isBalancedDepth(root.Right)) + 1
}

func isBalancedMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isBalancedSub(x, y int) int {
	if y > x {
		return y - x
	}
	return x - y
}
