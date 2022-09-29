// Функция принимает head связанного списка и число k.
// k - это количество частей, на которые надо разделить список.
// Функция делит список на k равных частей (каждая часть по длине
// не больше, чем на 1 в сравнении с другими частями).
// Возвращает массив связанных списков из k элементов,
// являющихся частями основного связанного списка.
// Если k будет больше количества элементов в списке,
// то вернётся массив с разбитым поэлементно списком и с пустыми элементами в конце массива.
// Input: head = [1,2,3], k = 5
// Output: [[1],[2],[3],[],[]]

package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var ret []*ListNode
	n := 0
	p := head

	// Подсчёт кол-ва элементов в списке
	for p != nil {
		n++
		p = p.Next
	}
	p = head
	// Остаток от деления показывает кол-во более длинных списков в начале массива
	rd := n % k
	// Если массив можно разбить поровну на k частей, то кол-во отрезаемых в начале узлов будет без изменения,
	// иначе rd элементов в начале массива будут на 1 больше по кол-ву узлов
	if rd == 0 {
		n = n / k
	} else {
		n = n/k + 1
	}
	for i := 0; i < k; i++ {
		if p != nil {
			ret = append(ret, cutList(p, n))
			for j := 0; j < n; j++ {
				p = p.Next
			}
		} else {
			ret = append(ret, nil)
		}
		rd--
		// Когда счётчик rd заканчивантся, то количество отрезаемых узлов уменьшается на 1
		if rd == 0 {
			n = n - 1
		}
	}
	return ret
}

// Функция отрезает с начала переданного List k узлов, делает из них новый List
// и возвращает указатель на новый List
func cutList(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	newList := new(ListNode)
	newList.Val = head.Val
	newList.Next = nil
	p := newList
	for k > 1 && head.Next != nil {
		head = head.Next
		node := new(ListNode)
		node.Val = head.Val
		node.Next = nil
		p.Next = node
		p = node
		k--
	}
	return newList
}
