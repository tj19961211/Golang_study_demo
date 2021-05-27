package main

func increasingBST(root *TreeNode) {
	min, _ := helper(root)
	return min
}

func helper(node *TreeNode) (*TreeNode, *TreeNode) {
	min, max := node, node
	// 向左孩子进发，因为树的左边总是比根小，所以通过不断遍历左孩子，
	// 来调整树从最小值起，右边孩子一直存在，无左孩子
	if node.Left != nil {
		lMin, lMax := helper(node.Left)
		min = lMin
		lMax.Right = node
	}

	// 同理，选出树的最大值，调整树的左孩子
	if node.Right != nil {
		rMin, rMax := helper(node.Right)
		max = rMax
		node.Right = rMin
	} 
	node.Left = nil
	return min, max
}