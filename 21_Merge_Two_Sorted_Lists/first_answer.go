/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var al *ListNode
	var ll *ListNode

	tl1 := list1
	tl2 := list2

	for {
		if tl1 == nil && tl2 == nil {
			break
		}

		if tl1 == nil {
			if al == nil {
				al = &ListNode{ Val: tl2.Val, Next: nil }
				ll = al
			} else {
				ll.Next = &ListNode{ Val: tl2.Val, Next: nil }
				ll = ll.Next
			}
			tl2 = tl2.Next
		} else if tl2 == nil {
			if al == nil {
				al = &ListNode{ Val: tl1.Val, Next: nil }
				ll = al
			} else {
				ll.Next = &ListNode{ Val: tl1.Val, Next: nil }
				ll = ll.Next
			}
			tl1 = tl1.Next
		} else {
			if tl1.Val > tl2.Val {
				if al == nil {
					al = &ListNode{ Val: tl2.Val, Next: nil }
					ll = al
				} else {
					ll.Next = &ListNode{ Val: tl2.Val, Next: nil }
					ll = ll.Next
				}
				tl2 = tl2.Next
			} else {
				if al == nil {
					al = &ListNode{ Val: tl1.Val, Next: nil }
					ll = al
				} else {
					ll.Next = &ListNode{ Val: tl1.Val, Next: nil }
					ll = ll.Next
				}
				tl1 = tl1.Next
			}
		}


	}

	return al
}
