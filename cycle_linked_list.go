// Функция определяет есть ли в связанном списке цикл.

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p := make(map[*ListNode]bool)
	for head != nil {
		if p[head] {
			return true
		}
		p[head] = true
		head = head.Next
	}
	return false
}
