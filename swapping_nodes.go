// Функция получает head связанного списка и значение k.
// Функция меняет местами k-ый элемент с начала списка
// и k-ый элемент списка с конца.
// Считается, что список проиндексирован с 1 (k > 0).

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	k1 := k - 1
	n := 1
	point1 := head
	point2 := head
	for point1.Next != nil {
		n++
		point1 = point1.Next
	}
	point1 = head
	for k1 > 0 || n-k > 0 {
		if k1 > 0 {
			point1 = point1.Next
			k1--
		}
		if n-k > 0 {
			point2 = point2.Next
			n--
		}
	}
	point1.Val, point2.Val = point2.Val, point1.Val
	return head
}
