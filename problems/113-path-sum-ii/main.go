package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func pathSum(root *TreeNode, sum int) [][]int {
	var vals []int
	return pathSumRev(root, vals, sum)
}

func pathSumRev(node *TreeNode, vals []int, sum int) [][]int {
	var result [][]int
	vals = append(vals, node.Val)
	sum -= node.Val
	if node.Left == nil && node.Right == nil {
		if sum == 0 {
			result = append(result, vals)
		}
	} else {
		if node.Left != nil {
			result = append(result, pathSumRev(node.Left, vals, sum)...)
		}
		if node.Right != nil {
			result = append(result, pathSumRev(node.Right, vals, sum)...)
		}
	}
	return result
}
