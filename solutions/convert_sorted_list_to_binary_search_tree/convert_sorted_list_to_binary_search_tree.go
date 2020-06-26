package convert_sorted_list_to_binary_search_tree

import (
	common "github.com/petrogdev/leetcode/common"
)

type ListNode *common.ListNode
type TreeNode *common.TreeNode

func sortedListToBST(head *common.ListNode) *common.TreeNode {
	return transMidToRoot(head, nil)
}

func transMidToRoot(begin, end *common.ListNode) *common.TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &common.TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &common.TreeNode{
		Val:   mid.Val,
		Left:  transMidToRoot(begin, mid),
		Right: transMidToRoot(mid.Next, end),
	}
}