package main

import (
	"strconv"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
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

		assert.Equal(t, actual, tc.expected)
	}
}

/**
 * Construct a binary search tree in golang.
 * https://leetcode.com/problems/recover-binary-search-tree/discuss/32539/Tree-Deserializer-and-Visualizer-for-Python
 */
func constructTree(tree string) *TreeNode {
	r := strings.NewReplacer("[", "", "]", "")
	leaves := strings.Split(r.Replace(tree), ",")
	nodes := make([]*TreeNode, len(leaves))
	for i, le := range leaves {
		if le != "null" {
			val, _ := strconv.Atoi(le)
			nodes[i] = &TreeNode{
				val,
				nil,
				nil,
			}
		}
	}
	rootNode := nodes[0]
	for i, node := range nodes {
		if node == nil {
			continue
		}
		if i+1 < len(nodes) {
			node.Left = nodes[i+1]
		}
		if i+2 < len(nodes) {
			node.Right = nodes[i+2]
		}
	}
	return rootNode
}
