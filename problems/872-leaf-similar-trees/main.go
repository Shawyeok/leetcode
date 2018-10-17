package main

func main() {
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := getLeafsRev(root1)
	leafs2 := getLeafsRev(root2)

	return compare(leafs1, leafs2)
}

func getLeafsRev(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	return append(getLeafsRev(root.Left), getLeafsRev(root.Right)...)
}

func compare(leafs1, leafs2 []int) bool {
	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}
