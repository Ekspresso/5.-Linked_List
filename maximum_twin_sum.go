// Функция определяет максимальную сумму двойников в связанном списке
// длинны n, где n чётно.
// Двойник для i-того элемента - это i-тый элемент с конца.

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Создание перевёрнутого связанного списка.
	reverse, k := reverseList(head)
	max := 0
	// Проход до половины каждого списка (оригинального и перевёрнутого)
	for (k+1)/2 > 0 {
		if reverse.Val+head.Val > max {
			max = reverse.Val + head.Val
		}
		reverse = reverse.Next
		head = head.Next
		k--
	}
	return max
}

func reverseList(head *ListNode) (*ListNode, int) {
	reverse := new(ListNode)
	normal := head
	reverse.Val = normal.Val
	reverse.Next = nil
	k := 1
	for normal.Next != nil {
		normal = normal.Next
		node := new(ListNode)
		node.Val = normal.Val
		node.Next = reverse
		reverse = node
		k++
	}
	return reverse, k
}
