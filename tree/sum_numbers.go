//给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
//
//例如，从根到叶子节点路径 1->2->3 代表数字 123。
//
//计算从根到叶子节点生成的所有数字之和。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例 1:
//
//输入: [1,2,3]
//    1
//   / \
//  2   3
//输出: 25
//解释:
//从根到叶子节点路径 1->2 代表数字 12.
//从根到叶子节点路径 1->3 代表数字 13.
//因此，数字总和 = 12 + 13 = 25.
//示例 2:
//
//输入: [4,9,0,5,1]
//    4
//   / \
//  9   0
// / \
//5   1
//输出: 1026
//解释:
//从根到叶子节点路径 4->9->5 代表数字 495.
//从根到叶子节点路径 4->9->1 代表数字 491.
//从根到叶子节点路径 4->0 代表数字 40.
//因此，数字总和 = 495 + 491 + 40 = 1026.

package tree

import (
	"fmt"
	"math"
)

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := 0
	var f func(node *TreeNode, sum int)
	f = func(node *TreeNode, sum int) {
		if r >= math.MaxInt64 || r <= math.MinInt64 {
			return
		}
		sum = sum*10 + node.Val
		if node.Left != nil {
			f(node.Left, sum)
		}
		if node.Right != nil {
			f(node.Right, sum)
		}
		if node.Right == nil && node.Left == nil {
			fmt.Printf("sum=%d node.Val=%d\n", sum, node.Val)
			r += sum
		}
	}
	f(root, 0)
	return r
}
