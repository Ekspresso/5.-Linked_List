// Функция принимает head связанного списка,
// создаёт новый перевёрнутый связанный список
// и возвращает head перевёрнутого связанного списка.

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newNode := new(ListNode)
	newNode.Val = head.Val
	newNode.Next = nil
	for head.Next != nil {
		head = head.Next
		node := new(ListNode)
		node.Val = head.Val
		node.Next = newNode
		newNode = node
	}
	return newNode
}
