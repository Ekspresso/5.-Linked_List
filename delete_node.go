// Функция получает узел, который необходимо удалить из связанного списка.
// Этот узел не является последним.
// Не подразумевается удаление узла из памяти.
// Происходит удаление значения из списка
// и кол-во элементов в списке уменьшается на 1.
// Значения перед и после узла остаются теми же и в таком же порядке.

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteNode(node *ListNode) {
//     for node.Next.Next != nil {
//         node.Val = node.Next.Val
//         node = node.Next
//     }
//     node.Val = node.Next.Val
//     node.Next = nil
// }

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
	next.Next = nil
}
