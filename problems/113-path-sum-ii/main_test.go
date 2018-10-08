package main

import (
	"strconv"
	"strings"
	"testing"
)

type testcase struct {
	tree     string
	sum      int
	expected [][]int
}

func Test_pathSum(t *testing.T) {
	suit := []testcase{
		testcase{
			"[5,4,8,11,null,13,4,7,2,null,null,5,1]",
			22,
			[][]int{
				[]int{5, 4, 11, 2},
				[]int{5, 8, 4, 5},
			},
		},
	}
	for _, tc := range suit {
		rootNode := constructTree(tc.tree)
		actual := pathSum(rootNode, tc.sum)

		if actual != tc.expected {
			t.Errorf("tree: %v, sum: %v, expected: %v, actual: %v", tc.tree, tc.sum, tc.expected, actual)
		}
	}
}

/**
 * Construct a binary search tree in golang.
 * https://leetcode.com/problems/recover-binary-search-tree/discuss/32539/Tree-Deserializer-and-Visualizer-for-Python
 */
func constructTree(tree string) *TreeNode {
	r := strings.NewReplacer("[", "", "]", "")
	nodes := strings.Split(r.Replace(tree), ",")
	val, _ := strconv.Atoi(nodes[0])
	rootNode := TreeNode{
		val,
		nil,
		nil,
	}
	for _, node := range nodes[1:] {
		if node == "null" {

		}
	}
	return &rootNode
}
