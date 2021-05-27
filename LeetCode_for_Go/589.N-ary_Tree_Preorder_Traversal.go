package main

// 树的前序遍历，遍历路径为 根节点、左孩子、右孩子
func preorder(root *Node) []int {
	res := []int{}

	if root == nil {
		return res
	}

        // 接入初始的根节点
	stack := []*Node{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, r.Val)
		n := len(r.Children)
		for i := n - 1; i >= 0; i-- {
			stack = append(stack, r.Children[i])
		}
	}
	return res
}
