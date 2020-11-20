package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Build(arr ...int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val: arr[0],
	}
	p := head

	for i := 1; i < len(arr); i++ {
		node := &ListNode{
			Val: arr[i],
		}
		p.Next = node
		p = node
	}

	return head
}

func Print(head *ListNode) {
	msg := ""

	p := head
	for p != nil {
		msg += fmt.Sprintf("%d", p.Val) + "->"
		p = p.Next
	}

	fmt.Println(msg)
}
