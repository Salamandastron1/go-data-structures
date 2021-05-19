package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		fmt.Println("empty List")
		return
	}

	if l.head.data == value {
		fmt.Println("Head contained value, head deleted")
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l linkedList) printListData() {

	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("\n")
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 20}
	node3 := &node{data: 50}
	node4 := &node{data: 68}
	node5 := &node{data: 73}
	node6 := &node{data: 90}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	fmt.Println(myList)
	myList.printListData()
	myList.deleteWithValue(68)
	myList.printListData()

	myList.deleteWithValue(100)
	myList.printListData()

	myList.deleteWithValue(90)
	myList.printListData()

	emptyList := linkedList{}
	emptyList.deleteWithValue(50)

}
