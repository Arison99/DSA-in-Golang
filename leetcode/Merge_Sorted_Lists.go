package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to act as the head of the merged list
	dummy := &ListNode{}
	current := dummy

	// Iterate through both lists
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// If there are remaining nodes in either list, append them
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the merged list, which starts at dummy.Next
	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example usage
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList)
}
