// Функции передаётся head связанного списка.
// Функция возвращает средний узел связанного списка.
// Если средних узлов два, то возвращается второй из них.

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	point := head
	k := 0
	for point.Next != nil {
		k++
		point = point.Next
	}
	k = (k + 1) / 2
	for k > 0 {
		head = head.Next
		k--
	}
	return head
}
