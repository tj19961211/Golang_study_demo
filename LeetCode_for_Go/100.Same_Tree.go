package main

// 2020-10-8

// Given two binary trees, write a function to check if they are the same or not.

// Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

// Example 1:

// Input:     1         1
//           / \       / \
//          2   3     2   3

//         [1,2,3],   [1,2,3]

// Output: true
// Example 2:

// Input:     1         1
//           /           \
//          2             2

//         [1,2],     [1,null,2]

// Output: false
// Example 3:

// Input:     1         1
//           / \       / \
//          2   1     1   2

//         [1,2,1],   [1,1,2]

// Output: false

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归出口，当两边均无值的时候返回 true
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil || p.Val != q.Val { // 返回 false 的情况
		return false
	} else { // 若值相同则进行下一层递归
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}
