// Функция принимает head связанного списка и x.
// Учитывая head связанного списка и значение x,
// функция делит его таким образом, чтобы все узлы,
// меньшие, чем x, предшествовали узлам,
// большим или равным x.
// Функция сохраняет исходный относительный
// порядок узлов в каждом из двух разделов.
// Input: head = [1,4,3,2,5,2], x = 3
// Output: [1,2,2,4,3,5]

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := new(ListNode)
	slow.Val = -1000
	slow.Next = head
	fast := slow
	curr := head
	head = slow
	for curr != nil {
		if curr.Val < x {
			fast.Next = fast.Next.Next
			curr.Next = slow.Next
			slow.Next = curr
			slow = slow.Next
			if fast.Next != nil && fast.Val < x {
				fast = fast.Next
				curr = fast.Next
			} else {
				curr = fast.Next
			}
		} else {
			fast = fast.Next
			curr = curr.Next
		}
	}
	return head.Next
}
