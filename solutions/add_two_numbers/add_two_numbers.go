package add_two_numbers

import (
	common "github.com/petrogdev/leetcode/common"
)

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	dummyHead := &common.ListNode{}
	current := dummyHead
	carry := 0

	if l1 == nil && l2 == nil {
		return nil
	}

	for carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		current.Next = &common.ListNode{Val: sum % 10}
		current = current.Next
	}
	return dummyHead.Next
}

func addTwoNumbersR(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	next := addTwoNumbersR(l1.Next, l2.Next)

	if sum >= 10 {
		sum %= 10
		next = addTwoNumbersR(next, &common.ListNode{
			Val: 1,
		})
	}

	return &common.ListNode{
		Val:  sum,
		Next: next,
	}
}
