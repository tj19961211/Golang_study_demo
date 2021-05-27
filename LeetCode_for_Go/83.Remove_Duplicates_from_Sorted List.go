package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates_1(head *ListNode) *ListNode {
	curr := head
	res := head
	for head != nil {
		// 如果 node 的 val 与当前的 val 不相等，则指向当前 node 的指针向后移动
		if head.Val != curr.Val {
			curr.Next = head
			curr = head
		} else {
			// 倘若前后两个 node val 相等则进行链表删减
			head, head.Next = head.Next, nil
		}
	}
	return res
}

// 方法二：比方法一更简单理解
func deleteDuplicates_2(head *ListNode) *ListNode {
	// 设定指向链表头的 ListNode ， 用作最后的返回
	res := head

	//判断 node是否有值与是否为最后一个 node
	for head != nil && head.Next != nil {

		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}
