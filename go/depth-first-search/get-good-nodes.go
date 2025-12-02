package main

import "math"

func doGetGoodNodes(node *TreeNode, maxVal int, nodes *[]*TreeNode) {
	if node == nil {
		return
	}

	if node.Value >= maxVal {
		maxVal = node.Value
		*nodes = append(*nodes, node)
	}

	doGetGoodNodes(node.Left, maxVal, nodes)
	doGetGoodNodes(node.Right, maxVal, nodes)
}

func GetGoodNodes(node *TreeNode) []*TreeNode {
	var nodes []*TreeNode
	doGetGoodNodes(node, math.MinInt, &nodes)

	return nodes
}
