// Функция принимает head двух связанных списков.
// Функция возвращает узел, в котором эти списки пересекаются
// (т.е. первый узел, который встречается в обоих списках).

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointers := make(map[*ListNode]int)
	pointers[headA] += 1
	pointers[headB] += 1
	// Проход по первому списку и запись его узлов в отображение.
	if pointers[headA] < 2 {
		for headA.Next != nil {
			pointers[headA.Next] += 1
			if pointers[headA.Next] > 1 {
				break
			}
			headA = headA.Next
		}
	}
	// Проход по второму списку и запись его значений в отображение
	// до момента его пересечения с первым списком
	if headA.Next == nil {
		for headB.Next != nil {
			pointers[headB.Next] += 1
			if pointers[headB.Next] > 1 {
				break
			}
			headB = headB.Next
		}
	}
	// Вывод узла пересечения из отображения
	for el, val := range pointers {
		if val > 1 {
			return el
		}
	}
	return nil
}
