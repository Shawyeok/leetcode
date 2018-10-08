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
	var nodes []int
	return pathSumRev(root, nodes, sum)
}

func pathSumRev(root *TreeNode, nodes []int, sum int) [][]int {
	var result [][]int
	nodes = append(nodes, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		result = append(result, nodes)
	} else {
		sum -= root.Val
		if sum < 0 {
			return result
		}
		if root.Right != nil {
			result = append(result, pathSumRev(root.Right, nodes, sum)...)
		}
		if root.Left != nil {
			result = append(result, pathSumRev(root.Left, nodes, sum)...)
		}
	}
	return result
}
