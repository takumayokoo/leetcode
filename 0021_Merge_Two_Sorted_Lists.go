package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergeNode := new(ListNode)
	root := mergeNode

	for {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				mergeNode.Val = list1.Val
				list1 = list1.Next
			} else {
				mergeNode.Val = list2.Val
				list2 = list2.Next
			}

		} else if list1 != nil && list2 == nil {
			mergeNode.Val = list1.Val
			list1 = list1.Next
		} else if list1 == nil && list2 != nil {
			mergeNode.Val = list2.Val
			list2 = list2.Next
		} else {
			break
		}

		if list1 != nil || list2 != nil {
			mergeNode.Next = new(ListNode)
			mergeNode = mergeNode.Next
		}
	}

	return root.Next
}
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var root *ListNode
	if list1.Val < list2.Val {
		root = list1
		root.Next = mergeTwoLists(list1.Next, list2)
	} else {
		root = list2
		root.Next = mergeTwoLists(list1, list2.Next)
	}

	return root
}
