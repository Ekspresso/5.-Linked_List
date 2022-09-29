// Функции передаётся head связанного списка, в котором
// хранится двоичное представление числа (наиболее значимый бит в начале списка).
// Функция возвращает десятичное представление числа.

package linkedlist

import "math"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	pointer := head
	var k float64
	k = 0
	ret := 0
	for pointer.Next != nil {
		k += 1
		pointer = pointer.Next
	}
	pointer = head
	for k >= 0 {
		if pointer.Val == 1 {
			ret += int(math.Pow(2, k))
		}
		pointer = pointer.Next
		k--
	}
	return ret
}
