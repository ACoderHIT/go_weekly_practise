package main

import "math"

type ListNode struct {
	Val int
	Next *ListNode
}

//给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
//
//你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//进阶:
//
//如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
//示例:
//
//输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出: 7 -> 8 -> 0 -> 7


//基础解法 翻转链表做加法之后再翻转回去
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reserveListNode(l1)
	l2 = reserveListNode(l2)
	jin := 0
	ans :=  &ListNode{
		Val:-1,
		Next:nil,
	}
	temp := ans
	for ;l1 != nil && l2 != nil; {
		ans.Next = &ListNode{
			Val:(l1.Val + l2.Val + jin) % 10,
			Next:nil,
		}
		jin = (l1.Val + l2.Val + jin) / 10
		ans = ans.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for ;l2 != nil; {
		ans.Next = &ListNode{
			Val:(l2.Val + jin) % 10,
			Next:nil,
		}
		jin = (l2.Val + jin) / 10
		ans = ans.Next
		l2 = l2.Next
	}
	for ;l1 != nil; {
		ans.Next = &ListNode{
			Val:(l1.Val + jin) % 10,
			Next:nil,
		}
		jin = (l1.Val + jin) / 10
		ans = ans.Next
		l1 = l1.Next
	}
	if (jin == 1) {
		ans.Next = &ListNode{
			Val:1,
			Next:nil,
		}
	}
	return reserveListNode(temp.Next)

}

func reserveListNode (l *ListNode) *ListNode {
	var pre = (*ListNode)(nil)
	for ;l != nil; {
		next := l.Next
		l.Next = pre
		pre = l
		l = next
	}
	return pre
}

//follow up
//思路：先把链表补全，再递归相加两个链表
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = buildList(l1, l2, getLen(l1), getLen(l2))
	res, jin := run(l1, l2);
	if (jin == 0) {
		return res
	}
	return &ListNode {
		Val:jin,
		Next:res,
	}
}

func run (l1 *ListNode, l2 *ListNode) (res *ListNode, jin int) {
	jin = 0
	if (l1 == nil && l2 == nil) {
		return nil, 0
	}
	tempres, j := run(l1.Next, l2.Next)
	res = &ListNode {
		Val:(l1.Val + l2.Val + j)%10,
		Next:tempres,
	}
	jin = (l1.Val + l2.Val + j)/10
	return
}

func buildList(l1 *ListNode, l2 *ListNode, len1 int, len2 int) (r1 *ListNode, r2 *ListNode)  {
	len := math.Abs(float64(len2-len1))
	t := &ListNode {
		Val:0,
		Next:nil,
	}
	h := t
	for i := 0; i < int(len - 1); i++ {
		t.Next = &ListNode {
			Val:0,
			Next:nil,
		}
		t = t.Next
	}
	if (len1 < len2) {
		t.Next = l1
		r1 = h
		r2 = l2
	} else if (len1 > len2) {
		t.Next = l2
		r2 = h
		r1 = l1
	} else {
		r2 = l2
		r1 = l1
	}
	return
}

func getLen(l1 *ListNode) (n int) {
	for {
		if l1 == nil {
			return
		}
		n++
		l1 = l1.Next
	}
}

